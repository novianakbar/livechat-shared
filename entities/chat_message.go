package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ChatMessage represents chat message
type ChatMessage struct {
	ID          string                `json:"id" gorm:"primaryKey"`
	SessionID   string                `json:"session_id" gorm:"column:session_id;not null"`
	Session     ChatSession           `json:"session" gorm:"foreignKey:SessionID;references:ID"`
	SenderID    sql.NullString        `json:"sender_id" gorm:"null;column:sender_id"`
	Sender      *User                 `json:"sender,omitempty" gorm:"foreignKey:SenderID;references:ID"`
	SenderType  string                `json:"sender_type" gorm:"column:sender_type;not null"` // customer, agent, system
	Message     string                `json:"message" gorm:"column:message;not null"`
	MessageType string                `json:"message_type" gorm:"column:message_type;default:'text'"` // text, image, file, system
	Attachments []string              `json:"attachments" gorm:"column:attachments;type:json"`
	ReadAt      sql.NullTime          `json:"read_at" gorm:"null;column:read_at"`
	CreatedAt   time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (c *ChatMessage) TableName() string {
	return "chat_messages"
}

func (c *ChatMessage) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
