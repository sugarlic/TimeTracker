package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/list", app.getList)
	mux.HandleFunc("/users/workloads", app.getWorkloads)
	mux.HandleFunc("/users/tasks/start", app.startTask)
	mux.HandleFunc("/users/tasks/end", app.endTask)
	mux.HandleFunc("/users", app.createUser)
	mux.HandleFunc("/users/", app.deleteUser)
	return mux
}
