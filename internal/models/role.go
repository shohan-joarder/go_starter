package models

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name" validate:"required,min=2,max=100"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// type HasPermissions struct {
// 	ID           uint      `gorm:"primaryKey" json:"id"`
// 	RoleID       uint      `gorm:"not null" json:"role_id"`
// 	PermissionID uint      `gorm:"not null" json:"permission_id"`
// 	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
// 	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
// }
