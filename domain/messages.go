package domain


type Message struct {
	Text string `json:"text"`
	Id string `json:"ts"`
	Attachments []Attachment `json:"attachments"` 
}

type Channel struct {
	Id string `json:"id"`
	Domain string `json:"domain"`
}

type Attachment struct {
	Fallback     string
	Color        string
	CallbackID   string
	AttachmentType string
	Actions      []Action
	
}

type Action struct {
	Text		string
	Name		string	
	Value		string
	Type		string
	Style		string
	Confirm Confirm
}

type Confirm struct {
	Title 				string
	Text 					string
	OkText 				string
	DismissText 	string
}

type MessageResponse struct {
	Ok				bool
	Channel		string
	Id				string
}