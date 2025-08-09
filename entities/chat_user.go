package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ChatUser represents chat user (anonymous or logged-in from OSS)
type ChatUser struct {
	ID          string                `json:"id" gorm:"primaryKey"`
	BrowserUUID sql.NullString        `json:"browser_uuid" gorm:"null;column:browser_uuid;uniqueIndex"` // For anonymous users
	OSSUserID   sql.NullString        `json:"oss_user_id" gorm:"null;column:oss_user_id"`               // For logged-in OSS users
	Email       sql.NullString        `json:"email" gorm:"null;column:email"`                           // For logged-in users
	IsAnonymous bool                  `json:"is_anonymous" gorm:"column:is_anonymous;default:true"`
	IPAddress   string                `json:"ip_address" gorm:"column:ip_address;not null"`
	UserAgent   sql.NullString        `json:"user_agent" gorm:"null;column:user_agent"`
	CreatedAt   time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (c *ChatUser) TableName() string {
	return "chat_users"
}

func (c *ChatUser) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
