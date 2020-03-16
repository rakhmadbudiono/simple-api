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
	r.HandleFunc("/", bookAPI.publishBooks).Methods("POST")
	r.HandleFunc("/", bookAPI.getBooks).Methods("GET")
}

func (m BookAPI) getBooks(w http.ResponseWriter, r *http.Request) {
	b := book.New()

	books := b.FetchBooks()

	handleJSONResponse(w, books)
}

func (m BookAPI) publishBooks(w http.ResponseWriter, r *http.Request) {	
	b := book.New()
	json.NewDecoder(r.Body).Decode(&b)

	b.PublishBooks()
	// if success == nil {
	// 	// handleError(w, NewInternalServerError(err))
	// 	return
	// }

	handleJSONResponse(w, r.Body)
}
