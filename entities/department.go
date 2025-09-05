package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// Department represents department for agents
type Department struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"column:name;not null"`
	Description sql.NullString `json:"description" gorm:"null;column:description"`
	IsActive    bool           `json:"is_active" gorm:"column:is_active;default:true"`

	// Ticketing specific fields
	CanHandleTickets   bool `json:"can_handle_tickets" gorm:"column:can_handle_tickets;default:true"`
	MaxTicketsPerAgent int  `json:"max_tickets_per_agent" gorm:"column:max_tickets_per_agent;default:10"`

	// Multi-Level Support Fields
	SupportLevel       int            `json:"support_level" gorm:"column:support_level;default:0"`                           // 0=L0, 1=L1, 2=L2, 3=L3
	ParentDeptID       sql.NullString `json:"parent_dept_id" gorm:"null;column:parent_dept_id"`                              // Parent department for hierarchy
	ParentDept         *Department    `json:"parent_dept,omitempty" gorm:"foreignKey:ParentDeptID;references:ID"`            // Parent department
	MaxEscalationLevel int            `json:"max_escalation_level" gorm:"column:max_escalation_level;default:3"`             // Maximum level this dept can handle
	AutoAssignmentRule string         `json:"auto_assignment_rule" gorm:"column:auto_assignment_rule;default:'round_robin'"` // round_robin, least_loaded, skill_based
	EscalationDeptID   sql.NullString `json:"escalation_dept_id" gorm:"null;column:escalation_dept_id"`                      // Default escalation target department
	EscalationDept     *Department    `json:"escalation_dept,omitempty" gorm:"foreignKey:EscalationDeptID;references:ID"`    // Escalation target department

	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (d *Department) TableName() string {
	return "departments"
}

func (d *Department) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
