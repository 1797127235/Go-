package model

import "time"

type Category struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:50;unique;not null" json:"name"`
	Slug string `gorm:"size:50;unique;not null" json:"slug"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}