package factory

import (
	"log/slog"
	"os"

	"github.com/OlegLaban/weather_service/pkg/logger"
	"github.com/OlegLaban/weather_service/pkg/logger/handler"
)

type Config struct {
	Env string
}

type factory struct {
	conf Config
}

func New(conf Config) *factory {
	return &factory{conf: conf}
}

func (f *factory) NewLogger(key string) *logger.Logger {
	var slogLogger *slog.Logger
	switch f.conf.Env {
	case "prod":
		opts := handler.PrettyHandlerOptions{
			SlogOpts: slog.HandlerOptions{
				Level: slog.LevelError,
			},
		}
		slogLogger = slog.New(handler.NewPrettyJSONHandler(os.Stdout, opts, key))
	case "dev":
	case "local":
		opts := handler.PrettyHandlerOptions{
			SlogOpts: slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}
		slogLogger = slog.New(handler.NewPrettyTextHandler(os.Stdout, opts, key))
	}
	return logger.New(slogLogger)
}
