package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Lexantes/api-url-shortener/internal/config"
	"github.com/Lexantes/api-url-shortener/internal/lib/logger/sl"
	"github.com/Lexantes/api-url-shortener/internal/storage/sqlite"

	"github.com/jessevdk/go-flags"
)

func main() {

	if _, err := config.Parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	cfg := config.MustLoad(config.Opt.Config)
	log := setupLogger(cfg.Env)

	log.Info("Запускаю url-shortener", slog.String("env", cfg.Env))

	fmt.Println(cfg)

	storage, err := sqlite.New(cfg.StoragePath)

	if err != nil {
		log.Error("Ошибка инициализация storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// todo: init router: chi, "chi render"

	// todo: run server:
}

const (
	envLocal = "local"
	// envDev   = "dev"
	envProd = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	// case envDev:
	// 	log = slog.New(
	// 		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		panic("Непонятное окружение. Проверьте конфиг")
	}

	return log

}
