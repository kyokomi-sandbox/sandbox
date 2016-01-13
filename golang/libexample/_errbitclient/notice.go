package _errbitclient

import (
	"fmt"
	"net/http"
)

// TODO: あとでなんとかする
const sampleURL = "http://104.155.235.152:3000"

func getCreateNoticeURL(apiKey string) string {
	return fmt.Sprintf(
		sampleURL + "/api/v3/projects/%s/notices?key=%s",
		apiKey, apiKey)
}

type Error struct {
	Type      string       `json:"type"`
	Message   string       `json:"message"`
	Backtrace []StackFrame `json:"backtrace"`
}

type notifier struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	URL     string `json:"url"`
}

type Notice struct {
	Notifier notifier               `json:"notifier"`
	Errors   []Error                `json:"errors"`
	Context  map[string]string      `json:"context"`
	Env      map[string]interface{} `json:"environment"`
	Session  map[string]interface{} `json:"session"`
	Params   map[string]interface{} `json:"params"`
}

func NewNotice(e interface{}, stack []StackFrame, req *http.Request) *Notice {
	notice := &Notice{
		Notifier: notifier{
			Name:    "gobrake",
			Version: "1.0",
			URL:     "https://github.com/airbrake/gobrake",
		},
		Errors: []Error{
			{
				Type:      fmt.Sprintf("%T", e),
				Message:   fmt.Sprint(e),
				Backtrace: stack,
			},
		},
		Context: map[string]string{},
		Env:     map[string]interface{}{},
		Session: map[string]interface{}{},
		Params:  map[string]interface{}{},
	}

	if req != nil {
		notice.Context["url"] = req.URL.String()
		if ua := req.Header.Get("User-Agent"); ua != "" {
			notice.Context["userAgent"] = ua
		}

		for k, v := range req.Header {
			if len(v) == 1 {
				notice.Env[k] = v[0]
			} else {
				notice.Env[k] = v
			}
		}

		// TODO: jsonがとれてない
		if err := req.ParseForm(); err == nil {
			for k, v := range req.Form {
				if len(v) == 1 {
					notice.Params[k] = v[0]
				} else {
					notice.Params[k] = v
				}
			}
		}
	}

	return notice
}
