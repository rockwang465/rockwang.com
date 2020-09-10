package model

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"net/url"
)

var DB *gorm.DB

type DBSource struct {
	Driver   string
	Username string
	Password string
	IPAddr   string
	Port     int
	DBName   string
	Charset  string
	Loc      string
}

func ConnectDB() error {
	var dbSource = &DBSource{
		Driver:   viper.GetString("databaseSource.driver"),
		Username: viper.GetString("databaseSource.username"),
		Password: viper.GetString("databaseSource.password"),
		IPAddr:   viper.GetString("databaseSource.ip"),
		Port:     viper.GetInt("databaseSource.port"),
		DBName:   viper.GetString("databaseSource.database"),
		Charset:  viper.GetString("databaseSource.charset"),
		Loc:      url.QueryEscape(viper.GetString("databaseSource.loc")),
	}
	// "root:UVlY88m9suHLsthK@tcp(10.151.3.79:6446)/demo2?charset=utf8mb4&parseTime=True&loc=Local"
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		dbSource.Username,
		dbSource.Password,
		dbSource.IPAddr,
		dbSource.Port,
		dbSource.DBName,
		dbSource.Charset,
		dbSource.Loc)
	db, err := gorm.Open(dbSource.Driver, args)
	if err != nil {
		return err
	}
	db.SingularTable(true) // 禁止表名复数
	DB = db
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func InitDB() {
	var userInfo = &UserInfo{}
	var articleCategoryInfo = &ArticleCategoryInfo{}
	var post = &Post{}
	DB.AutoMigrate(userInfo) // create userInfo table
	DB.AutoMigrate(articleCategoryInfo) // create articleCategoryInfo table
	DB.AutoMigrate(post) // create post table
}

func Close() {
	if err := DB.Close(); err != nil { // close DB
		glog.Fatal("DB close failed , err = ", err)
	}
}
