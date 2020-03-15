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

func (m *Object) FetchBooks(eventID uint32) ([]Book) {
	s := []

	return s
}