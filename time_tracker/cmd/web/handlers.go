package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"test.com/pkg/models"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// запрос ко внешнему API
	passportNumber := r.URL.Query().Get("passportNumber")

	parts := strings.SplitN(passportNumber, " ", 2)

	passportS, err := strconv.Atoi(parts[0])
	if err != nil {
		app.badRequest(w)
		return
	}

	passportN, err := strconv.Atoi(parts[1])
	if err != nil {
		app.badRequest(w)
		return
	}

	client := &http.Client{}

	url := fmt.Sprintf("http://127.0.0.1:5000/info?passportNumber=%d&passportSerie=%d", passportN, passportS)

	body, err := SendRequest(client, url)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var data *models.People
	err = json.Unmarshal(body, &data)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//

	err = app.userTasks.Create(data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Succes"))
}

func (app *application) startTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil || user_id < 1 {
		app.notFound(w)
		return
	}

	task_id, err := strconv.Atoi(r.URL.Query().Get("task_id"))
	if err != nil || task_id < 1 {
		app.notFound(w)
		return
	}

	err = app.userTasks.StartTask(user_id, task_id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("User id = ", user_id, " started task_id = ", task_id)
	w.Write([]byte("Started"))
}

func (app *application) endTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil || user_id < 1 {
		app.notFound(w)
		return
	}

	err = app.userTasks.EndTask(user_id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("User id = ", user_id, " complete his task")
	w.Write([]byte("Started"))
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

	err = app.userTasks.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("User id = ", id, " deleted")
	w.Write([]byte("Deleted"))
}

func SendRequest(client *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Time Tracker API")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, models.ErrNoRecord
	}

	return body, nil
}
