package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // Incluye ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"not null;size:255"`
	Age        int    `gorm:"not null"`
}
