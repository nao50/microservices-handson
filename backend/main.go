package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	repo := NewTODORepository(NewDB())
	// app := NewTODOApplication(repo)
	handler := NewTODOHandler(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/todos", handler.ListTODOs)
	r.Get("/todos/{id}", handler.GetTODO)
	r.Post("/todos", handler.CreateTODO)
	r.Put("/todos/{id}", handler.UpdateTODO)
	r.Delete("/todos/{id}", handler.DeleteTODO)

	http.ListenAndServe(":5050", r)
}
