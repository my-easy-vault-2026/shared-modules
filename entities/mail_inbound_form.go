package entities

type MailgunInboundEmail struct {
	From            string `json:"from"`
	To              string `json:"to"`
	Recipient       string `json:"recipient"`
	Subject         string `json:"subject"`
	BodyPlain       string `json:"body_plain"`
	BodyHTML        string `json:"body_html"`
	StrippedText    string `json:"stripped_text"`
	StrippedHTML    string `json:"stripped_html"`
	StrippedSig     string `json:"stripped_signature"`
	MessageID       string `json:"message_id"`
	AttachmentCount string `json:"attachment_count"`
	MessageURL      string `json:"message_url"`
}
