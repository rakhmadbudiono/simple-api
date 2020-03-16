package book

type Object struct {
	book *Book
}

type Book struct {
	BookId	 		int64 		`json:"id_book"`
	Title   		string  	`json:"title"`
}

func New() *Object {
	return &Object{}
}