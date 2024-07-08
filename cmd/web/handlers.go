package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"test.com/pkg/models"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var data models.User
	err = json.Unmarshal(body, &data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// err = app.commands.Insert(r.Context(), data.Title, data.Content)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Succes"))
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	err = app.users.Delete(r.Context(), id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("User id = ", id, " deleted")
	w.Write([]byte("Deleted"))
}
