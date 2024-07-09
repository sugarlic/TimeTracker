package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/info", app.getUser)
	mux.HandleFunc("/delete", app.deleteUser)
	mux.HandleFunc("/create", app.createUser)
	return mux
}
