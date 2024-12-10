package models

import "time"

type Warehouse struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"` // Auto-increment primary key
	UserID      uint      `json:"user_id" validate:"required"`        // Foreign key to a user
	Name        string    `json:"name" validate:"required"`           // Required field
	Phone       string    `json:"phone" validate:"omitempty"`         // Optional, validate as E.164 phone number format
	Address     string    `json:"address" validate:"omitempty"`       // Optional field
	Description string    `json:"description" validate:"omitempty"`   // Optional field
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`   // Automatically managed creation time
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`   // Automatically managed update time
}
