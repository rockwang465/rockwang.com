package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db username password address etc .. information
type DBInfo struct {
	user     string
	password string
	ip       string
	port     int
	database string
	driver   string
}

// define table structure
type Demo struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"` // 字段必须大写，否则无法创建字段
	Password  string `gorm:"size:255;not null"`
	Telephone string `gorm:"type:varchar(11);unique;not null"`
}

// Init database
func InitDB() (*gorm.DB, error) {
	// connect database
	// db, err := gorm.Open("mysql", "root:UVlY88m9suHLsthK@tcp(10.151.3.79:6446)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	// 这里可以做成配置文件来读取配置
	dbInfo := DBInfo{
		user:     "root",
		password: "UVlY88m9suHLsthK",
		ip:       "10.151.3.79",
		port:     30446,
		database: "go",
		driver:   "mysql",
	}
	dbInfoStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbInfo.user, dbInfo.password, dbInfo.ip, dbInfo.port, dbInfo.database)
	DB, err := gorm.Open(dbInfo.driver, dbInfoStr)
	if err != nil {
		return nil, err
	}

	// enable singular table name and create table
	DB.SingularTable(true)
	DB.AutoMigrate(&Demo{})

	return DB, nil
}

// close db connection
func DBClose(DB *gorm.DB) {
	DB.Close()
}
