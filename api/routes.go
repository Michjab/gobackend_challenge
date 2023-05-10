package main

import "net/http"

// Function initiates mux http server with routing REST paths for given application struct configuration. Returns http server instance
func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", app.EchoHandler)
	mux.HandleFunc("/sum", app.SumHandler)
	mux.HandleFunc("/multiply", app.MultiplyHandler)
	mux.HandleFunc("/flatten", app.FlattenHandler)
	mux.HandleFunc("/invert", app.TranspositionHandler)

	return mux
}
