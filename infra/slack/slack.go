package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hoiLightningTalk/app"
	"hoiLightningTalk/domain"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SlackService struct {
	postMessageURL string
	chatUpdateURL  string
	chatDeleteURL  string
	parser         SlackParser
}

func NewSlackService() app.MessageService {
	return SlackService{
		postMessageURL: "https://slack.com/api/chat.postMessage",
		chatUpdateURL:  "https://slack.com/api/chat.update",
		chatDeleteURL:  "https://slack.com/api/chat.delete",
		parser:         NewSlackParser(),
	}
}

func (ss SlackService) sendMessage(url string, rBody []byte) domain.MessageResponse {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(rBody))

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	secret := os.Getenv("SLACK_BOT_SECRET")

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", secret))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	log.Println(string(body))

	var response SlackResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.parser.ParseResponse(response)
}

func (ss SlackService) SendMessageToHook(url string, msg string) domain.MessageResponse {

	rBody, err := json.Marshal(SlackPayload{
		Text: msg,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.sendMessage(url, rBody)
}

func (ss SlackService) SendMessageToChannel(text string, channel string, attachments []domain.Attachment) domain.MessageResponse {
	rBody, err := json.Marshal(SlackPayload{
		Channel:     channel,
		Text:        text,
		Attachments: ss.parser.ParseAttachments(attachments),
		UnfurlLinks: false,
		UnfurlMedia: false,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.sendMessage(ss.postMessageURL, rBody)
}

func (ss SlackService) UpdateMessage(text string, channel string, ts string) domain.MessageResponse {
	rBody, err := json.Marshal(SlackPayload{
		Channel: channel,
		Text:    text,
		TS:      ts,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.sendMessage(ss.chatUpdateURL, rBody)
}

func (ss SlackService) RepplyMessage(text string, channel string, ts string) domain.MessageResponse {
	rBody, err := json.Marshal(SlackPayload{
		Channel:  channel,
		Text:     text,
		ThreadTS: ts,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.sendMessage(ss.postMessageURL, rBody)
}

func (ss SlackService) DeleteMessage(channel string, ts string) domain.MessageResponse {
	rBody, err := json.Marshal(SlackPayload{
		Channel: channel,
		TS:      ts,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}
	return ss.sendMessage(ss.chatDeleteURL, rBody)
}
