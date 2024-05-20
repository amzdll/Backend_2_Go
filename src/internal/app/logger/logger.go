package logger

import (
	"github.com/amzdll/backend_2_go/src/internal/config"
	"log/slog"
	"os"
)

func setupLogger(cfg *config.LogConfig) *slog.Logger {
	var log *slog.Logger

	switch cfg.Stage {
	case config.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
