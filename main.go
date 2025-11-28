package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	application "github.com/Gswaraw37/go-chi-test.git/app"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}