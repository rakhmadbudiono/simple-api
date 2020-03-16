package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakhmadbudiono/simple-api/internal/book"
)

type BookAPI struct{}

var bookAPI BookAPI

func (m *BookAPI) setup(r *mux.Router) {
	r.HandleFunc("/", bookAPI.publishBook).Methods("POST")
	r.HandleFunc("/", bookAPI.getBooks).Methods("GET")
}

func (m BookAPI) getBooks(w http.ResponseWriter, r *http.Request) {
	b := book.New()

	books := b.FetchBooks()

	handleJSONResponse(w, books)
}

func (m BookAPI) publishBook(w http.ResponseWriter, r *http.Request) {	
	var b book.Book
	json.NewDecoder(r.Body).Decode(&b)

	book.New().PublishBook(b)

	handleJSONResponse(w, b)
}