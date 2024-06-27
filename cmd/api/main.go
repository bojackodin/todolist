package main

import (
	"flag"
	"log"
	"log/slog"
)

type config struct {
	port int
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")

	flag.Parse()

	app := application{
		config: cfg,
	}

	slog.Info("starting server")

	err := app.run()
	if err != nil {
		log.Fatal(err)
	}
}
