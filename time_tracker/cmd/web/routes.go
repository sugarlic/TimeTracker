package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", nil)
	// mux.HandleFunc("/user", nil) // Same as /
	// mux.HandleFunc("/user/update", nil)
	// mux.HandleFunc("/user/create", nil)
	// mux.HandleFunc("/user/cd/start", nil)
	mux.HandleFunc("/create", app.createUser)
	mux.HandleFunc("/delete", app.deleteUser)
	return mux
}
