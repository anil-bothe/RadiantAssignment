package db

type AutherBooks struct {
	Author_id Authors `json:"author_id"`
	Book_id   Books   `json:"book_id"`
}
