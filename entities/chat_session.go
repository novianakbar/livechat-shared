package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ChatSession represents chat session
type ChatSession struct {
	ID           string                `json:"id" gorm:"primaryKey"`
	ChatUserID   string                `json:"chat_user_id" gorm:"column:chat_user_id;not null"`
	ChatUser     ChatUser              `json:"chat_user" gorm:"foreignKey:ChatUserID;references:ID"`
	AgentID      sql.NullString        `json:"agent_id" gorm:"null;column:agent_id"`
	Agent        *User                 `json:"agent,omitempty" gorm:"foreignKey:AgentID;references:ID"`
	DepartmentID sql.NullString        `json:"department_id" gorm:"null;column:department_id"`
	Department   *Department           `json:"department,omitempty" gorm:"foreignKey:DepartmentID;references:ID"`
	Topic        string                `json:"topic" gorm:"column:topic;not null"`
	Status       string                `json:"status" gorm:"column:status;not null;default:'waiting'"` // waiting, queued, active, closed
	Priority     string                `json:"priority" gorm:"column:priority;default:'normal'"`       // low, normal, high, urgent
	StartedAt    time.Time             `json:"started_at" gorm:"column:started_at"`
	EndedAt      sql.NullTime          `json:"ended_at" gorm:"null;column:ended_at"`
	Messages     []ChatMessage         `json:"messages,omitempty" gorm:"foreignKey:SessionID"`
	Contact      *ChatSessionContact   `json:"contact,omitempty" gorm:"foreignKey:SessionID"`
	CreatedAt    time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt    time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt    soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (c *ChatSession) TableName() string {
	return "chat_sessions"
}

func (c *ChatSession) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
