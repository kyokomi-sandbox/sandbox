package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kyokomi/emoji"
	"github.com/mattn/go-sixel"
	"github.com/nfnt/resize"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(
		os.Getenv("SLACK_BOT_TOKEN"),
		slack.OptionDebug(false),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		//fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			//			fmt.Println("Infos:", ev.Info)
			//fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			//rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C2147483705"))

		case *slack.MessageEvent:
			c, _ := api.GetChannelInfo(ev.Channel)
			u, _ := api.GetUserInfo(ev.User)

			var channelName string
			if c != nil {
				channelName = c.Name
			}
			var username string
			if u != nil {
				username = u.Name

				userFilepath := "icons/" + username + filepath.Ext(u.Profile.Image48)
				f, err := os.Open(userFilepath)
				if err != nil {
					// ファイルがなかったりエラーなら取得しなおす
					getImage(userFilepath, u.Profile.Image48)
					f, err = os.Open(userFilepath)
				}

				// 最終的にファイルがエラーじゃなければdecodeして表示する
				if err == nil {
					img, _, err := image.Decode(f)
					if err != nil {
						fmt.Println(err)
					} else {
						sixel.NewEncoder(os.Stdout).Encode(resize.Resize(40, 40, img, resize.Bicubic))
					}
					f.Close()
				}
			}

			fmt.Println(channelName, username, emoji.Sprint(ev.Text))

			for _, a := range ev.Attachments {
				if a.ImageURL != "" {
					filePath := "temp" + filepath.Ext(a.ImageURL)
					getImage(filePath, a.ImageURL)
					f, err := os.Open(filePath)
					if err == nil {
						img, _, err := image.Decode(f)
						if err != nil {
							fmt.Println(err)
						} else {
							sixel.NewEncoder(os.Stdout).Encode(img)
						}
						f.Close()
					}
				}
			}

			// TODO: 画像がattachされてたらアイコンと同じ要領で表示する

		case *slack.PresenceChangeEvent:
			//fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			//fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			//fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			//fmt.Printf("Invalid credentials")
			return

		default:
			// Ignore other events..
			//fmt.Printf("Unexpected: %v\n", msg.Data)

		}
	}
}

func getImage(filePath string, url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
}
