package request

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserSingUp struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var UserSingUpPasswordRegex = regexp.MustCompile(`\w+\d+`)

func (u UserSingUp) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FullName, validation.Required, validation.Length(3, 100), is.Alphanumeric),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Length(8, 50), is.ASCII, validation.Match(UserSingUpPasswordRegex)),
	)
}
