package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TicketSLA represents SLA tracking for tickets
type TicketSLA struct {
	ID       string `json:"id" gorm:"primaryKey"`
	TicketID string `json:"ticket_id" gorm:"column:ticket_id;uniqueIndex;not null"`

	// SLA Deadlines
	FirstResponseDue time.Time    `json:"first_response_due" gorm:"column:first_response_due"`
	ResolutionDue    time.Time    `json:"resolution_due" gorm:"column:resolution_due"`
	FirstResponseAt  sql.NullTime `json:"first_response_at" gorm:"null;column:first_response_at"`
	ResolvedAt       sql.NullTime `json:"resolved_at" gorm:"null;column:resolved_at"`

	// SLA Status
	FirstResponseBreached bool `json:"first_response_breached" gorm:"column:first_response_breached;default:false"`
	ResolutionBreached    bool `json:"resolution_breached" gorm:"column:resolution_breached;default:false"`

	// Calculations (in minutes)
	FirstResponseTime int `json:"first_response_time" gorm:"column:first_response_time"` // Actual time taken
	ResolutionTime    int `json:"resolution_time" gorm:"column:resolution_time"`         // Actual time taken

	// Timestamps
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	// Relations
	Ticket *Ticket `json:"ticket,omitempty" gorm:"foreignKey:TicketID;references:ID"`
}

func (ts *TicketSLA) TableName() string {
	return "ticket_sla"
}

func (ts *TicketSLA) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
