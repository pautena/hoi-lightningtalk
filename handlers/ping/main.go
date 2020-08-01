package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"hoiLightningTalk/app"
	"hoiLightningTalk/infra/mgo"
	"hoiLightningTalk/infra/slack"
)

func GetPingUser(text string) string {
	parts := strings.Split(text, " ")
	return strings.ReplaceAll(parts[0], "@", "")
}

func GetMessage(text string, pingUser string) string {
	return strings.ReplaceAll(text, fmt.Sprintf("@%v", pingUser), "")
}

func handler(request events.APIGatewayProxyRequest, userRepo app.UserRepository, messageService app.MessageService) (events.APIGatewayProxyResponse, error) {
	params, err := url.ParseQuery(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	text := params.Get("text")
	byUsername := params.Get("user_name")

	if byUsername == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Undefined user_name",
			StatusCode: 400,
		}, nil
	}

	if text == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Undefined text",
			StatusCode: 400,
		}, nil
	}

	pingUser := GetPingUser(text)
	message := GetMessage(text, pingUser)
	pingedCallback, _ := app.GetUserURL(pingUser, userRepo)

	err = app.SendPing(pingUser, message, byUsername, userRepo, messageService)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       pingedCallback,
	}, nil
}

func main() {
	userRepo := mgo.NewMongoUserRepository()
	messageService := slack.NewSlackService()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return handler(request, userRepo, messageService)
	})
}
