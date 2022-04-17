package entity

import "time"

type Model struct {
	ID        uint       `gorm:"primary:key"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"created_at"`
	DeletedAt *time.Time `sql:"index"`
}
