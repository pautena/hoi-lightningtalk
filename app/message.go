package app

import (
	"hoiLightningTalk/domain"
)

type MessageService interface {
	SendMessageToHook(url string, msg string) domain.SlackResponse
	SendMessageToChannel(text string, channel string,attachments []domain.SlackAttachment) domain.SlackResponse
	UpdateMessage(text string, channel string,ts string) domain.SlackResponse
	RepplyMessage(text string, channel string,ts string) domain.SlackResponse
	DeleteMessage(channel string, ts string) domain.SlackResponse
}