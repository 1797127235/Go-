package model

import "time"

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:50;unique;not null" json:"username"`
	Password string `json:"-"`
	Role uint `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
