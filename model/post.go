package model

import "time"

type Post struct {
    ID        uint      `gorm:"primaryKey" json:"id"` // ID
    Title     string    `gorm:"size:255;not null" json:"title"` // 标题
    Content   string    `gorm:"type:longtext" json:"content"` // 内容
    AuthorID  uint      `json:"author_id"` // 作者ID
    CategoryID uint `json:"category_id"` // 分类ID
    Status uint `json:"status"` // 状态
    CoverImage string `json:"cover_image"` // 封面图
    Views uint `json:"views"` // 浏览量
    Likes uint `json:"likes"` // 点赞数
    IsTop uint `json:"is_top"` // 是否置顶
    CreatedAt time.Time `json:"created_at"` // 创建时间
    UpdatedAt time.Time `json:"updated_at"` // 更新时间
}