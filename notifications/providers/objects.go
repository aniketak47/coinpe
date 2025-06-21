package providers

import "time"

type Recipient struct {
	Name      string    `json:"name"`
	ToEmails  []*string `json:"to_emails"`
	CCEmails  []*string `json:"cc_emails"`
	BCCEmails []*string `json:"bcc_emails"`
	Mobile    string    `json:"mobile"`
	DeviceID  string    `json:"device_id"`
}

type TemplateVariables struct {
	OTP            string `json:"otp,omitempty"`
	Username       string `json:"username,omitempty"`
	UserRole       string `json:"user_role,omitempty"`
	ToMobileNumber string `json:"to_mobile_number,omitempty"`
	ToEmail        string `json:"to_email,omitempty"`
	FirstName      string `json:"customer_first_name,omitempty"`
	LastName       string `json:"customer_last_name"`
	Link           string `json:"link,omitempty"`
}

type NotificationPayload struct {
	TemplateName       string      `json:"template_name" binding:"required"`
	Channel            string      `json:"channel" binding:"required"`
	ProviderKey        string      `json:"provider_key"`
	SenderID           string      `json:"sender_id"`
	SenderName         string      `json:"sender_name"`
	SenderEmail        string      `json:"sender_email"`
	ScheduleAt         *time.Time  `json:"schedule_at"`
	Recipients         []Recipient `json:"recipients" binding:"required"`
	Subject            string      `json:"subject"`
	Body               string
	TemplateVariables  TemplateVariables `json:"template_vars"`
	Category           string            `json:"category,omitempty"`
	ProviderTemplateID string            `json:"provider_template_id"`
}

type CreateTemplateRequest struct {
	Name     *string `json:"name"`
	Channel  *string `json:"channel"`
	Content  *string `json:"content"`
	Subject  *string `json:"subject"`
	Category *string `json:"category"`
}
