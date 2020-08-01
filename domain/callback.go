package domain

//Action

type ActionCallback struct {
	Type            string   `json:"type,omitempty"`
	Actions         []Action `json:"actions,omitempty"`
	CallbackID      string   `json:"channel_id"`
	Channel         Channel  `json:"channel"`
	User            User     `json:"user"`
	ActionTS        string   `json:"action_ts"`
	MessageTS       string   `json:"message_ts"`
	OriginalMessage Message  `json:"original_message"`
}
