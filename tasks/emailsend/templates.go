package emailsend

var EmailTemplateCollection = map[string]string{
	"email-subscription-verification": "Thank you for your interest in SaintSpace!\n\nPlease click on the link below to confirm your subscription. It helps us make sure you're human.\n\n {{.verificationLink}}\n\n",
}
