package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ChatTag represents chat tags
type ChatTag struct {
	ID        string                `json:"id" gorm:"primaryKey"`
	Name      string                `json:"name" gorm:"column:name;uniqueIndex;not null"`
	Color     string                `json:"color" gorm:"column:color;default:'#007bff'"`
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (c *ChatTag) TableName() string {
	return "chat_tags"
}

func (c *ChatTag) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
