package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	RoleID    uint      `gorm:"not null" json:"role_id"`
	Name      string    `gorm:"size:100;not null" json:"name" validate:"required,min=2,max=100"`
	Email     string    `gorm:"size:100;unique;not null" json:"email" validate:"required,email,unique=users_email"`
	Phone     string    `gorm:"size:100;unique;not null" json:"phone" validate:"required,min=10,max=15,unique=users_phone"`
	Password  string    `gorm:"size:100;not null" json:"password" validate:"required,min=6"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
