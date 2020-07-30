package main

import (
	"log"
	"net/url"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hoiLightningTalk/domain"

)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println(request.Body)

	params,err := url.ParseQuery(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}


	payload :=params.Get("payload")

	var slackAction domain.SlackAction;
	err = json.Unmarshal([]byte(payload), &slackAction)

	if err!=nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:"Action success",
	}, nil
}

func main() {
	lambda.Start(handler)
}
