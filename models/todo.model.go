package models

import "github.com/attapon-th/null"

const (
	fooTableName string = "Todo"
	fooSchema    string = "public"
	fooTable     string = fooSchema + "." + fooTableName
)

// Todo Demo model
type Todo struct {
	ModelID
	Name        null.String `json:"name" gorm:"column:name;type:varchar(100)"`
	StatusID    null.Int    `json:"status_id" gorm:"column:status_id;type:int;"`
	Comment     null.String `json:"comment" gorm:"column:comment;type:varchar(255)"`
	ComplatedAt null.Time   `json:"complated_at" gorm:"column:complated_at;type:timestamptz"`
	Tags        null.String `json:"tags" gorm:"column:tags;type:varchar(255)"`
	ModelRecordStamp
}

// TableName get table name
func (m Todo) TableName() string {
	return fooTable
}
