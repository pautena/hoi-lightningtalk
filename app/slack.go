package app

import (
	"bytes"
	"log"
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)



func SendSlackMessageToUrl(url string, msg string){

	rBody,err := json.Marshal(map[string]string{
		"text":msg,
	})
	
	client := &http.Client{}
	req,err := http.NewRequest("POST",url,bytes.NewBuffer(rBody))

	if err != nil {
		log.Println(err)
		return
	}

	secret :=os.Getenv("SLACK_BOT_SECRET")

	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v",secret))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	//TODO add headers

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)

	if err!=nil {
		log.Println(err)
		return
	}
	
	log.Println(string(body))

}