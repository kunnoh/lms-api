package response

import "github.com/google/uuid"

type BookResponse struct {
	BookId uuid.UUID `json:"bookId"`
	Title  string    `json:"title"`
}
