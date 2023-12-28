package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//Define a new command-line flag with the name 'addr', a default value of
	// and some short help text explaining what the flag controls. The value of
	// flag will be stored in the addr variable at runtime.

	addr := flag.String("addr", "4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You nedd to call this *BEFORE* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If anny error
	// is encountered during parsing the application will be terminated.

	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Creaate a logger for writing error messages in the same way, but use stderr
	// the destination and use the og.Lshortfile flag to include the relevant
	// file name and line number

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Use the http.NewServerMux() function to initialize a new servemux, then
	//register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//register two new handler functions and corresponding url patterns
	mux.HandleFunc("/snippet", snippet)
	mux.HandleFunc("/snippet/create", createSnippet)

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

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e
	// prefix it with the * symbol) before using it.
	// Write messages using the two new loggers, instead of the standard logger

	infoLog.Printf("Starting server on %s", *addr)

	// Use the http.ListenAndServe() function to start a new web server. We pass
	// two parameters: the TCP network address to listen on (in this case ":400
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.
	err := http.ListenAndServe(":4000", mux)
	errorLog.Fatal(err)
}
