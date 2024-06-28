package main

import (
	"flag"
	"log/slog"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	app := application{
		config: cfg,
	}

	err := app.run()
	if err != nil {
		slog.Error("server run", err)
		os.Exit(1)
	}
}
