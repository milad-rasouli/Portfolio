package model

import "time"

type Portfolio struct {
	Name string
	Summary string
	Skill string
	Experience string
	Education string
	Projects string
	CreatedAt time.Time
	ModifiedAt time.Time
}
