package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (app *application) readID(r *http.Request) (int64, error) {
	idString := r.PathValue("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id")
	}

	return id, nil
}
