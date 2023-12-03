package todosrv

import (
	"context"
	"time"

	"github.com/attapon-th/null"
	"github.com/attapon-th/template-fiber-api/models"
	"github.com/attapon-th/template-fiber-api/repositories"
	"github.com/attapon-th/template-fiber-api/schemas"
	"github.com/rs/zerolog/log"
)

type TodoService struct {
	ctx      context.Context
	todoRepo *repositories.TodoRepository
}

func NewTodoService(ctx context.Context) *TodoService {
	return &TodoService{
		ctx:      ctx,
		todoRepo: repositories.NewTodoRepository(),
	}
}

func (s *TodoService) Context(ctx context.Context) *TodoService {
	s.ctx = ctx
	return s
}

func (s *TodoService) Create(data *schemas.TodoItem) *schemas.TodoOne {
	r := s.todoRepo
	result := schemas.NewTodoOne()
	// convert schema to model
	md := models.Todo{
		Name:        data.Name,
		Comment:     data.Comment,
		StatusID:    data.StatusID,
		ComplatedAt: null.NewTime(time.Time{}, false),
		Tags:        data.Tags,
	}

	err := r.Context(s.ctx).Create(&md)
	if err != nil {
		log.Error().Err(err).Msg("create todo error")
		result.APIResponse = *schemas.NewAPIResponse(500, "Create Todo Error, "+err.Error())
		return result
	}

	// if succsss get data by id
	result = s.GetByID(md.ID.String())
	result.Message = "Create Todo Success"
	return result
}

func (s *TodoService) Update(id string, data *schemas.TodoItem) *schemas.TodoOne {
	r := s.todoRepo

	// check id if exists and no error or deleted
	result := s.GetByID(id)
	if result.Code == 404 {
		return result
	} else if result.Code != 200 {
		return result
	}

	md := models.Todo{
		Name:        data.Name,
		Comment:     data.Comment,
		StatusID:    data.StatusID,
		ComplatedAt: data.ComplatedAt,
		Tags:        data.Tags,
	}
	if err := md.ID.Scan(id); err != nil {
		result.APIResponse = *schemas.NewAPIResponse(500, "Update Todo Error, id not valid")
		return result
	}

	err := r.Context(s.ctx).Where("id = ?", id).Updates(&md)
	if err != nil {
		result.APIResponse = *schemas.NewAPIResponse(500, "Update Todo Error, "+err.Error())
		return result
	}

	result = s.GetByID(id)
	result.Message = "Update Todo Success"
	return result

}

func (s *TodoService) Delete(id string) *schemas.TodoOne {
	r := s.todoRepo
	result := schemas.NewTodoOne()
	if len(id) != 27 {
		result.APIResponse = *schemas.NewAPIResponse(400, "invalid id")
		return result
	}
	err := r.Context(s.ctx).Where("id = ?", id).Delete()
	if err != nil {
		log.Error().Err(err).Msg("delete todo error")
		result.APIResponse = *schemas.NewAPIResponse(500, "Delete Todo Error")
		return result
	}
	result.Data.ID = null.StringFrom(id)
	result.APIResponse = *schemas.NewAPIResponse(200, "Delete Todo Success")
	return result
}
