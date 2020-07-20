package main

import (
	"fmt"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"hoiLightningTalk/domain"
	"hoiLightningTalk/db"
	"hoiLightningTalk/app"
)

func SignIn(Id string){
	ur := db.NewUserRepository()
	ur.SaveUser(domain.User{Id:Id})
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	params,err := url.ParseQuery(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	team :=params.Get("team_domain")
	uID :=params.Get("user_id")
	rUrl :=params.Get("response_url")

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


	SignIn(uID)

	msg :=fmt.Sprintf("Hello, user %v, from team %v",uID,team)

	if rUrl != "" {
		app.SendSlackMessageToUrl(rUrl,fmt.Sprintf("_(Hook message)_ %v",msg))
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("_(Response message)_ %v",msg),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
