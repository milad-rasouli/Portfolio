package handler

import (
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/request"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

const (
	contactCreatedSuccessfully = "1"
	contactDatabaseError       = "2"
	contactInvalidInputFields  = "3"
)

type Contact struct {
	Logger       *zap.Logger
	ContactStore store.Contact
}

func (contact *Contact) GetContact(c fiber.Ctx) error {
	status := c.Query("status")
	contact.Logger.Info("status is " + status)

	return c.Render("contact", fiber.Map{"status": status})
}

func (contact *Contact) PostContact(c fiber.Ctx) error {
	var (
		err            error
		contactRequest request.Contact
		contactModel   model.Contact
	)
	{
		contactRequest.Subject = c.FormValue("subject")
		contactRequest.Email = c.FormValue("email")
		contactRequest.Message = c.FormValue("message")

		err = contactRequest.Validate()
		if err != nil {
			contact.Logger.Info("invalid contact fields", zap.Error(err))
			return postContactRedirect(c, contactInvalidInputFields)
		}
	}

	{
		contactModel.Subject = contactRequest.Subject
		contactModel.Email = contactRequest.Email
		contactModel.Message = contactRequest.Message
		contactModel.CreatedAt = time.Now()
		_, err = contact.ContactStore.CreateContact(c.Context(), contactModel)
		if err != nil {
			contact.Logger.Error("create contact", zap.Error(err))
			return postContactRedirect(c, contactDatabaseError)
		}
	}

	return postContactRedirect(c, contactCreatedSuccessfully)
}

func (contact *Contact) Register(g fiber.Router) {
	g.Get("/", contact.GetContact)
	g.Post("/", contact.PostContact)
}

func postContactRedirect(c fiber.Ctx, status string) error {
	return c.Redirect().To("/contact?status=" + status)
}
