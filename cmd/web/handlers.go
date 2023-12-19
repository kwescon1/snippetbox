package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const METHOD_NOT_ALLOWED = 405

// define a home handler function which writes a byte slice containing
// "Hello from snippetbox" as the response body.

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return

	}
	w.Write([]byte("Hello from snippetbox"))
}

func snippet(w http.ResponseWriter, r *http.Request) {

	// extract the value of the id parameter from the query string and try to parse
	// convert it to an integer using the strconv.Atoi() function. If it can't parse
	// be converted to an integer, or the value is less than 1, we return a 404
	// not found response

	// w.Write([]byte("Displaying Snippet from here"))
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)

		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our resp
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	//use r.Method to check whether the request is using post or not
	// if its not, use the w.writeHeader() method to send a 405 status code and
	//the w.Write() method to write a "Method Not Allowed" response body. We
	//then return from the function so that the subsequent code us not executed

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(METHOD_NOT_ALLOWED)
		// w.Write([]byte("Method Not Allowed"))

		//use the http.Error() method to send a 405 status code and "Method
		//Allowed" string as the response body.

		http.Error(w, "Method Not Allowed", METHOD_NOT_ALLOWED)

		return

	}

	w.Write([]byte("Creating new snippet"))
}
