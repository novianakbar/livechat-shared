package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// AgentStatus represents agent login session status
type AgentStatus struct {
	ID          string                `json:"id" gorm:"primaryKey"`
	AgentID     string                `json:"agent_id" gorm:"column:agent_id;not null"`
	Agent       User                  `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	Status      string                `json:"status" gorm:"column:status;not null"` // logged_in, logged_out
	LastLoginAt time.Time             `json:"last_login_at" gorm:"column:last_login_at"`
	CreatedAt   time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (a *AgentStatus) TableName() string {
	return "agent_status"
}

func (a *AgentStatus) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
