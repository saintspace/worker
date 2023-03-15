package handler

// SES is responsible for interfacing with AWS Simple Email Service
type EventHandler struct {
	emailService iEmailService
}

func New(emailService iEmailService) *EventHandler {
	return &EventHandler{
		emailService: emailService,
	}
}

type iEmailService interface {
	SendTemplatedEmail(
		templateName string,
		templateParams map[string]string,
		subjectLine string,
		senderAddress string,
		toAddresses []string,
	) error
}
