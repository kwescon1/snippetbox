package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const METHOD_NOT_ALLOWED = 405
const INTERNAL_SERVER_ERROR = 500

//change the signature of the home handler so it is defined as a method against
// *applications
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w) // use the notFound() helper

		return

	}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and
	// use the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())

		app.serverError(w, err) // use the serverError helper

		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents
	// dynamic data that we want tp pass in, which for now we'll leave as nil.

	err = ts.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())

		http.Error(w, "Internal Server Error", INTERNAL_SERVER_ERROR)
	}
}

// Change the signature of the snippet handler so it is defined as a method
// against *application.
func (app *application) snippet(w http.ResponseWriter, r *http.Request) {

	// extract the value of the id parameter from the query string and try to parse
	// convert it to an integer using the strconv.Atoi() function. If it can't parse
	// be converted to an integer, or the value is less than 1, we return a 404
	// not found response

	// w.Write([]byte("Displaying Snippet from here"))
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.

		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our resp
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Change the signature of the createSnippet handler so it is defined as a method
// against *application.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	//use r.Method to check whether the request is using post or not
	// if its not, use the w.writeHeader() method to send a 405 status code and
	//the w.Write() method to write a "Method Not Allowed" response body. We
	//then return from the function so that the subsequent code us not executed

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(METHOD_NOT_ALLOWED)
		// w.Write([]byte("Method Not Allowed"))

		app.clientError(w, http.StatusMethodNotAllowed) // use the clientError() helper

		return

	}

	// create some variables holding dummy data. We'll remove these later on during the build
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	//pass the data to the SnippetModel.store() method, receiving the ID of the new record back
	id, err := app.snippets.Store(title, content, expires)

	if err != nil {
		app.serverError(w, err)

		return
	}

	//Redirect user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
