package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"code.google.com/p/go.net/websocket"
	"github.com/k0kubun/pp"
)

// @_apiCall 'rtm.start', {agent: 'node-slack'}, @_onLogin

func main() {
	log.SetFlags(log.Llongfile)

	params := url.Values{}
	params.Set("agent", "node-slack")
	apiCall("rtm.start", params, onLogin)
}

type SlackBotClient struct {
	authenticated bool
	socketURL     string
	// Team
	teamID     string
	teamName   string
	teamDomain string
}

var slackBotClient SlackBotClient

func onLogin(data []byte) {
	var s SlackResponse
	if err := json.Unmarshal(data, &s); err != nil {
		log.Fatalln(err)
	}

	if !s.Ok {
		slackBotClient.authenticated = false
		log.Fatalln(errors.New("login error"))
		return
	}
	//	pp.Println(s)

	slackBotClient.authenticated = true
	slackBotClient.socketURL = s.URL
	slackBotClient.teamID = s.Team.ID
	slackBotClient.teamName = s.Team.Name
	slackBotClient.teamDomain = s.Team.Domain

	pp.Println(slackBotClient)

	connect()
}

func connect() {
	var origin = "http://localhost:8000/"

	config, err := websocket.NewConfig(slackBotClient.socketURL, origin)
	if err != nil {
		log.Fatal(err)
	}

	var client net.Conn
	if config.Location == nil {
		log.Fatal(websocket.DialError{config, websocket.ErrBadWebSocketLocation})
	}
	if config.Origin == nil {
		log.Fatal(websocket.DialError{config, websocket.ErrBadWebSocketOrigin})
	}
	switch config.Location.Scheme {
	case "ws":
		client, err = net.Dial("tcp", config.Location.Host)

	case "wss":
		client, err = tls.Dial("tcp", config.Location.Host, config.TlsConfig)

	default:
		err = websocket.ErrBadScheme
	}

	pp.Println(config)
	pp.Println(client)

	if err != nil {
		log.Fatal(websocket.DialError{config, err})
	}

	var wsConn *websocket.Conn
	wsConn, err = websocket.NewClient(config, client)
	if err != nil {
		log.Fatal(websocket.DialError{config, err})
	}

	pp.Println(wsConn)
	return

	//	wsOpen, err := websocket.Dial(slackBotClient.socketURL, "", origin)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	message := []byte("hello, world!")
	//	_, err = wsOpen.Write(message)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("Send: %s\n", message)
	//
	//	var msg = make([]byte, 512)
	//	_, err = wsOpen.Read(msg)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("Receive: %s\n", msg)
}

type ApiResponse struct {
	Success bool
	Message string
}

const SLACK_API_URL = "https://api.slack.com/api/"
const SLACK_TOKEN = "xoxb-3278979370-hcGgEzF2fNOlO05gixHbCeAe"

func apiCall(method string, params url.Values, callback func([]byte)) error {
	params.Set("token", SLACK_TOKEN)

	u := SLACK_API_URL + method
	res, err := http.PostForm(u, params)
	pp.Println(res.Request)
	pp.Println(res.Header)
	if err != nil {
		if data, err := createErrorResponse(res.StatusCode); err != nil {
			return err
		} else {
			callback(data)
		}
	}

	pp.Println(res.StatusCode)

	if res.StatusCode == http.StatusOK {
		if resData, err := ioutil.ReadAll(res.Body); err != nil {
			return err
		} else {
			callback(resData)
		}
	} else {
		if data, err := createErrorResponse(res.StatusCode); err != nil {
			return err
		} else {
			callback(data)
		}
	}
	return nil
}

func createErrorResponse(statusCode int) ([]byte, error) {
	apiRes := ApiResponse{
		Success: false,
		Message: "API response: " + strconv.Itoa(statusCode),
	}
	if data, err := json.Marshal(apiRes); err != nil {
		return nil, errors.New("apiCall error: " + err.Error())
	} else {
		return data, nil
	}
}
