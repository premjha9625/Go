package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// "github.com/milap-neupane/golang-chi-crud-api/books"
	"github.com/milap-neupane/golang-chi-crud-api/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/books", BookRoutes())

	http.ListenAndServe(":3000", r)
}

func BookRoutes() chi.Router {
	r := chi.NewRouter()
	//bookHandler := handlers.BookHandler{}
	r.Get("/", handlers.ListBooks)
	r.Post("/", handlers.CreateBook)
	r.Get("/{id}", handlers.GetBooks)
	r.Put("/{id}", handlers.UpdateBook)
	r.Delete("/{id}", handlers.DeleteBook)
	return r
}
