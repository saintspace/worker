package email

import (
	"fmt"
)

// SES is responsible for interfacing with AWS Simple Email Service
type EmailService struct {
	emailSender iEmailSender
}

func New(emailSender iEmailSender) *EmailService {
	return &EmailService{
		emailSender: emailSender,
	}
}

type iEmailSender interface {
	SendEmail(
		senderAddress string,
		toAddresses []string,
		subjectLine string,
		emailBody string,
	) error
}

func (s *EmailService) SendTemplatedEmail(
	templateName string,
	templateParams map[string]string,
	subjectLine string,
	senderAddress string,
	toAddresses []string,
) error {
	emailBody, err := generateEmailBodyFromTemplate(templateName, templateParams)
	if err != nil {
		return fmt.Errorf("error while generating email body => %v", err.Error())
	}
	return s.emailSender.SendEmail(senderAddress, toAddresses, subjectLine, emailBody)
}
