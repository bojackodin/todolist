package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) run() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	slog.Info("msg", "starting server addr", srv.Addr)

	return srv.ListenAndServe()
}
