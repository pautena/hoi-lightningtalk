package app

import (
	"fmt"
	"hoiLightningTalk/domain"
)

func Strikethrough(channel domain.Channel, message domain.Message, messageService MessageService) {
	newMessage := fmt.Sprintf("~%v~", message.Text)
	messageService.UpdateMessage(newMessage, channel.ID, message.ID)
}

func Italic(channel domain.Channel, message domain.Message, messageService MessageService) {
	newMessage := fmt.Sprintf("_%v_", message.Text)
	messageService.UpdateMessage(newMessage, channel.ID, message.ID)
}

func ThermonuclearWar(channel domain.Channel, message domain.Message, messageService MessageService) {
	nuclearURL := "https://media.giphy.com/media/3o7abwbzKeaRksvVaE/giphy.gif"
	messageService.RepplyMessage(nuclearURL, channel.ID, message.ID)
}

func DeleteMessage(channel domain.Channel, message domain.Message, messageService MessageService) {
	messageService.DeleteMessage(channel.ID, message.ID)
}
