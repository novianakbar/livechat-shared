package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// TicketEscalation represents escalation history and tracking
type TicketEscalation struct {
	ID       string  `json:"id" gorm:"primaryKey"`
	TicketID string  `json:"ticket_id" gorm:"column:ticket_id;not null;index"`
	Ticket   *Ticket `json:"ticket,omitempty" gorm:"foreignKey:TicketID;references:ID"`

	// Escalation Level Tracking
	FromLevel int `json:"from_level" gorm:"column:from_level;not null"` // Previous level (0=L0, 1=L1, etc)
	ToLevel   int `json:"to_level" gorm:"column:to_level;not null"`     // New level

	// Department and Agent Changes
	FromDepartmentID sql.NullString `json:"from_department_id" gorm:"null;column:from_department_id"`
	FromDepartment   *Department    `json:"from_department,omitempty" gorm:"foreignKey:FromDepartmentID;references:ID"`
	ToDepartmentID   sql.NullString `json:"to_department_id" gorm:"null;column:to_department_id"`
	ToDepartment     *Department    `json:"to_department,omitempty" gorm:"foreignKey:ToDepartmentID;references:ID"`

	FromAgentID sql.NullString `json:"from_agent_id" gorm:"null;column:from_agent_id"`
	FromAgent   *User          `json:"from_agent,omitempty" gorm:"foreignKey:FromAgentID;references:ID"`
	ToAgentID   sql.NullString `json:"to_agent_id" gorm:"null;column:to_agent_id"`
	ToAgent     *User          `json:"to_agent,omitempty" gorm:"foreignKey:ToAgentID;references:ID"`

	// Escalation Details
	Reason        string    `json:"reason" gorm:"column:reason;type:text;not null"`                       // Reason for escalation
	EscalatedByID string    `json:"escalated_by_id" gorm:"column:escalated_by_id;not null"`               // Who triggered the escalation
	EscalatedBy   *User     `json:"escalated_by,omitempty" gorm:"foreignKey:EscalatedByID;references:ID"` // User who escalated
	EscalatedAt   time.Time `json:"escalated_at" gorm:"column:escalated_at;autoCreateTime;<-:create"`     // When escalated

	// Escalation Type and Trigger
	IsAutoEscalation bool   `json:"is_auto_escalation" gorm:"column:is_auto_escalation;default:false"` // Auto or manual escalation
	TriggerType      string `json:"trigger_type" gorm:"column:trigger_type;size:50"`                   // sla_breach, manual, priority_change, workload
	TriggerData      string `json:"trigger_data" gorm:"column:trigger_data;type:text"`                 // Additional trigger information (JSON)

	// Success tracking
	WasSuccessful bool           `json:"was_successful" gorm:"column:was_successful;default:true"` // Was escalation successful
	FailureReason sql.NullString `json:"failure_reason" gorm:"null;column:failure_reason"`         // If failed, why
	ResolvedAt    sql.NullTime   `json:"resolved_at" gorm:"null;column:resolved_at"`               // When this escalation was resolved

	// Timestamps
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (te *TicketEscalation) TableName() string {
	return "ticket_escalations"
}

func (te *TicketEscalation) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}

// GetLevelName returns human readable level name
func (te *TicketEscalation) GetLevelName(level int) string {
	switch level {
	case 0:
		return "L0 - First Level Support"
	case 1:
		return "L1 - Technical Support"
	case 2:
		return "L2 - Senior Support"
	case 3:
		return "L3 - Expert/Management"
	default:
		return "Unknown Level"
	}
}

// GetFromLevelName returns human readable from level name
func (te *TicketEscalation) GetFromLevelName() string {
	return te.GetLevelName(te.FromLevel)
}

// GetToLevelName returns human readable to level name
func (te *TicketEscalation) GetToLevelName() string {
	return te.GetLevelName(te.ToLevel)
}
