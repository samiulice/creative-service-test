package main

import (
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "home", nil); err != nil {
		app.errorLog.Println(err)
	}
}
func (app *application) PageNotFound(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "page-not-found", nil); err != nil {
		app.errorLog.Println(err)
	}
}
