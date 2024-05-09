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
type Category struct {
	ID   int64
	Name string
}

type Relation struct {
	CategoryID int64
	PostID     int64
}
type BlogWithCategory struct {
	Blog
	Category []Category
}
