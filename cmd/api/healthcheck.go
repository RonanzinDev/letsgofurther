package main

import (
	"fmt"
	"net/http"
)

// Handler witch writes a plain text response
func (app *application) healtcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status :available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
