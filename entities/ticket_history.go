package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TicketHistory represents audit trail for ticket changes
type TicketHistory struct {
	ID       string         `json:"id" gorm:"primaryKey"`
	TicketID string         `json:"ticket_id" gorm:"column:ticket_id;not null;index"`
	UserID   sql.NullString `json:"user_id" gorm:"null;column:user_id"`           // null if system action
	Action   string         `json:"action" gorm:"column:action;size:50;not null"` // created, status_changed, assigned, commented, etc

	// Change Details
	FieldName string `json:"field_name" gorm:"column:field_name;size:50"`
	OldValue  string `json:"old_value" gorm:"column:old_value;size:255"`
	NewValue  string `json:"new_value" gorm:"column:new_value;size:255"`

	// Additional Info
	Description string `json:"description" gorm:"column:description;type:text"`
	ActorName   string `json:"actor_name" gorm:"column:actor_name;size:100"` // Who did the action

	// Timestamps
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`

	// Relations
	Ticket *Ticket `json:"ticket,omitempty" gorm:"foreignKey:TicketID;references:ID"`
	User   *User   `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
}

func (th *TicketHistory) TableName() string {
	return "ticket_history"
}

func (th *TicketHistory) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
