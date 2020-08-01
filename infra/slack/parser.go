package slack

import (
	"hoiLightningTalk/domain"
)



type SlackParser struct {

}

func NewSlackParser () SlackParser {
	return SlackParser {}
}

func (slackParser SlackParser) ParseResponse(slackResponse SlackResponse) domain.MessageResponse {
	return domain.MessageResponse {
		Ok:slackResponse.Ok,
		Channel: slackResponse.Channel,
		Id: slackResponse.Ts,
	}
}

func (slackParser SlackParser) ParseAttachments(attachments []domain.Attachment) []SlackAttachment{

	var slackAttachments []SlackAttachment

	for _, attachment := range attachments {
		slackAttachments = append(slackAttachments, SlackAttachment{
			Fallback: attachment.Fallback,
			Color: attachment.Color,
			CallbackID: attachment.CallbackID,
			AttachmentType: attachment.AttachmentType,
			Actions:slackParser.parseActions(attachment.Actions),
		})
	}
	return slackAttachments

}

func (slackParser SlackParser) parseActions(actions []domain.Action) []SlackAction {
	var slackActions []SlackAction

	for _, action := range actions {
		slackActions = append(slackActions, SlackAction{
			Type:action.Type,
			Text:action.Text,
			Style:action.Style,
			Name:action.Name,
			Value:action.Value,
			Confirm: SlackConfirm {
				Text:action.Text,
			},
		})
	}
	return slackActions
}