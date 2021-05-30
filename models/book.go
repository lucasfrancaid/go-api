package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string
	Author    string
	Price     float32 `gorm:"type:decimal(10,2);"`
	Published time.Time
}
