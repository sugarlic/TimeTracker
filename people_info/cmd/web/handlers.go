package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"people_info.com/pkg/models"
	_ "people_info.com/pkg/models/postgre"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.users.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.render(w, s)
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	passportSerie, err := strconv.Atoi(r.URL.Query().Get("passportSerie"))
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w)
		return
	}

	passportNumber, err := strconv.Atoi(r.URL.Query().Get("passportNumber"))
	if err != nil {
		app.errorLog.Println(err)
		app.badRequest(w)
		return
	}

	s, err := app.users.Get(passportSerie, passportNumber)
	if err != nil {
		app.serverError(w, err)
		return
	}

	people := &models.People{
		Name:       s.Name,
		Surname:    s.Surname,
		Patronymic: s.Patronymic,
		Address:    s.Address,
	}

	app.render(w, people)
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

	err = app.users.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("User id = ", id, " deleted")
	w.Write([]byte("Deleted"))
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.badRequest(w)
		return
	}

	var data models.User
	err = json.Unmarshal(body, &data)
	if err != nil {
		app.badRequest(w)
		return
	}

	err = app.users.Insert(data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Succes"))
}
