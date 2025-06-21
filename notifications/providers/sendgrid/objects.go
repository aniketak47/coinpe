package sendgrid

type SendgridWebhookResponse struct {
	Email                string `json:"email,omitempty"`
	Event                string `json:"event,omitempty"`
	Ip                   string `json:"ip,omitempty"`
	Reason               string `json:"reason,omitempty"`
	Response             string `json:"response,omitempty"`
	SgEventId            string `json:"sg_event_id,omitempty"`
	SgMessageId          string `json:"sg_message_id,omitempty"`
	SgMachineOpen        bool   `json:"sg_machine_open,omitempty"`
	SmtpId               string `json:"smtp-id,omitempty"`
	Status               string `json:"status,omitempty"`
	Timestamp            int    `json:"timestamp,omitempty"`
	Url                  string `json:"url,omitempty"`
	Useragent            string `json:"useragent,omitempty"`
	AsmGroupId           int    `json:"asm_group_id,omitempty"`
	Attempt              string `json:"attempt,omitempty"`
	BounceClassification string `json:"bounce_classification,omitempty"`
	Category             string `json:"category,omitempty"`
}

type CreateTemplateRequest struct {
	Name       string `json:"name"`
	Generation string `json:"generation"`
}

type CreateTemplateResponse struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Generation string        `json:"generation"`
	UpdatedAt  string        `json:"updated_at"`
	Versions   []interface{} `json:"versions"`
}

type CreateTemplateVersion struct {
	TemplateID           string `json:"template_id"`
	Active               int    `json:"active"`
	Name                 string `json:"name"`
	HTMLContent          string `json:"html_content"`
	GeneratePlainContent bool   `json:"generate_plain_content"`
	Subject              string `json:"subject"`
	Editor               string `json:"editor"`
}
