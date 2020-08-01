package app

import (
	"errors"
	"fmt"
	"log"
	"hoiLightningTalk/domain"
)

func SendPing(PingUserId string, Message, ByUsername string, userRepo UserRepository,messageService MessageService) error{
	PingUser,err := userRepo.GetUser(PingUserId)

	if err != nil{
		return errors.New(fmt.Sprintf("%v user doesn't exists",PingUserId))
	}

	By,err := userRepo.GetUser(ByUsername)

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
			CallbackID: "action_callback_id",
			Color: "#3AA3E3",
			AttachmentType: "default",
			Actions: []domain.SlackAction{
				domain.SlackAction{
					Name: "strikethrough",
					Text: "Strikethrough",
					Type: "button",
					Value: "strikethrough",
				},
				domain.SlackAction{
					Name: "italic",
					Text: "Italic",
					Type: "button",
					Value: "italic",
				},
				domain.SlackAction{
					Name: "war",
					Text: "Thermonuclear War",
					Style: "danger",
					Type: "button",
					Value: "war",
					Confirm: domain.SlackConfirm{
						Title: "Are you sure?",
						Text: "Wouldn't you prefer something less permanent?",
						OkText: "Yes",
						DismissText: "No",
					},
				},
				domain.SlackAction{
					Name: "delete",
					Text: "Delete message",
					Style: "danger",
					Type: "button",
					Value: "delete",
					Confirm: domain.SlackConfirm{
						Title: "Are you sure?",
						Text: "Does you want to delete this message?",
						OkText: "Yes",
						DismissText: "No",
					},
				},
			},
		},
	}

	response := messageService.SendMessageToChannel(text,PingUser.SlackId,attachments)
	log.Println(response)

	return nil
}