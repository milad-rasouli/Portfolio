package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type AboutMe struct {
	Content string `json:"content"`
}

func (am AboutMe) Validate() error {
	return validation.ValidateStruct(&am,
		validation.Field(&am.Content, validation.Required, validation.Length(50, 10000)),
	)
}
