package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"test.com/pkg/models"
)

// @Summary Get list of users
// @Description Get a paginated list of users with optional filters
// @Accept  json
// @Produce  json
// @Param   page     query    int     false  "Page number"
// @Param   pageSize query    int     false  "Page size"
// @Param   surname  query    string  false  "Surname"
// @Param   name     query    string  false  "Name"
// @Param   patronymic query  string  false  "Patronymic"
// @Param   address  query    string  false  "Address"
// @Success 200 {object} []models.UserTask
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/list [get]
func (app *application) getList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	filter := make(map[string]interface{})
	if surname := r.URL.Query().Get("surname"); surname != "" {
		filter["surname"] = surname
	}
	if name := r.URL.Query().Get("name"); name != "" {
		filter["name"] = name
	}
	if patronymic := r.URL.Query().Get("patronymic"); patronymic != "" {
		filter["patronymic"] = patronymic
	}
	if address := r.URL.Query().Get("address"); address != "" {
		filter["address"] = address
	}

	users, err := app.userTasks.GetList(filter, page, pageSize)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.render(w, users)
}

// @Summary Get user workloads
// @Description Get workloads for a specific user within a date range
// @Accept  json
// @Produce  json
// @Param   user_id    query    int     true  "User ID"
// @Param   start_date query    string  true  "Start date (YYYY-MM-DD)"
// @Param   end_date   query    string  true  "End date (YYYY-MM-DD)"
// @Success 200 {object} models.TaskWorkload
// @Failure 500 {string} string "Internal Server Error"
// @Failure 400 {string} string "Bad Request"
// @Router /users/workloads [get]
func (app *application) getWorkloads(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	user_id, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		app.badRequest(w)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		app.badRequest(w)
		return
	}

	data, err := app.userTasks.GetUserWorkload(user_id, startDate, endDate)
	if err != nil {
		app.serverError(w, err)
	}

	w.WriteHeader(http.StatusOK)
	app.render(w, data)
}

// @Summary Create a new user
// @Description Create a new user based on passport information
// @Accept  json
// @Produce  json
// @Param   PassportNumber query string true "Passport Number"
// @Success 201 {string} string "Success"
// @Failure 500 {string} string "Internal Server Error"
// @Failure 400 {string} string "Bad Request"
// @Router /users [post]
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// запрос ко внешнему API
	passportNumber := r.URL.Query().Get("PassportNumber")

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

// @Summary Start a task
// @Description Start a specific task for a user
// @Accept  json
// @Produce  json
// @Param   user_id query int true "User ID"
// @Param   task_id query int true "Task ID"
// @Success 200 {string} string "Started"
// @Failure 500 {string} string "Internal Server Error"
// @Failure 400 {string} string "Bad Request"
// @Router /users/tasks/start [post]
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Started"))
}

// @Summary End a task
// @Description End the current task for a user
// @Accept  json
// @Produce  json
// @Param   user_id query int true "User ID"
// @Success 200 {string} string "Completed"
// @Failure 500 {string} string "Internal Server Error"
// @Failure 404 {string} string "Not Found"
// @Router /users/tasks/end [post]
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Completed"))
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Accept  json
// @Produce  json
// @Param   id query int true "User ID"
// @Success 200 {string} string "Deleted"
// @Failure 500 {string} string "Internal Server Error"
// @Failure 404 {string} string "Not Found"
// @Router /users [delete]
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}

// @Summary Update a user
// @Description Update user information
// @Accept  json
// @Produce  json
// @Param   user body models.UserTask true "User data"
// @Success 200
// @Failure 500 {string} string "Internal Server Error"
// @Failure 400 {string} string "Bad Request"
// @Router /users/update [put]
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var user models.UserTask
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		app.badRequest(w)
		return
	}

	if err := app.userTasks.Update(&user); err != nil {
		app.serverError(w, err)
	}

	w.WriteHeader(http.StatusOK)
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
