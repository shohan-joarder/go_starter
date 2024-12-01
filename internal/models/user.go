package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
	Name      string    `gorm:"size:100;not null" json:"name" validate:"required,min=2,max=100"`
	Email     string    `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`
	Phone     string    `gorm:"size:100;unique;not null" json:"phone" validate:"required"`
	Password  string    `gorm:"size:100;not null" json:"password" validate:"required,min=6"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
