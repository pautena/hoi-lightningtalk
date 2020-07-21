package main

import (
	"fmt"
	"net/url"
	"strings"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"hoiLightningTalk/db"
	"hoiLightningTalk/app"
	"hoiLightningTalk/domain"
)

func GetPingUser(Text string) string{
	parts := strings.Split(Text, " ")
	return strings.ReplaceAll(parts[0],"@","")
}

func GetMessage(Text string, PingUser string) string{
	return strings.ReplaceAll(Text,fmt.Sprintf("@%v",PingUser),"")
}

func SendPing(PingUserId string, Message, ByUsername string) error{
	ur := db.NewUserRepository()

	PingUser,err := ur.GetUser(PingUserId)

	log.Printf(fmt.Sprintf("@%v",PingUser))

	if err != nil{
		return errors.New(fmt.Sprintf("%v user doesn't exists",PingUserId))
	}

	By,err := ur.GetUser(ByUsername)

	if err != nil{
		return errors.New(fmt.Sprintf("%v user doesn't exists",PingUserId))
	}

	var text string;
	if Message != ""{
		text = fmt.Sprintf("<@%v> sended to you a hoi %s",By.SlackId, Message)
	}else {
		text = fmt.Sprintf("<@%v> sended to you a hoi",By.SlackId)
	}

	attachments:= []domain.SlackAttachment{
		{
			Fallback: "You are unable to choose a game",
			CallbackID: "wopr_game",
			Color: "#3AA3E3",
			AttachmentType: "default",
			Actions: []domain.SlackAction{
				domain.SlackAction{
					Name: "game",
					Text: "Chess",
					Type: "button",
					Value: "chess",
				},
				domain.SlackAction{
						Name: "game",
						Text: "Falken's Maze",
						Type: "button",
						Value: "maze",
				},
				domain.SlackAction{
					Name: "game",
					Text: "Thermonuclear War",
						Style: "danger",
						Type: "button",
						Value: "war",
						Confirm: domain.SlackConfirm{
								Title: "Are you sure?",
								Text: "Wouldn't you prefer a good game of chess?",
								OkText: "Yes",
								DismissText: "No",
						},
					},
				},
			},
		}

	app.SendSlackMessageToUser(text,PingUser.SlackId,attachments)

	return nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	params,err := url.ParseQuery(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	text :=params.Get("text")
	byUsername :=params.Get("user_name")

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
	message := GetMessage(text,pingUser)

	err2:=SendPing(pingUser,message,byUsername)

	if err2!= nil{
		return events.APIGatewayProxyResponse{
			Body:       err2.Error(),
			StatusCode: 400,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
