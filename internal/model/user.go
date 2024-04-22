package model

import "time"

type User struct {
	ID         int64
	FullName   string
	Email      string
	Password   string
	IsGithub   int64
	OnlineAt   time.Time
	CreatedAt  time.Time
	ModifiedAt time.Time
}
