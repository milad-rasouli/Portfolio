package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Contact struct {
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (c Contact) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Subject, validation.Required, validation.Length(3, 100)),
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Message, validation.Length(5, 500)),
	)
}
