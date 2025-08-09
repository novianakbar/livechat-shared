package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// User represents system user (admin, agent, etc.)
type User struct {
	ID           string                `json:"id" gorm:"primaryKey"`
	Email        string                `json:"email" gorm:"column:email;uniqueIndex;not null"`
	Password     string                `json:"-" gorm:"column:password;not null"`
	Name         string                `json:"name" gorm:"column:name;not null"`
	Role         string                `json:"role" gorm:"column:role;not null"` // admin, agent
	IsActive     bool                  `json:"is_active" gorm:"column:is_active;default:true"`
	DepartmentID sql.NullString        `json:"department_id" gorm:"null;column:department_id"`
	Department   *Department           `json:"department,omitempty" gorm:"foreignKey:DepartmentID;references:ID"`
	CreatedAt    time.Time             `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt    time.Time             `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt    soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete:second;"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	uuidV7, _ := uuid.NewV7()
	uuid := uuidV7.String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
