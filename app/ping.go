package app

import (
	"fmt"
	"hoiLightningTalk/domain"
	"log"
)

func SendPing(userID string, message, byUsername string, userRepo UserRepository, messageService MessageService) error {
	PingUser, err := userRepo.GetUser(userID)

	if err != nil {
		return fmt.Errorf("%v user doesn't exists", userID)
	}

	byUser, err := userRepo.GetUser(byUsername)

	if err != nil {
		return fmt.Errorf("%v user doesn't exists", userID)
	}

	var text string
	if message != "" {
		text = fmt.Sprintf("<@%v> sended to you a hoi %s", byUser.AppID, message)
	} else {
		text = fmt.Sprintf("<@%v> sended to you a hoi", byUser.AppID)
	}

	attachments := GetAttachments()

	response := messageService.SendMessageToChannel(text, PingUser.AppID, attachments)
	log.Println(response)

	return nil
}

func GetAttachments() []domain.Attachment {
	return []domain.Attachment{
		{
			Fallback:       "You are unable to choose a game",
			CallbackID:     "action_callback_id",
			Color:          "#3AA3E3",
			AttachmentType: "default",
			Actions: []domain.Action{
				{
					Name:  "strikethrough",
					Text:  "Strikethrough",
					Type:  "button",
					Value: "strikethrough",
				},
				{
					Name:  "italic",
					Text:  "Italic",
					Type:  "button",
					Value: "italic",
				},
				{
					Name:  "war",
					Text:  "Thermonuclear War",
					Style: "danger",
					Type:  "button",
					Value: "war",
					Confirm: domain.Confirm{
						Title:       "Are you sure?",
						Text:        "Wouldn't you prefer something less permanent?",
						OkText:      "Yes",
						DismissText: "No",
					},
				},
				{
					Name:  "delete",
					Text:  "Delete message",
					Style: "danger",
					Type:  "button",
					Value: "delete",
					Confirm: domain.Confirm{
						Title:       "Are you sure?",
						Text:        "Does you want to delete this message?",
						OkText:      "Yes",
						DismissText: "No",
					},
				},
			},
		},
	}
}
