package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User3 struct {
	gorm.Model
	Name string
	Age  int
	Active bool
}

func main() {
	// 连接Mysql数据库
	db, err := gorm.Open("mysql", "root:UVlY88m9suHLsthK@tcp(10.151.3.79:6446)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SingularTable(true)
	db.AutoMigrate(&User3{})

	var user User3
	//user = User3{Age: 19, Name: "nova"}
	//db.Debug().Create(&user)
	//db.Debug().First(&user)
	db.Debug().First(&user) // 1.先找到第一条数据

	// 更新
	user.Name = "Mark"
	user.Age = 20

	db.Debug().Save(&user)  // 2.修改第一条数据 (默认修改所有字段)
	//  UPDATE `user3` SET `created_at` = '2020-07-15 11:41:59', `updated_at` = '2020-07-15 17:22:27', `deleted_at` = NULL, `name` = 'Mark', `age` = 20, `active` = false  WHERE `user3`.`deleted_at` IS NULL AND `user3`.`id` = 1

	db.Debug().Model(&user).Update("name", "小王子") // 3.使用Update修改name字段
	//  UPDATE `user3` SET `name` = '小王子', `updated_at` = '2020-07-15 17:22:27'  WHERE `user3`.`deleted_at` IS NULL AND `user3`.`id` = 1

	m1 := map[string]interface{}{
		"name":   "liwenzhou",
		"age":    28,
		"active": true,
	}

	db.Debug().Model(&user).Update(m1)  // 4.m1列出来的字段都会更新
	// UPDATE `user3` SET `active` = true, `age` = 28, `name` = 'liwenzhou', `updated_at` = '2020-07-15 17:23:58'  WHERE `user3`.`deleted_at` IS NULL AND `user3`.`id` = 1

	db.Debug().Model(&user).Select("age").Updates(m1)  // 5.只会更新m1的age字段，因为这里select固定了字段名，其他字段不会更新
	// UPDATE `user3` SET `age` = 28, `updated_at` = '2020-07-15 17:24:59'  WHERE `user3`.`deleted_at` IS NULL AND `user3`.`id` = 1

	db.Debug().Model(&user).Omit("active").Updates(m1)  // 6.除了active字段外都会更新
	//  UPDATE `user3` SET `age` = 28, `name` = 'liwenzhou', `updated_at` = '2020-07-15 17:26:44'  WHERE `user3`.`deleted_at` IS NULL AND `user3`.`id` = 1
}
