package models

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name" validate:"required,min=2,max=100,unique=roles_name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type RolePermission struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	RoleID      uint      `gorm:"not null" json:"role_id"`
	Permissions string    `gorm:"type:json" json:"permissions"`   // Store as JSON string
	CreatedBy   *uint     `gorm:"default:null" json:"created_by"` // Nullable
	UpdatedBy   *uint     `gorm:"default:null" json:"updated_by"` // Nullable
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Permissions []Permission

type Permission struct {
	URL     string `json:"url"`
	Methods string `json:"methods"`
	Status  bool   `json:"status"`
}
