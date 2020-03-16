package book

var books []Book

func (m *Object) FetchBooks() ([]Book) {
	return books
}

func (m *Object) PublishBook(book Book) () {
	books = append(books, book)
}