run-container:
	docker run -e BROKERURL=amqp://rabbitBroker -e EMAILHOST=smtp.gmail.com -e EMAILUSER=mulingealexmuli@gmail.com -e EMAILPASSWORD=iizqdvpqpyoipxpn -e EMAILPORT=587 --network ecommerce mailer
build-image:
	docker build -t mailer .