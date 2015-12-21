package actions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const contentType = "application/x-www-form-urlencoded"
const successMessage = "ok"

type SlackMessage struct {
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Text      string `json:"text"`
	Hook      string `json:"-"`
}

func NewSlackMessage(username string, emoji string, hook string) *SlackMessage {
	return &SlackMessage{Username: username, IconEmoji: emoji, Hook: hook}
}

func (sm *SlackMessage) Execute(text string) string {
	sm.Text = text
	marshalledPayload, err := json.Marshal(sm)
	if err != nil {
		panic(err)
	}

	stringPayload := string(marshalledPayload[:])
	data := url.Values{}
	data.Set("payload", stringPayload)

	resp, err := http.Post(sm.Hook, contentType, bytes.NewBufferString(data.Encode()))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if successMessage == string(responseBody[:]) {
		return "Successfully Sent Slack Message: " + text
	} else {
		return "Failed to send slack message: " + string(responseBody[:])
	}

}
