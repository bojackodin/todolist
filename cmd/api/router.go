package main

import (
	"net/http"
)

// TODO:
// get /v1/task
// post /v1/task
// get /v1/task/id
// put /v1/task/id
// put /v1/task/id/done
// delete /v1/task/id

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthcheck", app.healthCheckHandler)

	handle := func(pattern string, next func(w http.ResponseWriter, r *http.Request) error) {
		mux.Handle(pattern, app.errorHandler(next))
	}

	handle("GET /v1/tasks", app.listTasksHandler)
	handle("POST /v1/tasks", app.createTaskHandler)
	handle("GET /v1/tasks/{id}", app.getTaskHandler)
	handle("PUT /v1/tasks/{id}", app.updateTaskHandler)
	handle("PUT /v1/tasks/{id}/done", app.markTaskDoneHandler)
	handle("DELETE /v1/tasks/{id}", app.deleteTaskHandler)

	return app.recoverPanic(mux)
}
