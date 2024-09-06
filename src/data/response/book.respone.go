package response

import (
	"github.com/google/uuid"
)

type BookResponse struct {
	BookId      uuid.UUID `json:"bookId"`
	Title       string    `json:"title"`
	ISBN        string    `json:"isbn"`
	Publication string    `json:"publication"`
	Author      string    `json:"author"`
	Genre       string    `json:"genre"`
}
