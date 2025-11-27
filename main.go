package main

import (
	"context"
	"fmt"

	application "github.com/Gswaraw37/go-chi-test.git/app"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}