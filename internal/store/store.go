package store

import (
	"context"
	"errors"

	"github.com/Milad75Rasouli/portfolio/internal/model"
)

var DuplicateUserError = errors.New("the user is already exist")
var UserNotFountError = errors.New("could not find the user")
var CannotCreateTableError = errors.New("Cannot create tables")

var BlogNotFoundError = errors.New("could not find the target blog")
var BlogCreateError = errors.New("could not create the blog")
var CategoryNotFoundError = errors.New("could not find the target category")
var CategoryCreateError = errors.New("could not create category")
var CategoryRelationNotFoundError = errors.New("could not find the target category relation")
var CategoryRelationCreateError = errors.New("could not create the category relation")

type User interface {
	Create(context.Context, model.User) (int64, error)
	GetByEmail(context.Context, string) (model.User, error)
	GetByID(context.Context, int64) (model.User, error)
	GetAll(context.Context) ([]model.User, error)
	DeleteByID(context.Context, int64) error
	UpdatePasswordFullName(context.Context, int64, string, string) error
}

type Blog interface {
	CreateBlog(context.Context, model.Blog) (int64, error)
	GetByID(context.Context, int64) (model.Blog, error)
	GetAll(context.Context) ([]model.Blog, error)
	DeleteByID(context.Context, int64) error
}
type Store interface {
	// User
	// Blog
}
