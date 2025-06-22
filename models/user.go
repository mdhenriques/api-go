package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // Inclui automaticamente os campos ID, CreatedAt, UpdatedAt, DeletedAt (opcional)

	Username       string `gorm:"not null" json:"username"`
	Email          string `gorm:"unique;not null" json:"email"`
	HashedPassword string `gorm:"not null" json:"hashed_password"`
	Xp             int    `gorm:"default:0" json:"xp"`
	Coins          int    `gorm:"default:0" json:"coins"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Xp        int    `json:"xp"`
	Coins     int    `json:"coins"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
