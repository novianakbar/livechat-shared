package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// TicketCategory represents ticket classification
type TicketCategory struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"column:name;size:100;not null"`
	Code        string `json:"code" gorm:"column:code;size:20;uniqueIndex;not null"` // TECH, BILLING, GENERAL
	Description string `json:"description" gorm:"column:description;type:text"`
	Color       string `json:"color" gorm:"column:color;size:7"` // HEX color
	IsActive    bool   `json:"is_active" gorm:"column:is_active;default:true"`

	// SLA Settings (in minutes)
	SLAFirstResponse int `json:"sla_first_response" gorm:"column:sla_first_response;default:60"` // 1 hour
	SLAResolution    int `json:"sla_resolution" gorm:"column:sla_resolution;default:1440"`       // 24 hours

	// Auto Assignment
	DefaultDepartmentID sql.NullString `json:"default_department_id" gorm:"null;column:default_department_id"`
	DefaultDepartment   *Department    `json:"default_department,omitempty" gorm:"foreignKey:DefaultDepartmentID;references:ID"`

	// Timestamps
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`

	// Relations
	Tickets []Ticket `json:"tickets,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
}

func (tc *TicketCategory) TableName() string {
	return "ticket_categories"
}

func (tc *TicketCategory) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
