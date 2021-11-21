package teletype

import "time"

type Article struct {
	ID          int64     `json:"id"`
	URI         string    `json:"uri"`
	Text        string    `json:"text,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
}

type articles struct {
	Length   int64     `json:"length"`
	Articles []Article `json:"articles"`
}
