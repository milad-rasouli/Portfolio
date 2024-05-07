package model

import "time"

type Blog struct {
	ID         int64
	Title      string
	Body       string
	Caption    string
	ImagePath  string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
