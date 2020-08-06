package main

import (
	"database/sql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 定义模型
type User1 struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"` // 指定类型；唯一索引
	Role         string  `gorm:"size:255"`                       // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"`                // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                 // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`                     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                              // 忽略本字段
}

type User2 struct {
	gorm.Model
	Name sql.NullString
	Age  sql.NullInt64
}

func main() {
	// 连接Mysql数据库
	db, err := gorm.Open("mysql", "root:UVlY88m9suHLsthK@tcp(10.151.3.79:6446)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		glog.Fatal(err)
	}
	defer db.Close()

	// 创建数据库
	db.AutoMigrate(&User1{})
	db.AutoMigrate(&User2{})

	// 设置空值的默认值
	user := User2{Name: sql.NullString{"", true}, Age: sql.NullInt64{18, true}}
	db.Create(&user)

}
