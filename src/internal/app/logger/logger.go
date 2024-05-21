package logger

import (
	"github.com/amzdll/backend_2_go/src/internal/config"
	"github.com/rs/zerolog"
	"os"
)

func setupLogger(cfg *config.LogConfig) *zerolog.Logger {
	var logger zerolog.Logger
	switch cfg.Stage {
	case config.EnvLocal:
		logger = setupLocalLogger()
	case config.EnvDev:
		logger = setupDevLogger()
	case config.EnvProd:
		logger = setupProdLogger()
	}
	return &logger
}

func setupLocalLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05 02/01/2006"}
	return zerolog.New(output).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func setupDevLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func setupProdLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().Timestamp().Logger()
}
