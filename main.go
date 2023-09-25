package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/muling3/go-mailer/mail"
	"github.com/muling3/go-mailer/models"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func main() {
	config, err := envConfig()
	if err != nil {
		log.Fatal("Error reading application config ", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", config)

	connection, _ := amqp.Dial(config.BrokerUrl)
	defer connection.Close()

	go func(con *amqp.Connection) {
		channel, _ := connection.Channel()
		defer channel.Close()
		durable, exclusive := true, false
		autoDelete, noWait := false, true
		q, _ := channel.QueueDeclare(config.QueueName, durable, autoDelete, exclusive, noWait, nil)
		channel.QueueBind(q.Name, config.RoutingKey, config.Topic, false, nil)
		autoAck, exclusive, noWait, noLocal := false, false, false, false
		messages, _ := channel.Consume(q.Name, "", autoAck, exclusive, noLocal, noWait, nil)
		// multiAck := false
		for msg := range messages {
			fmt.Println("Body:", string(msg.Body), "Timestamp:", msg.Timestamp)
			// mailMessage
			mailMessage := models.MailMessage{}
			json.Unmarshal(msg.Body, &mailMessage)

			// msg.Ack(multiAck)
			mail.SendEmail(mailMessage, config)
		}
	}(connection)

	select {}
}

func envConfig() (config models.Config, err error) {
	viper.SetConfigName("environment")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
