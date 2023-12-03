package todoctl

import (
	"strconv"

	"github.com/attapon-th/template-fiber-api/schemas"
	"github.com/attapon-th/template-fiber-api/services/todosrv"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var (
	log zerolog.Logger
)

// NewTodoCtl creates a new todo controller
func NewTodoCtl(r fiber.Router) {
	log = zlog.With().Str("ctl", "todo").Logger()
	r.Route("/", func(r fiber.Router) {
		r.Get("/", gets)         // get list data
		r.Get("/:id", getByID)   // get by id
		r.Post("/", create)      // add data
		r.Put("/:id", update)    // update all data
		r.Patch("/:id", patch)   // update some data
		r.Delete("/:id", delete) // soft delete by default
	}).Name("todo")

	log.Info().Msg("New Todo Contoller initialized")
}

// @Summary		Show todos
// @Description	get todos
// @Tags			todos
// @Accept			json
// @Param			id			query	string	false	"Filter  ID"
// @Param			name		query	string	false	"Filter name"
// @Param			status_id	query	int		false	"Todo id status"
// @Param			_page		query	int		false	"Page number of Todo Data"	default(1)
// @Param			_limit		query	int		false	"Limit Todo Data"			default(10)
// @Produce		json
// @Success		200		{object}	schemas.Todos
// @Failure		default	{object}	schemas.Todos	"not success"
// @Router			/todos [get]
func gets(c *fiber.Ctx) error {

	sv := todosrv.NewTodoService(c.UserContext())

	q := c.Queries()
	limit, _ := strconv.ParseInt(c.Query("_limit", "10"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("_page", "1"), 10, 64)

	result := sv.Gets(limit, page, q)
	if result.Code != 200 {
		result.Data = nil
	}
	return c.Status(result.Code).JSON(result)

}

// @Summary		Show todo by id
// @Description	get todos by id
// @Tags			todos
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Todo ID"
// @Success		200		{object}	schemas.TodoOne
// @Failure		default	{object}	schemas.TodoOne
// @Router			/todos/{id} [get]
func getByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) != 27 {
		return fiber.NewError(400, "invalid id")
	}

	sv := todosrv.NewTodoService(c.UserContext())
	result := sv.GetByID(id)
	return c.JSON(result)
}

// @Summary		Create Todo
// @Description	create todo
// @Tags			todos
// @Accept			json
// @Produce		json
// @Param			request	body		schemas.TodoItem	true	"Todo Item"
// @Success		200		{object}	schemas.TodoOne
// @Failure		default	{object}	schemas.TodoOne
// @Router			/todos [post]
func create(c *fiber.Ctx) error {
	data := schemas.TodoItem{}

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(400, "can't parse body, "+err.Error())
	}
	sv := todosrv.NewTodoService(c.UserContext())

	res := sv.Create(&data)

	return c.Status(res.Code).JSON(res)
}

// @Summary		Update Todo
// @Description	Update todo by ID
// @Tags			todos
// @Accept			json
// @Produce		json
// @Param			request	body		schemas.TodoItem	true	"Todo Item"
// @Param			id		path		int					true	"Todo ID"
// @Success		200		{object}	schemas.TodoOne
// @Failure		default	{object}	schemas.TodoOne
// @Router			/todos/{id} [put]
func update(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) != 27 {
		return fiber.NewError(400, "invalid id")
	}
	data := schemas.TodoItem{}

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(400, "can't parse body, "+err.Error())
	}
	sv := todosrv.NewTodoService(c.UserContext())

	result := sv.Update(id, &data)

	return c.Status(result.Code).JSON(result)
}

// @Summary		Update Todo some data
// @Description	Update todo by ID
// @Tags			todos
// @Accept			json
// @Produce		json
// @Param			request	body		schemas.TodoItem	true	"Todo Item"
// @Param			id		path		int					true	"Todo ID"
// @Success		200		{object}	schemas.TodoOne
// @Failure		default	{object}	schemas.TodoOne
// @Router			/todos/{id} [patch]
func patch(c *fiber.Ctx) error {
	// same update
	return update(c)
}

// @Summary		Delete Todo by ID
// @Description	Update todo by ID
// @Tags			todos
// @Accept			json
// @Produce		json
// @Param			id		path		int	true	"Todo ID"
// @Success		200		{object}	schemas.TodoOne
// @Failure		default	{object}	schemas.TodoOne
// @Router			/todos/{id} [delete]
func delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) != 27 {
		return fiber.NewError(400, "invalid id")
	}
	sv := todosrv.NewTodoService(c.UserContext())
	result := sv.Delete(id)
	return c.Status(result.Code).JSON(result.APIResponse)
}
