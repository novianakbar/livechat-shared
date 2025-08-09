package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ChatLog represents chat activity log
type ChatLog struct {
	ID        string                `json:"id" gorm:"primaryKey"`
	SessionID string                `json:"session_id" gorm:"column:session_id;not null"`
	Session   ChatSession           `json:"session" gorm:"foreignKey:SessionID;references:ID"`
	Action    string                `json:"action" gorm:"column:action;not null"` // started, waiting, response, closed, transferred
	Details   sql.NullString        `json:"details" gorm:"null;column:details"`
	UserID    sql.NullString        `json:"user_id" gorm:"null;column:user_id"`
	User      *User                 `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (c *ChatLog) TableName() string {
	return "chat_logs"
}

func (c *ChatLog) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
