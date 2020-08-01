package app

import (
	"hoiLightningTalk/domain"
)

type MessageService interface {
	SendMessageToHook(url string, msg string) domain.MessageResponse
	SendMessageToChannel(text string, channel string, attachments []domain.Attachment) domain.MessageResponse
	UpdateMessage(text string, channel string, ts string) domain.MessageResponse
	RepplyMessage(text string, channel string, ts string) domain.MessageResponse
	DeleteMessage(channel string, ts string) domain.MessageResponse
}
