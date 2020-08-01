package slack

import (
	"bytes"
	"log"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"hoiLightningTalk/domain"
	"hoiLightningTalk/app"
)


type SlackService struct {
	postMessageUrl string
	chatUpdateUrl string
	chatDeleteUrl string
	parser SlackParser
}

func NewSlackService() app.MessageService{
	return SlackService {
		postMessageUrl: "https://slack.com/api/chat.postMessage",
		chatUpdateUrl: "https://slack.com/api/chat.update",
		chatDeleteUrl: "https://slack.com/api/chat.delete",
		parser: NewSlackParser(),
	}
}


func (ss SlackService) sendMessage(url string, rBody []byte) domain.MessageResponse{
	client := &http.Client{}
	req,err := http.NewRequest("POST",url,bytes.NewBuffer(rBody))

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	secret :=os.Getenv("SLACK_BOT_SECRET")

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v",secret))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)

	if err!=nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	log.Println(string(body))
	
	var response SlackResponse
	err = json.Unmarshal([]byte(body), &response)

	if err!=nil {
		log.Println(err)
		return domain.MessageResponse{}
	}
	
	
	return ss.parser.ParseResponse(response)
}



func (ss SlackService) SendMessageToHook(url string, msg string) domain.MessageResponse{

	rBody,err := json.Marshal(SlackPayload{
		Text:msg,
	})	

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.sendMessage(url,rBody)
}

func (ss SlackService) SendMessageToChannel(text string, channel string,attachments []domain.Attachment) domain.MessageResponse{
	rBody,err := json.Marshal(SlackPayload{
		Channel:channel,
		Text:text,
		Attachments:ss.parser.ParseAttachments(attachments),
		UnfurlLinks:false,
		UnfurlMedia:false,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.sendMessage(ss.postMessageUrl,rBody)
}

func (ss SlackService) UpdateMessage(text string, channel string,ts string) domain.MessageResponse{
	rBody,err := json.Marshal(SlackPayload{
		Channel:channel,
		Text:text,
		Ts:ts,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.sendMessage(ss.chatUpdateUrl,rBody)
}

func (ss SlackService) RepplyMessage(text string, channel string,ts string) domain.MessageResponse{
	rBody,err := json.Marshal(SlackPayload{
		Channel:channel,
		Text:text,
		ThreadTs:ts,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}

	return ss.sendMessage(ss.postMessageUrl,rBody)
}

func (ss SlackService) DeleteMessage(channel string, ts string) domain.MessageResponse {
	rBody,err := json.Marshal(SlackPayload{
		Channel:channel,
		Ts:ts,
	})

	if err != nil {
		log.Println(err)
		return domain.MessageResponse{}
	}
	return ss.sendMessage(ss.chatDeleteUrl,rBody)
}