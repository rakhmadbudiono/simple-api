package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rakhmadbudiono/simple-api/internal/book"
)

type BookAPI struct{}

var bookAPI BookAPI

func (m *BookAPI) setup(r *mux.Router) {
	r.HandleFunc("/", bookAPI.getBooks).Methods("GET")
}

func (m BookAPI) getBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if err != nil {
		handleError(w, NewErrorNoMessage(400))
		return
	}

	ss := book.New()

	books, err := ss.FetchBooks()
	if err != nil {
		handleError(w, NewInternalServerError(err))
		return
	}

	handleJSONResponse(w, books)
}
