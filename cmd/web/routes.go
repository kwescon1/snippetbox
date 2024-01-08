package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	//Use the http.NewServerMux() function to initialize a new servemux, then
	//register the home function as the handler for the "/" URL pattern.

	mux := http.NewServeMux()

	// Swap the route declarations to use the application struct's methods as the
	// handler functions.
	mux.HandleFunc("/", app.home)

	//register two new handler functions and corresponding url patterns
	mux.HandleFunc("/snippet", app.snippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// create a file server which serves files out of the "./ui/static" directory
	// Note that the path given to the http.Dir function is relative to the program
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// serving single file
	// func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./ui/static/file.zip")
	// }

	return mux
}
