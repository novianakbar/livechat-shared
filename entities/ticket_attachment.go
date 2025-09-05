package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// TicketAttachment represents file attachments on tickets
type TicketAttachment struct {
	ID         string `json:"id" gorm:"primaryKey"`
	TicketID   string `json:"ticket_id" gorm:"column:ticket_id;not null;index"`
	FileName   string `json:"file_name" gorm:"column:file_name;size:255;not null"`
	FilePath   string `json:"file_path" gorm:"column:file_path;size:500;not null"`
	FileSize   int64  `json:"file_size" gorm:"column:file_size"`
	FileType   string `json:"file_type" gorm:"column:file_type;size:50"`
	UploadedBy string `json:"uploaded_by" gorm:"column:uploaded_by;size:20"` // customer, agent

	// Timestamps
	CreatedAt time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`

	// Relations
	Ticket *Ticket `json:"ticket,omitempty" gorm:"foreignKey:TicketID;references:ID"`
}

func (ta *TicketAttachment) TableName() string {
	return "ticket_attachments"
}

func (ta *TicketAttachment) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
