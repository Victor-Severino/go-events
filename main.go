package main

import (
	"github.com/Victor-Severino/go-events/event"
	"github.com/Victor-Severino/go-events/user"
)

func main() {
	ed := event.NewEventDispatcher()
	sendEmailListener := user.NewSendEmailListener()
	publishRabbitMQListener := user.NewPublishRabbitMQListener()

	ed.AddListener("user_created", sendEmailListener)
	ed.AddListener("user_created", publishRabbitMQListener)

	user := user.NewUser(ed)
	user.Create("User1")
}