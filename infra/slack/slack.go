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
}

func NewSlackService() app.MessageService{
	return SlackService {
		postMessageUrl: "https://slack.com/api/chat.postMessage",
		chatUpdateUrl: "https://slack.com/api/chat.update",
		chatDeleteUrl: "https://slack.com/api/chat.delete",
	}
}


func (ss SlackService) sendMessage(url string, rBody []byte) domain.SlackResponse{
	client := &http.Client{}
	req,err := http.NewRequest("POST",url,bytes.NewBuffer(rBody))

	if err != nil {
		log.Println(err)
		return domain.SlackResponse{}
	}

	secret :=os.Getenv("SLACK_BOT_SECRET")

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v",secret))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return domain.SlackResponse{}
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)

	if err!=nil {
		log.Println(err)
		return domain.SlackResponse{}
	}

	log.Println(string(body))
	
	var response domain.SlackResponse;
	err = json.Unmarshal([]byte(body), &response)

	if err!=nil {
		log.Println(err)
		return domain.SlackResponse{}
	}
	
	
	return response
}



func (ss SlackService) SendMessageToHook(url string, msg string) domain.SlackResponse{

	rBody,err := json.Marshal(domain.SlackPayload{
		Text:msg,
	})	

	if err != nil {
		log.Println(err)
		return domain.SlackResponse{}
	}

	return ss.sendMessage(url,rBody)
}

func (ss SlackService) SendMessageToChannel(text string, channel string,attachments []domain.SlackAttachment) domain.SlackResponse{
	rBody,err := json.Marshal(domain.SlackPayload{
		Channel:channel,
		Text:text,
		Attachments:attachments,
		UnfurlLinks:false,
		UnfurlMedia:false,
	})

	if err != nil {
		log.Println(err)
		return domain.SlackResponse{}
	}

	return ss.sendMessage(ss.postMessageUrl,rBody)
}

func (ss SlackService) UpdateMessage(text string, channel string,ts string) domain.SlackResponse{
	rBody,err := json.Marshal(domain.SlackPayload{
		Channel:channel,
		Text:text,
		Ts:ts,
	})

	if err != nil {
		log.Println(err)
		return domain.SlackResponse{}
	}

	return ss.sendMessage(ss.chatUpdateUrl,rBody)
}

func (ss SlackService) RepplyMessage(text string, channel string,ts string) domain.SlackResponse{
	rBody,err := json.Marshal(domain.SlackPayload{
		Channel:channel,
		Text:text,
		ThreadTs:ts,
	})

	if err != nil {
		log.Println(err)
		return domain.SlackResponse{}
	}

	return ss.sendMessage(ss.postMessageUrl,rBody)
}

func (ss SlackService) DeleteMessage(channel string, ts string) domain.SlackResponse {
	rBody,err := json.Marshal(domain.SlackPayload{
		Channel:channel,
		Ts:ts,
	})

	if err != nil {
		log.Println(err)
		return domain.SlackResponse{}
	}

	return ss.sendMessage(ss.chatDeleteUrl,rBody)
}