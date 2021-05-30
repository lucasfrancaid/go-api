package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Price     float32        `gorm:"type:decimal(10,2);" json:"price"`
	Published time.Time      `json:"published"`
	AuthorID  int            `json:"-"`
	Author    Author         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;,References:Name" json:"author"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
