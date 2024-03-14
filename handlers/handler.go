// Handlers for the router yet to be implemented
package handlers

import (
	"encoding/json"
	"net/http"

	//"./../models/models"
	"github.com/milap-neupane/golang-chi-crud-api/models"
)

type BookHandler struct {
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.ListBooks())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
