package domain


//Action

type ActionCallback struct {
	Type       			string       					`json:"type,omitempty"`
	Actions    			[]Action     		`json:"actions,omitempty"`
	CallbackId    	string      					`json:"channel_id"`
	Channel    			Channel      		`json:"channel"`
	User    				User      				`json:"user"`
	ActionTs 				string 								`json:"action_ts"`
	MessageTs 			string 								`json:"message_ts"`
	OriginalMessage Message 	`json:"original_message"`
}