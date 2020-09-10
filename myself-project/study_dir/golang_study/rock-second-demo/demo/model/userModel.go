package model

type UserInfo struct {
	//gorm.Model
	ID        uint      `json:"id"gorm:"primary_key"`
	CreatedAt LocalTime `json:"created_at"gorm:"type:timestamp"`
	UpdatedAt LocalTime `json:"created_at"gorm:"type:timestamp"`
	Name      string    `json:"name"gorm:"not null;type:varchar(50)"`
	Telephone string    `json:"telephone"gorm:"unique;type:varchar(11);not null"`
	Password  string    `json:"password"gorm:"size:255"`
}
