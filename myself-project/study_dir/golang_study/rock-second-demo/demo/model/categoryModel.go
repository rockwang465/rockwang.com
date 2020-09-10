package model

// 存到数据库的文章分类model
type ArticleCategoryInfo struct {
	ID uint `json:"id"gorm:"primary_key"`
	Name string `json:"name"gorm:"not null;type:varchar(50)"`
	CreatedAt LocalTime `json:"created_at"gorm:"type:timestamp"`
	UpdatedAt LocalTime `json:"updated_at"gorm:"type:timestamp"`
}