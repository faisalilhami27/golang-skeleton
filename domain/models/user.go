package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int
	Name      string
	Username  string
	Password  string
	Photo     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
