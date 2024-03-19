// Handlers for the router yet to be implemented
package books

import (
	// "encoding/json"
	// "net/http"

	//"./../models/models"
	"github.com/milap-neupane/golang-chi-crud-api/models"
)

type BookHandler struct {
}

func GetBooks(id string) *models.Book {
	for _, book := range models.Books {
		if book.ID == id {
			return book
		}
	}
	return nil
}

func StoreBook(book models.Book) {
	models.Books = append(models.Books, &book)
}

func DeleteBook(id string) *models.Book {
	for i, book := range models.Books {
		if book.ID == id {
			models.Books = append(models.Books[:i], (models.Books)[i+1:]...)
			return &models.Book{}
		}
	}
	return nil
}

func UpdateBook(id string, bookUpdate models.Book) *models.Book {
	for i, book := range models.Books {
		if book.ID == id {
			models.Books[i] = &bookUpdate
			return book
		}
	}
	return nil
}

// func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
// 	err := json.NewEncoder(w).Encode(models.ListBooks())
// 	if err != nil {
// 		http.Error(w, "Internal error", http.StatusInternalServerError)
// 		return
// 	}
// }
// func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request)   {}
// func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {}
// func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {}
// func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {}
