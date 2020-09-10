package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// db username password address etc .. information
type DBInfo struct {
	username string
	password string
	ip       string
	port     int
	database string
	driver   string
	charset  string
	loc      string
}

// define user table
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"` // 字段必须大写，否则无法创建字段
	Password  string `gorm:"size:255;not null"`         // 密码要加密的，所以很长
	Telephone string `gorm:"type:varchar(11);unique;not null"`
}

// define article category table 文章分类表
type Category struct {
	ID        uint      `json:"id" gorm:"primary_key"` // ID会自增的
	Name      string    `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt LocalTime `json:"create_at" gorm:"type:timestamp"`
	//CreatedAt time.Time `json:"create_at" gorm:"type:timestamp"`
	UpdatedAt LocalTime `json:"update_at" gorm:"type timestamp"`
	//UpdatedAt time.Time `json:"update_at" gorm:"type timestamp"`
}

// define article info table 单篇文章信息表
type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36); primary_key"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	CategoryID uint      `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string    `json:"title" gorm:"type:char(50); not null"`
	HeadImage  string    `json:"head_image"`
	Content    string    `json:"content" gorm:"type:text;not null"`
	CreatedAt  LocalTime `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  LocalTime `json:"updated_at" gorm:"type:timestamp"`
}

// https://studygolang.com/articles/25636
// http://gorm.book.jasperxu.com/changelog.html
// 在调用之前，先生成uuid
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4()) // 将uuid.NewV4生成的值赋值给 ID
}
