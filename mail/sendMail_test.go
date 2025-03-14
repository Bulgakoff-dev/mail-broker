package mail

import "testing"

const (
	smtpServer = "mail.eastern-gate.ru"
	smtpPort   = "587"

	username = "hello"
	password = "NovoBlack88"

	from = "hello@eastern-gate.ru"
	to   = "lithiumrobot@gmail.com"
)

func TestSendMail(t *testing.T) {
	recipients := []string{
		to,
	}

	err := SendMail(recipients, "Some Subject", "Message body", from, username, password, smtpServer, smtpPort)
	if err != nil {
		t.Error(err)
	}
}
