package main

import (
	"net/http"
)

// TODO:
// get /v1/todolist
// post /v1/todolist
// get /v1/todolist/id
// put /v1/todolist/id
// delete /v1/todolist/id

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/tasks", app.listTasksHandler)
	mux.HandleFunc("POST /v1/tasks", app.createTaskHandler)
	mux.HandleFunc("GET /v1/tasks/{id}", app.getTaskHandler)
	mux.HandleFunc("PUT /v1/tasks/{id}", app.updateTaskHandler)
	mux.HandleFunc("DELETE /v1/tasks/{id}", app.deleteTaskHandler)

	return app.recoverPanic(mux)
}
