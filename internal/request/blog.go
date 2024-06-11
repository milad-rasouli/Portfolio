package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Blog struct {
	Title   string `json:"title,omitempty"`
	Body    string `json:"body"`
	Caption string `json:"caption"`
}

func (b Blog) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Title, validation.Required, validation.Length(3, 100)),
		validation.Field(&b.Body, validation.Required, validation.Length(100, 300000)),
		validation.Field(&b.Caption, validation.Required, validation.Length(10, 350)),
	)
}
