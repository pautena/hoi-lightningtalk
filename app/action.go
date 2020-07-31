package app

import (
	"fmt"
	"hoiLightningTalk/domain"
)


func Strikethrough(channel domain.SlackChannel, message domain.SlackMessage){
	newMessage := fmt.Sprintf("~%v~",message.Text)
	UpdateSlackMessage(newMessage,channel.Id,message.Ts)
}

func Italic(channel domain.SlackChannel, message domain.SlackMessage){
	newMessage := fmt.Sprintf("_%v_",message.Text)
	UpdateSlackMessage(newMessage,channel.Id,message.Ts)
}

func ThermonuclearWar(channel domain.SlackChannel, message domain.SlackMessage){
	nuclearUrl:="https://media.giphy.com/media/3o7abwbzKeaRksvVaE/giphy.gif"
	RepplySlackMessage(nuclearUrl,channel.Id,message.Ts)
}

func DeleteMessage(channel domain.SlackChannel, message domain.SlackMessage){
	DeleteSlackMessage(channel.Id,message.Ts)
}