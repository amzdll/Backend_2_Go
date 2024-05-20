package main

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/app"
)

//	@title			Chat API
//	@version		1.0
//	@description	API Server for Chat Application

// @host		127.0.0.1:8000
// @BasePath	/
func main() {
	ctx := context.Background()
	application := app.CreateApp()
	err := application.Start(ctx)
	if err != nil {
		return
	}
}
