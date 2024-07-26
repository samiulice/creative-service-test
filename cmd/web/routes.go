package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(SessionLoad)

	mux.Get("/", app.Home)

	//404 not found route
	mux.NotFound(app.PageNotFound)

	//Public file server
	
	publicFileServer := http.FileServer(http.Dir("./static/public"))
	mux.Handle("/static/public/*", http.StripPrefix("/static/public", publicFileServer))

	
	return mux
}
