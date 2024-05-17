package handler

import (
	"errors"
	"fmt"
	"strconv"

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
	// show a report to the users
	var (
		err     error
		contact []model.Contact
	)
	contact, err = cp.DB.GetAllContact(c.Context())
	if err != nil {
		cp.Logger.Error("GetAllContact error", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Render("control-panel", fiber.Map{
		"contact": contact,
	})
}
func (cp *ControlPanel) GetCreateORModifyBlog(c fiber.Ctx) error {
	blogID := c.Params("blogID")
	if blogID == "new" {
		return c.JSON("create blog " + blogID)
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
	return c.JSON("create blog")
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
	return c.JSON("modify home")
}

func (cp *ControlPanel) PostModifyAboutMe(c fiber.Ctx) error {
	var am request.AboutMe
	c.Bind().Body(&am)
	fmt.Printf("about me content %+v", am)
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
