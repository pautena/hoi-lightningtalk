package domain



type SlackUser struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type SlackChannel struct {
	Id string `json:"id"`
	Domain string `json:"domain"`
}

type SlackOriginalMessage struct {
	Text string `json:"text"`
	Ts string `json:"ts"`
	Attachments []SlackAttachment `json:"attachments"` 
}

type SlackActionCallback struct {
	Type       			string       					`json:"type,omitempty"`
	Actions    			[]SlackAction     		`json:"actions,omitempty"`
	CallbackId    	string      					`json:"channel_id"`
	Channel    			SlackChannel      		`json:"channel"`
	User    				SlackUser      				`json:"user"`
	ActionTs 				string 								`json:"action_ts"`
	MessageTs 			string 								`json:"message_ts"`
	OriginalMessage SlackOriginalMessage 	`json:"original_message"`
}