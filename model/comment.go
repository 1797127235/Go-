package model

import "time"

type Comment struct {
	ID uint `gorm:"primaryKey" json:"id"`//评论ID
	UserID uint `json:"user_id"` //评论者ID
	PostID uint `json:"post_id"` //评论哪篇文章
	Content string `json:"content"` //评论内容
	ParentID uint `json:"parent_id"` //回复哪条评论
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}