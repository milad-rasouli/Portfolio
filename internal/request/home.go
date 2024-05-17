package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Home struct {
	Name       string `json:"name"`
	Slogan     string `json:"slogan"`
	ShortIntro string `json:"short_intro"`
	GithubUrl  string `json:"github_url"`
}

func (h Home) Validate() error {
	return validation.ValidateStruct(&h,
		validation.Field(&h.Name, validation.Length(3, 100), is.ASCII),
		validation.Field(&h.Slogan, validation.Length(10, 250), is.ASCII),
		validation.Field(&h.ShortIntro, validation.Length(10, 250), is.ASCII),
		validation.Field(&h.GithubUrl, is.URL),
	)
}
