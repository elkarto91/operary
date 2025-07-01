package model

// WebhookRequest defines the payload for sending a webhook.
type WebhookRequest struct {
	URL     string            `json:"url"`
	Method  string            `json:"method,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Payload map[string]any    `json:"payload,omitempty"`
}

// EmailRequest defines the payload for sending an email via SMTP.
type EmailRequest struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}
