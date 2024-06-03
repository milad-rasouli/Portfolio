package handler

import (
	"errors"
	"html/template"
	"strconv"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/request"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type ControlPanel struct {
	Logger *zap.Logger
	DB     store.Store
}

func (cp *ControlPanel) GetControlPanel(c fiber.Ctx) error {
	var (
		err     error
		contact []model.Contact
		aboutMe model.AboutMe
		home    model.Home
	)
	{
		contact, err = cp.DB.GetAllContact(c.Context())
		if errors.Is(err, store.ContactNotFountError) == false && err != nil {
			cp.Logger.Error("GetAllContact error", zap.Error(err))
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	{
		home, err = cp.DB.GetHome(c.Context())
		if errors.Is(err, store.HomeNotFountError) == false && err != nil {
			cp.Logger.Error("GetHome error", zap.Error(err))
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	{
		aboutMe, err = cp.DB.GetAboutMe(c.Context())
		if errors.Is(err, store.AboutMeNotFountError) == false && err != nil {
			cp.Logger.Error("AboutMe error", zap.Error(err))
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	return c.Render("control-panel", fiber.Map{
		"contact":        contact,
		"aboutMeContent": template.HTML(aboutMe.Content),
		"home":           home,
	})
}
func (cp *ControlPanel) GetCreateORModifyBlog(c fiber.Ctx) error {
	blogID := c.Params("blogID")
	if blogID == "new" {
		return c.Render("create-modify.blog", fiber.Map{})
	}
	return c.JSON("modify blog " + blogID)
}
func (cp *ControlPanel) PostDeleteBlog(c fiber.Ctx) error {
	data := struct {
		Data string `json:"data"`
	}{}

	err := c.Bind().Body(&data)
	if err != nil {
		cp.Logger.Error("invalid json", zap.Error(err))
		return Message(c, errors.New("unable to delete the Blog"))
	}
	return Message(c, errors.New("delete user "+data.Data))
}

func (cp *ControlPanel) PostCreateBlog(c fiber.Ctx) error {
	var (
		blog request.Blog
		err  error
	)
	{
		c.Bind().Body(&blog)
		err = blog.Validate()
		if err != nil {
			cp.Logger.Error("create post error", zap.Error(err))
		}
		cp.Logger.Info("create post", zap.Any("data", blog))
	}
	{
		dbBlog := model.Blog{
			Title:      blog.Title,
			Body:       blog.Body,
			Caption:    blog.Caption,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}
		_, err = cp.DB.CreateBlog(c.Context(), dbBlog)
		if err != nil {
			cp.Logger.Error("create blog error", zap.Error(err))
			Message(c, err)
		}
		cp.Logger.Info("create blog successfully")
	}
	return Message(c, errors.New("created blog"))
}

func (cp *ControlPanel) PostModifyBlog(c fiber.Ctx) error {
	return c.JSON("modify blog")
}

func (cp *ControlPanel) PostDeleteUser(c fiber.Ctx) error {
	data := struct {
		Data string `json:"data"`
	}{}

	err := c.Bind().Body(&data)
	if err != nil {
		cp.Logger.Error("invalid json", zap.Error(err))
		return Message(c, errors.New("unable to delete the user"))
	}
	return Message(c, errors.New("delete user "+data.Data))
}
func (cp *ControlPanel) PostDeleteContact(c fiber.Ctx) error {
	data := struct {
		Data string `json:"data"`
	}{}

	err := c.Bind().Body(&data)
	if err != nil {
		cp.Logger.Error("invalid json", zap.Error(err))
		return Message(c, errors.New("unable to parse input"))
	}
	id, err := strconv.ParseInt(data.Data, 10, 64)
	if err != nil {
		cp.Logger.Error("invalid id", zap.Error(err))
		return Message(c, errors.New("invalid id for deleting the contact message"))
	}
	err = cp.DB.DeleteContactByID(c.Context(), id)
	if err != nil {
		return Message(c, errors.New("unable to delete the contact message"))
	}
	return Message(c, errors.New("delete contact message "+data.Data))
}

func (cp *ControlPanel) PostModifyHome(c fiber.Ctx) error {
	var (
		homeRequest request.Home
		home        model.Home
		err         error
	)
	{
		err = c.Bind().Body(&homeRequest)
		if err != nil {
			cp.Logger.Error("home parse error", zap.Error(err))
			return Message(c, err)
		}
		err = homeRequest.Validate()
		if err != nil {
			return Message(c, err)
		}
	}
	{
		home.Name = homeRequest.Name
		home.ShortIntro = homeRequest.ShortIntro
		home.Slogan = homeRequest.Slogan
		home.GithubUrl = homeRequest.GithubUrl
		err = cp.DB.UpdateHome(c.Context(), home)
		if err != nil {
			return Message(c, err)
		}
	}
	return Message(c, errors.New("updated home"))
}

func (cp *ControlPanel) PostModifyAboutMe(c fiber.Ctx) error {
	var (
		aboutMeRequest request.AboutMe
		aboutMe        model.AboutMe
		err            error
	)

	{
		err = c.Bind().Body(&aboutMeRequest)
		if err != nil {
			cp.Logger.Error("about me parse error", zap.Error(err))
			return Message(c, err)
		}
		err = aboutMeRequest.Validate()
		if err != nil {
			return Message(c, err)
		}
	}
	{
		aboutMe.Content = aboutMeRequest.Content
		cp.DB.UpdateAboutMe(c.Context(), aboutMe)
	}
	return Message(c, errors.New("updated about-me!"))
}

func (cp *ControlPanel) Register(g fiber.Router) {
	g.Get("/", cp.GetControlPanel)                                 //
	g.Get("/create-modify-blog/:blogID", cp.GetCreateORModifyBlog) //
	g.Post("/delete-blog", cp.PostDeleteBlog)                      //                      //
	g.Post("/create-blog", cp.PostCreateBlog)
	g.Post("/modify-blog", cp.PostModifyBlog)
	g.Post("/delete-user", cp.PostDeleteBlog)       //
	g.Post("/delete-contact", cp.PostDeleteContact) //
	g.Post("/modify-home", cp.PostModifyHome)
	g.Post("/modify-about-me", cp.PostModifyAboutMe)
}
