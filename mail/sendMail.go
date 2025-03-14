package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func SendMail(recipients []string, subject, message, mail, username, password, smtpServer, smtpPort string) error {
	toHeader := strings.Join(recipients, ", ")

	body := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		mail, toHeader, subject, message)

	auth := smtp.PlainAuth("", username, password, smtpServer)

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, mail, recipients, []byte(body))
	if err != nil {
		log.Println("Error sending mail", err)
		return err
	}

	return nil
}
