package main

import (
	"fmt"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hoiLightningTalk/app"
)


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	params,err := url.ParseQuery(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	team :=params.Get("team_domain")
	uID :=params.Get("user_id")
	username :=params.Get("user_name")
	rUrl :=params.Get("response_url")
	text :=params.Get("text")

	if team == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Undefined team_domain",
			StatusCode: 400,
		}, nil
	}

	if uID == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Undefined user_id",
			StatusCode: 400,
		}, nil
	}

	if username == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Undefined username",
			StatusCode: 400,
		}, nil
	}


	app.SignIn(uID,username,text)

	msg :=fmt.Sprintf("Hello, user %v, from team %v",uID,team)

	if rUrl != "" {
		app.SendSlackMessageToUrl(rUrl,msg)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
