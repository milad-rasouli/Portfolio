package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Blog struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content"`
}

func (b Blog) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Title, validation.Required, validation.Length(3, 100)),
		validation.Field(&b.Content, validation.Required, validation.Length(100, 10000)),
	)
}
