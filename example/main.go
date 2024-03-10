package main

import (
	"context"
	"log"

	"github.com/hexennacht/tunda"
)

type User struct {
	ID                   int
	Username             string
	Password             string
	ConfirmationPassword string
}

func main() {
	t := tunda.NewTundaClient()

	user := User{
		Username: "hexennacht",
		Password: "superSecretP4$W0RD",
		ConfirmationPassword: "superSecretP4$W0RD",
	}

	kerjaan := tunda.Kerjaan{
		Data:         user,
		TimeDuration: 5,
	}

	id, err := t.Create(context.Background(), &kerjaan)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("Success adding to queue with ID: %s \n", id)
}
