package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// Department represents department for agents
type Department struct {
	ID          string                `json:"id" gorm:"primaryKey"`
	Name        string                `json:"name" gorm:"column:name;not null"`
	Description sql.NullString        `json:"description" gorm:"null;column:description"`
	IsActive    bool                  `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt   time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (d *Department) TableName() string {
	return "departments"
}

func (d *Department) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
