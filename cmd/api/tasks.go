package main

import (
	"fmt"
	"net/http"
)

func (app *application) listTasksHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "list tasks")
	return nil
}

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "create task")
	return nil
}

func (app *application) getTaskHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := app.readID(r)
	if err != nil {
		http.NotFound(w, r)
		return err
	}
	fmt.Fprint(w, id)
	return nil
}

func (app *application) updateTaskHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "update task by id")
	return nil
}

func (app *application) markTaskDoneHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "mark task done by id")
	return nil
}

func (app *application) deleteTaskHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "delete task by id")
	return nil
}
