package book

type Object struct {
	bo *Book
}

type Book struct {
	BookId	 		int64 		`json:"id_book"`
	Title   		string  	`json:"title"`
}

func New() *Object {
	return &Object{}
}

func (m *Object) FetchBooks() ([]Book) {
	s := []Book{Book{1, "Buku 1"}, Book{2, "Buku 2"}}

	return s
}