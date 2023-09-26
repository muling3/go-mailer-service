package models

type Config struct {
	EmailHost     string `mapstructure: "EMAILHOST"`
	EmailPort     int    `mapstructure: "EMAILPORT"`
	EmailUser     string `mapstructure: "EMAILUSER"`
	EmailPassword string `mapstructure: "EMAILPASSWORD"`
	BrokerUrl     string `mapstructure: "BROKERURL"`
	QueueName     string `mapstructure: "QUEUENAME"`
	Topic         string `mapstructure: "TOPIC"`
	RoutingKey    string `mapstructure: "ROUTINGKEY"`
	MongoUri      string `mapstructure: "MONGOURI"`
}

type MailMessage struct {
	FromName    string `json:"fromName"`
	ToName      string `json:"toName"`
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	Body        string `json:"body"`
	Subject     string `json:"subject"`
}

type SmtpServer struct {
	Host string
	Port string
}
