package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type application struct {
}

func main() {
	app := new(application)

	router := chi.NewRouter()
	router.Get("/", app.handleHome)
}

func (app *application) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
