package mail

import (
	"fmt"
	"net/smtp"

	"github.com/muling3/go-mailer/models"
)

func SendEmail(message models.MailMessage, config models.Config) {
	auth := initEmail(config)

	// smtp server configuration.
	smtpServer := models.SmtpServer{Host: config.EmailHost, Port: fmt.Sprintf("%d", config.EmailPort)}

	headers := []byte("From: " + message.FromAddress + "\r\n" +
		"Subject: " + message.Subject + "\r\n" +
		"\r\n")

	body := append(headers, message.Body...)

	// Sending email.
	err := smtp.SendMail(smtpServer.Host+":"+smtpServer.Port, auth, message.FromAddress, []string{message.ToAddress}, body)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")

}

func initEmail(config models.Config) smtp.Auth {
	from := config.EmailUser
	password := config.EmailPassword

	// smtp server configuration.
	smtpServer := models.SmtpServer{Host: config.EmailHost, Port: fmt.Sprintf("%d", config.EmailPort)}

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.Host)
	return auth
}
