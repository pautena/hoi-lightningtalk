package domain

type SlackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type SlackConfirm struct {
	Title 				string `json:"title"`
	Text 					string `json:"text"`
	OkText 				string `json:"ok_text"`
	DismissText 	string `json:"dismiss_text"`
}

type SlackAction struct {
	Type		string   			`json:"type"`
	Text		string   			`json:"text"`
	Url 		string   			`json:"url"`
	Style		string   			`json:"style"`
	Name		string   			`json:"name"`
	Value		string   			`json:"value"`
	Confirm SlackConfirm	`json:"confirm"`
}

type SlackAttachment struct {
	Fallback     string   			`json:"fallback"`
	Color        string   			`json:"color"`
	PreText      string   			`json:"pretext"`
	AuthorName   string   			`json:"author_name"`
	AuthorLink   string   			`json:"author_link"`
	AuthorIcon   string   			`json:"author_icon"`
	Title        string   			`json:"title"`
	TitleLink    string   			`json:"title_link"`
	Text         string   			`json:"text"`
	ImageUrl     string   			`json:"image_url"`
	Fields       []SlackField  	`json:"fields"`
	Footer       string   			`json:"footer"`
	FooterIcon   string   			`json:"footer_icon"`
	Timestamp    int64    			`json:"ts"`
	MarkdownIn   []string 			`json:"mrkdwn_in"`
	Actions      []SlackAction 	`json:"actions"`
	CallbackID   string   			`json:"callback_id"`
	ThumbnailUrl string   			`json:"thumb_url"`
	AttachmentType string   		`json:"attachment_type"`
}

type SlackPayload struct {
	Parse       string       			`json:"parse,omitempty"`
	Username    string       			`json:"username,omitempty"`
	IconUrl     string       			`json:"icon_url,omitempty"`
	IconEmoji   string       			`json:"icon_emoji,omitempty"`
	Channel     string       			`json:"channel,omitempty"`
	Text        string       			`json:"text,omitempty"`
	LinkNames   string       			`json:"link_names,omitempty"`
	Attachments []SlackAttachment `json:"attachments,omitempty"`
	UnfurlLinks bool         			`json:"unfurl_links,omitempty"`
	UnfurlMedia bool         			`json:"unfurl_media,omitempty"`
	Markdown    bool         			`json:"mrkdwn,omitempty"`
}