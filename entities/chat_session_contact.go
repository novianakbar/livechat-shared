package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ChatSessionContact represents contact information for a chat session
type ChatSessionContact struct {
	ID           string                `json:"id" gorm:"primaryKey"`
	SessionID    string                `json:"session_id" gorm:"column:session_id;not null;uniqueIndex"`
	Session      ChatSession           `json:"session" gorm:"foreignKey:SessionID;references:ID"`
	ContactName  string                `json:"contact_name" gorm:"column:contact_name;not null"`
	ContactEmail string                `json:"contact_email" gorm:"column:contact_email;not null"`
	ContactPhone sql.NullString        `json:"contact_phone" gorm:"null;column:contact_phone"`
	Position     sql.NullString        `json:"position" gorm:"null;column:position"` // Job position
	CompanyName  sql.NullString        `json:"company_name" gorm:"null;column:company_name"`
	CreatedAt    time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt    time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt    soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (c *ChatSessionContact) TableName() string {
	return "chat_session_contacts"
}

func (c *ChatSessionContact) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
