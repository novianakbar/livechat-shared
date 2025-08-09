package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ChatSessionTag represents many-to-many relationship between sessions and tags
type ChatSessionTag struct {
	ID        string                `json:"id" gorm:"primaryKey"`
	SessionID string                `json:"session_id" gorm:"column:session_id;not null"`
	Session   ChatSession           `json:"session" gorm:"foreignKey:SessionID;references:ID"`
	TagID     string                `json:"tag_id" gorm:"column:tag_id;not null"`
	Tag       ChatTag               `json:"tag" gorm:"foreignKey:TagID;references:ID"`
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (c *ChatSessionTag) TableName() string {
	return "chat_session_tags"
}

func (c *ChatSessionTag) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
