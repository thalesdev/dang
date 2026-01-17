package rest

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/thalesdev/dang/internal/application/queries"
	"github.com/thalesdev/dang/internal/interface/rest/dto/response"
)

type ContentController struct {
	findAllContentHandler  *queries.FindAllContentHandler
	findByIdContentHandler *queries.FindByIdContentHandler
}

func NewContentController(app *fiber.App, findAllContentHandler *queries.FindAllContentHandler, findByIdContentHandler *queries.FindByIdContentHandler) *ContentController {
	controller := ContentController{
		findAllContentHandler:  findAllContentHandler,
		findByIdContentHandler: findByIdContentHandler,
	}
	group := app.Group("/content")

	group.Get("/", controller.getAll)
	group.Get("/:id", controller.getById)

	return &controller
}

func (c *ContentController) getAll(ctx fiber.Ctx) error {
	title := ctx.Query("title")

	var query queries.FindAllContentQuery
	if title != "" {
		query = queries.NewFindAllContentQuery(&title)
	}

	items, err := c.findAllContentHandler.Execute(ctx, query)
	if err != nil {
		return err
	}

	response := response.ToContentResponseList(items)
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	ctx.Send(responseBytes)

	return nil
}

func (c *ContentController) getById(ctx fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}
	content, err := c.findByIdContentHandler.Execute(ctx, queries.NewFindByIdContentQuery(id))
	if err != nil {
		return err
	}

	response, err := json.Marshal(response.ToContentResponse(*content))
	if err != nil {
		return err
	}
	return ctx.Send(response)
}
