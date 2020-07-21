package app

import (
	"bytes"
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"hoiLightningTalk/domain"
)


func sendMessage(url string, rBody []byte) {
	client := &http.Client{}
	req,err := http.NewRequest("POST",url,bytes.NewBuffer(rBody))

	if err != nil {
		log.Println(err)
		return
	}

	secret :=os.Getenv("SLACK_BOT_SECRET")

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v",secret))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)

	if err!=nil {
		log.Println(err)
		return
	}
	
	log.Println(string(body))
}



func SendSlackMessageToUrl(url string, msg string){

	rBody,err := json.Marshal(domain.SlackPayload{
		Text:msg,
	})	

	if err != nil {
		log.Println(err)
		return
	}

	sendMessage(url,rBody)
}

func SendSlackMessageToUser(text string, channel string,attachments []domain.SlackAttachment){
	rBody,err := json.Marshal(domain.SlackPayload{
		Channel:channel,
		Text:text,
		Attachments:attachments,
		UnfurlLinks:false,
		UnfurlMedia:false,
	})

	if err != nil {
		log.Println(err)
		return
	}

	sendMessage("https://slack.com/api/chat.postMessage",rBody)
}