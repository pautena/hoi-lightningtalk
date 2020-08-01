package app

import (
	"fmt"
	"hoiLightningTalk/domain"
)


func Strikethrough(channel domain.SlackChannel, message domain.SlackMessage, messageService MessageService){
	newMessage := fmt.Sprintf("~%v~",message.Text)
	messageService.UpdateMessage(newMessage,channel.Id,message.Ts)
}

func Italic(channel domain.SlackChannel, message domain.SlackMessage, messageService MessageService){
	newMessage := fmt.Sprintf("_%v_",message.Text)
	messageService.UpdateMessage(newMessage,channel.Id,message.Ts)
}

func ThermonuclearWar(channel domain.SlackChannel, message domain.SlackMessage, messageService MessageService){
	nuclearUrl:="https://media.giphy.com/media/3o7abwbzKeaRksvVaE/giphy.gif"
	messageService.RepplyMessage(nuclearUrl,channel.Id,message.Ts)
}

func DeleteMessage(channel domain.SlackChannel, message domain.SlackMessage, messageService MessageService){
	messageService.DeleteMessage(channel.Id,message.Ts)
}