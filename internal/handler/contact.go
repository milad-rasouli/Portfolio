package handler

import (
	"github.com/Milad75Rasouli/portfolio/internal/request"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type Contact struct {
	Logger *zap.Logger
}

func (h *Contact) GetContact(c fiber.Ctx) error {
	// var message string
	status := c.Query("status")
	h.Logger.Info("status is " + status)
	// message = "I got your message. I will reply it soon!"

	return c.Render("contact", fiber.Map{"status": status})
}

func (h *Contact) PostContact(c fiber.Ctx) error {
	// save the message and valid the information
	var (
		err            error
		contactRequest request.Contact
	)
	contactRequest.Subject = c.Params("subject")
	contactRequest.Email = c.Params("email")
	contactRequest.Message = c.Params("message")

	err = contactRequest.Validate()
	if err != nil {
		return c.Redirect().To("/contact?status=3")
	}
	// save to the database

	return c.Redirect().To("/contact?status=1")
}

func (h *Contact) Register(g fiber.Router) {
	g.Get("/", h.GetContact)
	g.Post("/", h.PostContact)
}
