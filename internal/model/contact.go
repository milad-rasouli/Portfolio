package model

import "time"

type Contact struct {
	ID        int64
	Subject   string
	Email     string
	Message   string
	CreatedAt time.Time
}
