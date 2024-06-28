package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version: %s\n", version)
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
}
