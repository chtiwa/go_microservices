package main

import (
	"context"
	"fmt"

	"github.com/chtiwa/go_microservices/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failmed to start app")
	}
}
