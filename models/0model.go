package models

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ModelID primary key use ksuid
type ModelID struct {
	ID ksuid.KSUID `json:"id" gorm:"column:id;type:varchar(27);primary_key"`
}

// ModelRecordStamp record version timestamp and soft delete flag
type ModelRecordStamp struct {
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // use `1` `0`
}

// NewID create new ID
func (m *ModelID) NewID() ksuid.KSUID {
	m.ID = ksuid.New()
	return m.ID
}

// BeforeCreate gorm hook for before create
func (m *ModelID) BeforeCreate(*gorm.DB) (err error) {
	m.NewID()
	return nil
}
