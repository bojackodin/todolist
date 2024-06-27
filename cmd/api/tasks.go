package main

import (
	"fmt"
	"net/http"
)

func (app *application) listTasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "list tasks")
}

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create task")
}

func (app *application) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readID(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, id)
}

func (app *application) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "update task by id")
}

func (app *application) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete task by id")
}
