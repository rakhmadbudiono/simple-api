package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type API struct {
	book BookAPI
}

var api API

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", welcome)

	api := API{}

	api.book.setup(r.PathPrefix("/book").Subrouter())

	return r
}

type timestampResponse struct {
	Timestamp int64 `json:"timestamp"`
}

func welcome(w http.ResponseWriter, r *http.Request) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)

	data := timestampResponse{
		Timestamp: millis,
	}

	handleJSONResponse(w, data)
}
