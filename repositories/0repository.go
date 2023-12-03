package repositories

import (
	"context"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// TableModel get table name interface of gorm
type TableModel interface {
	TableName() string
}

// Repository  Base repository
type Repository[T TableModel] struct {
	ctx   context.Context
	db    *gorm.DB
	tx    *gorm.DB
	model *T
}

// NewRepository create new repository for Table Model
func NewRepository[T TableModel](db *gorm.DB, model T) *Repository[T] {
	if viper.GetBool("dev") {
		db = db.Debug()
	}
	return &Repository[T]{
		db:    db,
		tx:    db.Model(&model),
		ctx:   context.Background(),
		model: &model,
	}
}

// Context function set context
func (r *Repository[T]) Context(ctx context.Context) *Repository[T] {
	r.ctx = ctx
	r.tx = r.db.WithContext(ctx).Model(r.model)
	return r
}

// GormDB function get gorm db
func (r *Repository[T]) GormDB() *gorm.DB {
	return r.db
}

// GetTx function get context transaction continue
func (r *Repository[T]) GetTx() *gorm.DB {
	return r.tx
}

// Where function for create WHERE condition same as gorm
func (r *Repository[T]) Where(query string, args ...interface{}) *Repository[T] {
	if r.tx == nil {
		r.tx = r.db.Model(r.model).Where(query, args...)
		return r
	}
	r.tx = r.tx.Where(query, args...)
	return r
}

// Or function for create OR condition
// example:
//
// r.Or(
//
//	[]any{"name = ?", "john"},
//	[]any{"age BETWEEN ? AND ?", 10, 20},
//
// )
func (r *Repository[T]) Or(queryAgrs ...[]any) *Repository[T] {
	if len(queryAgrs) == 0 {
		return r
	}
	var fns []any
	db := r.db
	for _, where := range queryAgrs {
		w := where[0]
		agrs := []any{}
		if len(where) > 1 {
			agrs = where[1:]
		}
		fns = append(fns, db.Where(w, agrs...))
	}
	if len(fns) == 1 {
		r.tx = r.tx.Or(fns[0])
		return r
	}
	r.tx = r.tx.Or(fns[0], fns...)
	return r
}

// Count function get record count without limit in table if continue query data
func (r *Repository[T]) Count() int64 {
	var count int64
	if r.tx == nil {
		r.db.Model(r.model).Count(&count)
		return count
	}
	r.tx = r.tx.Limit(-1).Offset(-1).Count(&count)
	return count
}

// Limit function get limit record in table
// example:
//
// r.Where("name = ?", "john").Limit(10).Find(results)
func (r *Repository[T]) Limit(limit int) *Repository[T] {
	if r.tx == nil {
		r.tx = r.db.Model(r.model).Limit(limit)
		return r
	}
	r.tx = r.tx.Limit(limit)
	return r
}

// Offset function get offset record in table
// example:
//
// r.Where("name = ?", "john").Offset(10).Find(results)
func (r *Repository[T]) Offset(offset int) *Repository[T] {
	if r.tx == nil {
		r.tx = r.db.Model(r.model).Offset(offset)
		return r
	}
	r.tx = r.tx.Offset(offset)
	return r
}

// OrderBy function get order by record in table
// (default: id DESC)
// example:
//
// r.Where("name = ?", "john").Order("id desc").Find(results)
func (r *Repository[T]) OrderBy(order string) *Repository[T] {
	if r.tx == nil {
		r.tx = r.db.Model(r.model).Order(order)
		return r
	}
	r.tx = r.tx.Order(order)
	return r
}

// Find function get all record in table
//
// example: find and get count all record without limit
//
// results := []User{}
// r.Where("name = ?", "john").Limit(10).Find(&results)
//
// count = r.Count()
func (r *Repository[T]) Find(results any) *Repository[T] {
	if r.tx == nil {
		r.db.Model(r.model).Find(results)
		return r
	}
	r.tx = r.tx.Find(results)
	return r
}

// First function get first record in table
// example:
//
// result := User{}
// r.Where("name = ?", "john").First()
func (r *Repository[T]) First(result any) *Repository[T] {
	if r.tx == nil {
		r.db.Model(r.model).First(result)
		return r
	}
	r.tx = r.tx.First(result)
	return r
}

// Create function create record
// example:
//
// r.Create(&User{Name: "john"})
func (r *Repository[T]) Create(model *T) error {
	if r.tx == nil {
		return r.db.Model(model).Create(model).Error
	}
	r.tx = r.tx.Create(model)
	return r.tx.Error
}

// Updates function update record
// example:
//
// r.Where("name = ?", "john").Updates(&User{Name: "johnny"})
func (r *Repository[T]) Updates(model *T) error {
	if r.tx == nil {
		return r.db.Model(model).Updates(model).Error
	}
	r.tx = r.tx.Updates(model)
	return r.tx.Error
}

// Delete function delete record
// example:
//
// r.Delete(1, 2, 3)
func (r *Repository[T]) Delete(primaryKey ...any) error {
	if r.tx == nil {
		return r.db.Model(r.model).Delete(r.model, primaryKey).Error
	}
	r.tx = r.tx.Delete(r.model, primaryKey)
	return r.tx.Error
}

// Err function get error
func (r *Repository[T]) Err() error {
	if r.tx == nil {
		return nil
	}
	return r.tx.Error
}
