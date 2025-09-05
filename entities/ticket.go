package entities

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// Ticket represents a support ticket
type Ticket struct {
	// Basic Info
	ID          string `json:"id" gorm:"primaryKey"`
	TicketCode  string `json:"ticket_code" gorm:"column:ticket_code;uniqueIndex;size:20;not null"` // TKT-20250809-001
	Subject     string `json:"subject" gorm:"column:subject;size:255;not null"`
	Description string `json:"description" gorm:"column:description;type:text"`

	// Customer Info
	CustomerName  string `json:"customer_name" gorm:"column:customer_name;size:100"`
	CustomerEmail string `json:"customer_email" gorm:"column:customer_email;size:100"`
	CustomerPhone string `json:"customer_phone" gorm:"column:customer_phone;size:20"`

	// Classification
	CategoryID sql.NullString  `json:"category_id" gorm:"null;column:category_id"`
	Category   *TicketCategory `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	Priority   string          `json:"priority" gorm:"column:priority;size:20;default:'medium'"` // low, medium, high, urgent
	Status     string          `json:"status" gorm:"column:status;size:20;default:'open'"`       // open, in_progress, resolved, closed, escalated

	// Assignment
	AssignedToID sql.NullString `json:"assigned_to_id" gorm:"null;column:assigned_to_id"`
	AssignedTo   *User          `json:"assigned_to,omitempty" gorm:"foreignKey:AssignedToID;references:ID"`
	DepartmentID sql.NullString `json:"department_id" gorm:"null;column:department_id"`
	Department   *Department    `json:"department,omitempty" gorm:"foreignKey:DepartmentID;references:ID"`

	// Multi-Level Escalation Fields
	CurrentLevel    int    `json:"current_level" gorm:"column:current_level;default:0"`       // 0=L0, 1=L1, 2=L2, 3=L3
	EscalationPath  string `json:"escalation_path" gorm:"column:escalation_path;type:text"`   // JSON history: ["L0_dept_id", "L1_dept_id"]
	CanEscalate     bool   `json:"can_escalate" gorm:"column:can_escalate;default:true"`      // Can this ticket be escalated
	MaxLevel        int    `json:"max_level" gorm:"column:max_level;default:3"`               // Maximum escalation level (0-3)
	EscalationCount int    `json:"escalation_count" gorm:"column:escalation_count;default:0"` // Number of times escalated

	// Previous assignment tracking for escalation
	PreviousAssignedToID sql.NullString `json:"previous_assigned_to_id" gorm:"null;column:previous_assigned_to_id"`
	PreviousAssignedTo   *User          `json:"previous_assigned_to,omitempty" gorm:"foreignKey:PreviousAssignedToID;references:ID"`
	PreviousDepartmentID sql.NullString `json:"previous_department_id" gorm:"null;column:previous_department_id"`
	PreviousDepartment   *Department    `json:"previous_department,omitempty" gorm:"foreignKey:PreviousDepartmentID;references:ID"`

	// SLA & Tracking
	CreatedByID     sql.NullString `json:"created_by_id" gorm:"null;column:created_by_id"` // Agent ID who created (null if customer/AI)
	CreatedBy       *User          `json:"created_by,omitempty" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedVia      string         `json:"created_via" gorm:"column:created_via;size:20"` // customer, agent, ai
	FirstResponseAt sql.NullTime   `json:"first_response_at" gorm:"null;column:first_response_at"`
	ResolvedAt      sql.NullTime   `json:"resolved_at" gorm:"null;column:resolved_at"`
	ClosedAt        sql.NullTime   `json:"closed_at" gorm:"null;column:closed_at"`

	// Security & Access
	AccessToken string `json:"access_token" gorm:"column:access_token;size:64;index"` // For customer access

	// Timestamps
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`

	// Relations
	Comments    []TicketComment    `json:"comments,omitempty" gorm:"foreignKey:TicketID;references:ID"`
	Attachments []TicketAttachment `json:"attachments,omitempty" gorm:"foreignKey:TicketID;references:ID"`
	History     []TicketHistory    `json:"history,omitempty" gorm:"foreignKey:TicketID;references:ID"`
	SLA         *TicketSLA         `json:"sla,omitempty" gorm:"foreignKey:TicketID;references:ID"`
}

func (t *Ticket) TableName() string {
	return "tickets"
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)

	// Generate ticket code if not provided
	if t.TicketCode == "" {
		ticketCode := t.generateTicketCode()
		tx.Statement.SetColumn("TicketCode", ticketCode)
	}

	// Generate unique access token
	accessToken := t.generateAccessToken()
	tx.Statement.SetColumn("AccessToken", accessToken)

	return nil
}

// generateTicketCode creates a unique ticket code
func (t *Ticket) generateTicketCode() string {
	now := time.Now()
	return fmt.Sprintf("TKT-%s-%s",
		now.Format("20060102"),
		now.Format("150405")) // HHMMSS
}

// generateAccessToken creates a secure access token for customer portal
func (t *Ticket) generateAccessToken() string {
	tokenUUID, _ := uuid.NewV7()
	return tokenUUID.String()[:32] // Use first 32 chars for token
}
