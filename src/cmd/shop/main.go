package main

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/app"
)

func main() {
	ctx := context.Background()
	application := app.CreateApp()
	err := application.Start(ctx)
	if err != nil {
		return
	}
}
