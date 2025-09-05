package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// TicketComment represents comments and responses on tickets
type TicketComment struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	TicketID       string         `json:"ticket_id" gorm:"column:ticket_id;not null;index"`
	UserID         sql.NullString `json:"user_id" gorm:"null;column:user_id"` // null if customer comment
	Content        string         `json:"content" gorm:"column:content;type:text;not null"`
	IsInternal     bool           `json:"is_internal" gorm:"column:is_internal;default:false"` // true = internal note, false = public
	IsFromCustomer bool           `json:"is_from_customer" gorm:"column:is_from_customer;default:false"`

	// Metadata
	AuthorName  string `json:"author_name" gorm:"column:author_name;size:100"`   // Customer name or Agent name
	AuthorEmail string `json:"author_email" gorm:"column:author_email;size:100"` // For customer comments

	// Timestamps
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`

	// Relations
	Ticket *Ticket `json:"ticket,omitempty" gorm:"foreignKey:TicketID;references:ID"`
	User   *User   `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
}

func (tc *TicketComment) TableName() string {
	return "ticket_comments"
}

func (tc *TicketComment) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
