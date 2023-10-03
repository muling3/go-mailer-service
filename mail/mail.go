package mail

import (
	"context"
	"fmt"
	"log"
	"net/smtp"

	"github.com/muling3/go-mailer/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gomail.v2"
)

func SendMailUsingGoMail(message models.MailMessage, config models.Config, client *mongo.Client) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailUser)
	m.SetHeader("To", message.ToAddress)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", message.Body)

	d := gomail.NewDialer(config.EmailHost, config.EmailPort, config.EmailUser, config.EmailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	// save a copy of the the message to the db
	coll := client.Database("logs").Collection("ecommerce-mails")

	result, err := coll.InsertOne(context.Background(), message)

	if err != nil {
		log.Fatalf("Error Listing Databases %v", err)
	}

	fmt.Printf("-> %s", result)
}

func SendEmail(message models.MailMessage, config models.Config, client *mongo.Client) {
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

	// save a copy of the the message to the db
	coll := client.Database("logs").Collection("ecommerce-mails")

	result, err := coll.InsertOne(context.Background(), message)

	if err != nil {
		log.Fatalf("Error Listing Databases %v", err)
	}

	fmt.Printf("-> %s", result)
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
