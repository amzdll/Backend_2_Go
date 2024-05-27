package logger

import (
	"github.com/amzdll/shop/internal/config"
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	logger *zerolog.Logger
}

func NewLogger(cfg *config.LogConfig) *Logger {
	var l zerolog.Logger
	switch cfg.Stage {
	case config.EnvLocal:
		l = setupLocalLogger()
	case config.EnvDev:

		l = setupDevLogger()
	case config.EnvProd:
		l = setupProdLogger()
	}
	l.Info().Msg("Logger initialized")
	return &Logger{logger: &l}
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Err(err).Msg(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Fatal().Err(err).Msg(msg)
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
