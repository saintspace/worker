package emailsend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type EmailSendTask struct {
	TemplateName  string            `json:"templateName"`
	SenderAddress string            `json:"senderAddress"`
	SubjectLine   string            `json:"subjectLine"`
	ToAddresses   []string          `json:"toAddresses"`
	Parameters    map[string]string `json:"parameters"`
}

func HandleTask(serializedTask string) error {
	task, err := parseTask(serializedTask)
	if err != nil {
		return fmt.Errorf("error while parsing EmailSendTask => %v", err.Error())
	}
	emailBody, err := generateEmailFromTemplate(task.TemplateName, task.Parameters, EmailTemplateCollection)
	if err != nil {
		return fmt.Errorf("error while generating email body => %v", err.Error())
	}
	return sendEmail(task, emailBody)
}

func parseTask(taskString string) (EmailSendTask, error) {
	task := EmailSendTask{}
	err := json.Unmarshal([]byte(taskString), &task)
	return task, err
}

func generateEmailFromTemplate(templateName string, params map[string]string, templates map[string]string) (string, error) {
	emailTemplates := map[string]*template.Template{}
	for k, v := range templates {
		emailTemplates[k] = template.Must(template.New("index").Parse(v))
	}
	currentTemplate, ok := emailTemplates[templateName]
	if !ok {
		return "", fmt.Errorf("unknown email template  => %s", templateName)
	}
	emailBody := &bytes.Buffer{}
	err := currentTemplate.Execute(emailBody, params)
	return emailBody.String(), err
}

func sendEmail(task EmailSendTask, emailBody string) error {
	email := &ses.SendEmailInput{
		Source: aws.String(task.SenderAddress),
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice(task.ToAddresses),
		},
		Message: &ses.Message{
			Subject: &ses.Content{Data: aws.String(task.SubjectLine)},
			Body:    &ses.Body{Text: &ses.Content{Data: aws.String(emailBody)}},
		},
	}
	// Initialize the AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ses.New(sess)
	// Send the email via Amazon Simple Email Service (SES)
	_, err := svc.SendEmail(email)
	if err != nil {
		return fmt.Errorf("error while sending email via SES => %v", err.Error())
	}
	return nil
}
