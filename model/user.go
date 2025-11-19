package model

import "time"

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:50;unique;not null" json:"username"`
	Password string `json:"-"`
	Avatar string `gorm:"size:255" json:"avatar"` //存放用户头像的url
	Role uint `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
