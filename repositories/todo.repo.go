package repositories

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/wiriyamanuwong/template-fiber-api/models"
	"github.com/wiriyamanuwong/template-fiber-api/pkg"
)

const todoFields string = "id,name,status_id,comment,complated_at,tags"

var todoRepo = (*TodoRepository)(nil)

// TodoRepository todo repository
type TodoRepository struct {
	*Repository[models.Todo]
	Fields []string
}

func NewTodoRepository() *TodoRepository {
	if todoRepo != nil {
		return todoRepo
	}
	td := &TodoRepository{}
	db := pkg.ConnectPostgreSQL(viper.GetString("DB_DSN"))
	td.Repository = NewRepository(db, models.Todo{})
	td.Fields = strings.Split(todoFields, ",")
	todoRepo = td
	return td
}
