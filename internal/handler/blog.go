package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Milad75Rasouli/portfolio/frontend/views/pages"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type Blog struct {
	Logger    *zap.Logger
	BlogStore store.Blog
}

func (b *Blog) GetList(c fiber.Ctx) error {
	var (
		err       error
		blog      []model.Blog
		sBlogList string
	)
	{
		blog, err = b.BlogStore.GetAllBlog(c.Context())
		if errors.Is(err, store.BlogNotFoundError) == false && err != nil {
			b.Logger.Error("GetAllBlog error", zap.Error(err))
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	for i := 0; i < len(blog); i++ {
		sBlogList += fmt.Sprintf(`<div class="col">
		  <div class="card bg-light text-dark border-light">
			<div class="card-body">
			  <h5 class="card-title">%s</h5>
			  <p class="card-text">%s</p>
			  <a href="/blog/%d" class="btn btn-dark float-start">Read</a>
			</div>
			<div class="card-footer text-dark border-light">
			  <div class="float-start text-small">
				%s
			  </div> 
			</div>
		  </div>
		</div>`, blog[i].Title, blog[i].Caption, blog[i].ID, blog[i].ModifiedAt.Format("2006/01/02"))
	}
	base := pages.BlogList(sBlogList)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	err = base.Render(context.Background(), c.Response().BodyWriter()) // Pass the blogs here
	if err != nil {
		b.Logger.Error("tempel render error", zap.Error(err))
		return Message(c, err)
	}
	return c.SendStatus(fiber.StatusOK)
}

func (b *Blog) GetBlog(c fiber.Ctx) error {
	var (
		// role  = c.Locals("userRole")
		// email = c.Locals("userEmail")
		blogID = c.Params("blogID")
		blog   model.Blog
		err    error
	)
	{
		err = validation.Validate(blogID,
			validation.Required,
			is.Int,
		)
		if err != nil {
			b.Logger.Error("get blog validation error", zap.Error(err))
			return Message(c, err)
		}
	}
	{
		// e, _ := email.(string)
		// r, _ := role.(string)
		id, err := strconv.ParseInt(blogID, 10, 64)
		if err != nil {
			b.Logger.Error("get blog parse error", zap.Error(err))
			return err
		}
		blog, err = b.BlogStore.GetBlogByID(c.Context(), id)
		if err != nil {
			b.Logger.Error("get blog database error", zap.Error(err))
			return Message(c, err)
		}
		base := pages.Blog(blog.Title, blog.Body)
		base.Render(context.Background(), c.Response().BodyWriter())
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	}
	return c.SendStatus(fiber.StatusOK)
}

func (b *Blog) Register(g fiber.Router) {
	g.Get("/", b.GetList)
	g.Get("/:blogID", b.GetBlog)
}
