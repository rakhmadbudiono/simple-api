package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakhmadbudiono/simple-api/internal/book"
)

type BookAPI struct{}

var bookAPI BookAPI

func (m *BookAPI) setup(r *mux.Router) {
	r.HandleFunc("/", bookAPI.getBooks).Methods("POST")
}

func (m BookAPI) getBooks(w http.ResponseWriter, r *http.Request) {
	ss := book.New()

	books := ss.FetchBooks()

	handleJSONResponse(w, books)
}
