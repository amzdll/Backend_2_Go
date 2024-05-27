package main

import (
	"context"
	"github.com/amzdll/shop/internal/app"
	"github.com/rs/zerolog/log"
)

//	@title			Chat API
//	@version		1.0
//	@description	Shop Api
//
// @host		127.0.0.1:8000
// @BasePath	/
func main() {
	ctx := context.Background()
	application := app.CreateApp()
	err := application.Start(ctx)
	if err != nil {
		log.Err(err)
		return
	}
}
