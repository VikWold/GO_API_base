package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) htmxHandler(w http.ResponseWriter, r *http.Request) {
	err := app.tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		app.logger.Error("template error", "error", err)
	}
}

func (app *application) clickedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div>Items loaded via htmx at %s</div>", time.Now().Format("15:04:05"))
}
