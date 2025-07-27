package main

import (
	"fmt"
	"os"

	"github.com/Lexantes/api-url-shortener/internal/config"

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

	fmt.Println(cfg)


	// TODO: logger: log/slog

	// todo: init storage: sqlite3

	// todo: init router: chi, "chi render"

	// todo: run server:
}


