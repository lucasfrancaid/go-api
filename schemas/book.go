package schemas

import "time"

type CreateBookSchema struct {
	Title     string    `json:"title"`
	AuthorID  int       `json:"author_id"`
	Price     float32   `json:"price"`
	Published time.Time `json:"published"`
}
