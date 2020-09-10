package model

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"net/url"
)

// define DB for GetDB() func return
var DBClient *gorm.DB

// Init database
func InitDB() error {
	// connect database
	// db, err := gorm.Open("mysql", "root:UVlY88m9suHLsthK@tcp(10.151.3.79:6446)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	// 这里可以做成配置文件来读取配置
	dbInfo := DBInfo{
		username: viper.GetString("databaseSource.username"),
		password: viper.GetString("databaseSource.password"),
		ip:       viper.GetString("databaseSource.ip"),
		port:     viper.GetInt("databaseSource.port"),
		database: viper.GetString("databaseSource.database"),
		driver:   viper.GetString("databaseSource.driver"),
		charset:  viper.GetString("databaseSource.charset"),
		loc:      url.QueryEscape(viper.GetString("databaseSource.loc")),  // url.QueryEscape将/转义为%2f
		//loc:      "Asia%2fShanghai",
	}
	dbInfoStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s", dbInfo.username, dbInfo.password, dbInfo.ip, dbInfo.port, dbInfo.database, dbInfo.charset, dbInfo.loc)
	//dbInfoStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbInfo.user, dbInfo.password, dbInfo.ip, dbInfo.port, dbInfo.database)
	DB, err := gorm.Open(dbInfo.driver, dbInfoStr)
	if err != nil {
		return err
	}

	// enable singular table name
	DB.SingularTable(true)

	// create table
	DB.AutoMigrate(&User{}) // 此部分最好写到对应的controller中最佳
	DB.AutoMigrate(&Category{}) // 创建表
	DB.AutoMigrate(&Post{})

	DBClient = DB
	return nil
}

// return DB
func GetDB() *gorm.DB {
	return DBClient
}

// close db connection
func DBClose() {
	if err := DBClient.Close(); err != nil {
		glog.Fatal("db close failed: ", err)
	}
}
