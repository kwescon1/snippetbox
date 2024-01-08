package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom logger
// we will add more to it as the build progresses.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

	// Initialize a new instance of application containing the dependencies

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Initialize a new http.Server struct. We set the Addr and Handler fields
	// that the server uses the same network address and routes as before, and
	// the Errorlog field so that the server now uses the custom reeroLog logged
	// in event of any problems.

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // call the new app.routes() method
	}

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e
	// prefix it with the * symbol) before using it.
	// Write messages using the two new loggers, instead of the standard logger

	infoLog.Printf("Starting server on %s", *addr)

	//Call the listenAndServe() method on our new http.Server struct.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
