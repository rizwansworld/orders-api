package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/rizwansworld/orders-api/application"
)

func main() {
	app := application.New()

	// The NotifyContext method returns two things
	// A derived context from a root context which can be used to cancel any ongoing tasks based on signal interrupt.
	// A cancel function which can be called to perform the cancellations.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// defer keyword calls the cancel function at the end of the current function, as our requirement.
	// In our case, end of the main function.
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
