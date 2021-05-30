package schemas

import "time"

type CreateBookSchema struct {
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Price     float32   `json:"price"`
	Published time.Time `json:"published"`
}
