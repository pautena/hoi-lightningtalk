package app

import (
	"fmt"
	"hoiLightningTalk/domain"
)


func Strikethrough(channel domain.Channel, message domain.Message, messageService MessageService){
	newMessage := fmt.Sprintf("~%v~",message.Text)
	messageService.UpdateMessage(newMessage,channel.Id,message.Id)
}

func Italic(channel domain.Channel, message domain.Message, messageService MessageService){
	newMessage := fmt.Sprintf("_%v_",message.Text)
	messageService.UpdateMessage(newMessage,channel.Id,message.Id)
}

func ThermonuclearWar(channel domain.Channel, message domain.Message, messageService MessageService){
	nuclearUrl:="https://media.giphy.com/media/3o7abwbzKeaRksvVaE/giphy.gif"
	messageService.RepplyMessage(nuclearUrl,channel.Id,message.Id)
}

func DeleteMessage(channel domain.Channel, message domain.Message, messageService MessageService){
	messageService.DeleteMessage(channel.Id,message.Id)
}