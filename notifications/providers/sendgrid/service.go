package sendgrid

import (
	"coinpe/notifications/providers"
	"coinpe/pkg/logger"
	"fmt"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type client struct {
	baseURL     string
	apiKey      string
	fromName    string
	fromAddress string
}

func (c *client) Send(payload *providers.NotificationPayload) (string, error) {
	email := mail.NewV3Mail()
	email.From = mail.NewEmail(payload.SenderName, payload.SenderEmail)
	email.Subject = payload.Subject
	email.TemplateID = payload.ProviderTemplateID

	personalization := mail.NewPersonalization()
	personalization.Subject = payload.Subject

	// Add TO recipients
	for _, recipient := range payload.Recipients {
		for _, toEmail := range recipient.ToEmails {
			personalization.AddTos(mail.NewEmail("", *toEmail))
		}
	}

	// Add CC recipients
	for _, recipient := range payload.Recipients {
		for _, ccEmail := range recipient.CCEmails {
			personalization.AddCCs(mail.NewEmail("", *ccEmail))
		}
	}

	// Add BCC recipients
	for _, recipient := range payload.Recipients {
		for _, bccEmail := range recipient.BCCEmails {
			personalization.AddBCCs(mail.NewEmail("", *bccEmail))
		}
	}

	personalization.DynamicTemplateData = map[string]interface{}{
		"DynamicTemplateData": map[string]interface{}{
			"otp":              payload.TemplateVariables.OTP,
			"username":         payload.TemplateVariables.Username,
			"user_role":        payload.TemplateVariables.UserRole,
			"to_mobile_number": payload.TemplateVariables.ToMobileNumber,
			"to_email":         payload.TemplateVariables.ToEmail,
		},
	}
	email.AddPersonalizations(personalization)

	client := sendgrid.NewSendClient(c.apiKey)
	res, err := client.Send(email)
	if res.StatusCode != http.StatusAccepted {
		return "", fmt.Errorf("%d: %s", res.StatusCode, res.Body)
	}
	msgID := res.Headers["X-Message-Id"][0]
	logger.Info("sendgrid msgID: ", msgID)
	return msgID, err
}

// SetBaseURL ...
func (sg *client) SetBaseURL(baseURL string) {
	sg.baseURL = baseURL
}

// func (sg *client) CreateTemplate(payload *providers.CreateTemplateRequest) (string, error) {
// 	var (
// 		templateResponse CreateTemplateResponse
// 	)
// 	templateRequest := CreateTemplateRequest{
// 		Name:       *payload.Name,
// 		Generation: generationDynamic,
// 	}
// 	byteRequest, err := json.Marshal(templateRequest)
// 	if err != nil {
// 		logger.Error("failed to create template in sendgrid: ", err)
// 		return "", err
// 	}
// 	request := sendgrid.GetRequest(sg.apiKey, "/v3/templates", sg.baseURL)
// 	request.Method = http.MethodPost
// 	request.Body = byteRequest
// 	response, err := sendgrid.API(request)
// 	if err != nil {
// 		logger.Error("failed to create template in sendgrid: ", err)
// 		return "", err
// 	}
// 	err = json.Unmarshal([]byte(response.Body), &templateResponse)
// 	if err != nil {
// 		logger.Error("failed to get sendgrid response: ", err)
// 		return "", err
// 	}

// 	createTemplateVersion := CreateTemplateVersion{
// 		TemplateID:           templateResponse.ID,
// 		Active:               1,
// 		Name:                 *payload.Name,
// 		HTMLContent:          *payload.Content,
// 		GeneratePlainContent: true,
// 		Subject:              *payload.Subject,
// 	}

// 	err = sg.CreateTemplateVersion(&createTemplateVersion)
// 	if err != nil {
// 		logger.Error("failed to create template version on sendgrid: ", err)
// 		return "", err
// 	}
// 	return templateResponse.ID, nil
// }
