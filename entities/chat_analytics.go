package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// ChatAnalytics represents chat analytics
type ChatAnalytics struct {
	ID                  string                `json:"id" gorm:"primaryKey"`
	Date                time.Time             `json:"date" gorm:"column:date;not null"`
	TotalSessions       int                   `json:"total_sessions" gorm:"column:total_sessions;default:0"`
	CompletedSessions   int                   `json:"completed_sessions" gorm:"column:completed_sessions;default:0"`
	AverageResponseTime float64               `json:"average_response_time" gorm:"column:average_response_time;default:0"` // in seconds
	TotalMessages       int                   `json:"total_messages" gorm:"column:total_messages;default:0"`
	DepartmentID        sql.NullString        `json:"department_id" gorm:"null;column:department_id"`
	Department          *Department           `json:"department,omitempty" gorm:"foreignKey:DepartmentID;references:ID"`
	AgentID             sql.NullString        `json:"agent_id" gorm:"null;column:agent_id"`
	Agent               *User                 `json:"agent,omitempty" gorm:"foreignKey:AgentID;references:ID"`
	CreatedAt           time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt           time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt           soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (c *ChatAnalytics) TableName() string {
	return "chat_analytics"
}

func (c *ChatAnalytics) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
