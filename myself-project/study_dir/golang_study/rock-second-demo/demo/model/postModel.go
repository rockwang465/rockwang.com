package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID         uuid.UUID            `json:"id"grom:"primary_key;type:char(36)"`
	UserId     uint                 `json:"user_id"gorm:"not null"`
	CategoryId uint                 `json:"category_id"gorm:"not null"`
	Category   *ArticleCategoryInfo // 展示ArticleCategoryInfo表信息给用户时需要
	Title      string               `json:"title"gorm:"not null;type:varchar(30)"`
	HeadImg    string               `json:"head_img"gorm:"not null"`
	Content    string               `json:"content"gorm:"not null;type:text"`
	CreatedAt  LocalTime            `json:"created_at"gorm:"type:timestamp"`
	UpdatedAt  LocalTime            `json:"updated_at"gorm:"type:timestamp"`
}

// https://studygolang.com/articles/25636
// http://gorm.book.jasperxu.com/changelog.html
// 在调用之前，先生成uuid
func (p *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4()) // 将uuid.NewV4生成的值赋值给ID
}
