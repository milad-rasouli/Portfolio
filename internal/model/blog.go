package model

import "time"

type Blog struct {
	Title      string
	Body       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
