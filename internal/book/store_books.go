package book

import (
	"github.com/rakhmadbudiono/simple-api/util/pubsub"
)

func (m *Object) FetchBooks() ([]Book) {
	return pubsub.books
}

func (m *Object) PublishBooks() (int) {
	return 1
}