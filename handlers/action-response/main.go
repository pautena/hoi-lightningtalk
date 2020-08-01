package main

import (
	"log"
	"net/url"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hoiLightningTalk/domain"
	"hoiLightningTalk/app"
	"hoiLightningTalk/infra/slack"

)

func handler(request events.APIGatewayProxyRequest, messageService app.MessageService) (events.APIGatewayProxyResponse, error) {

	log.Println(request.Body)

	params,err := url.ParseQuery(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}


	payload :=params.Get("payload")

	var callback domain.ActionCallback;
	err = json.Unmarshal([]byte(payload), &callback)

	if err!=nil {
		return events.APIGatewayProxyResponse{}, err
	}

	actionValue  := callback.Actions[0].Value

	if (actionValue == "strikethrough"){
		app.Strikethrough(callback.Channel,callback.OriginalMessage, messageService)
	}else if (actionValue == "italic"){
		app.Italic(callback.Channel,callback.OriginalMessage, messageService)
	}else if (actionValue == "war"){
		app.ThermonuclearWar(callback.Channel,callback.OriginalMessage, messageService)
	}else if (actionValue == "delete"){
		app.DeleteMessage(callback.Channel,callback.OriginalMessage, messageService)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	messageService:= slack.NewSlackService();
	lambda.Start(func (request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
		return handler(request,messageService)
	})
}
