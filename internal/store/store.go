package store

import (
	"context"
	"errors"

	"github.com/Milad75Rasouli/portfolio/internal/model"
)

var DuplicateUserError = errors.New("the user is already exist")
var UserNotFountError = errors.New("could not find the user")
var AboutMeNotFountError = errors.New("could not find about me")
var CannotCreateTableError = errors.New("Cannot create tables")

var BlogNotFoundError = errors.New("could not find the target blog")
var BlogCreateError = errors.New("could not create the blog")
var CategoryNotFoundError = errors.New("could not find the target category")
var CategoryCreateError = errors.New("could not create category")
var CategoryRelationNotFoundError = errors.New("could not find the target category relation")
var CategoryRelationCreateError = errors.New("could not create the category relation")

type User interface {
	CreateUser(context.Context, model.User) (int64, error)
	GetUserByEmail(context.Context, string) (model.User, error)
	GetUserByID(context.Context, int64) (model.User, error)
	GetAllUser(context.Context) ([]model.User, error)
	DeleteUserByID(context.Context, int64) error
	UpdateUserByPasswordFullName(context.Context, int64, string, string) error
}

type Blog interface {
	CreateBlog(context.Context, model.Blog) (int64, error)
	GetBlogByID(context.Context, int64) (model.Blog, error)
	GetAllBlog(context.Context) ([]model.Blog, error)
	DeleteBlogByID(context.Context, int64) error
	UpdateBlogByID(context.Context, model.Blog) error
	CreateCategory(context.Context, model.Category) (int64, error)
	GetCategoryByID(context.Context, int64) (model.Category, error)
	GetAllCategory(context.Context) ([]model.Category, error)
	DeleteCategoryByID(context.Context, int64) error
	UpdateCategoryByID(context.Context, model.Category) error
	CreateCategoryRelation(context.Context, model.Relation) error
	GetCategoryRelationAllByPostID(context.Context, int64) ([]model.Relation, error)
	GetCategoryRelationAllByCategoryID(context.Context, int64) ([]model.Relation, error)
	DeleteCategoryRelationAllByPostID(context.Context, int64) error
	DeleteCategoryRelationAllByCategoryID(context.Context, int64) error
	GetAllPostsWithCategory(context.Context) ([]model.BlogWithCategory, error)
}

type Contact interface {
	CreateContact(context.Context, model.Contact) (int64, error)
	GetContactByID(context.Context, int64) (model.Contact, error)
	GetAllContact(context.Context) ([]model.Contact, error)
	DeleteContactByID(context.Context, int64) error
}
type AboutMe interface {
	UpdateAboutMe(context.Context, model.AboutMe) error
	GetAboutMe(context.Context) (model.AboutMe, error)
}
type Store interface {
	User
	Blog
	AboutMe
	Contact
}
