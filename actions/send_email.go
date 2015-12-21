package actions

type SendEmail struct {
	smtpServer string
	recipient  string
	sender     string
	subject    string
	body       string
}

func NewSendEmailAction(smtpServer string, recipient string, sender string, subject string, body string) *SendEmail {
	return &SendEmail{smtpServer, recipient, sender, subject, body}
}

func (se *SendEmail) Execute(parameters string) string {
	return "Sent email to " + se.recipient
}
