package main

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User4 struct {
	gorm.Model
	Name  string
	Age   int
	Hobby string
	//Active bool
}

func main() {
	// 连接Mysql数据库
	db, err := gorm.Open("mysql", "root:UVlY88m9suHLsthK@tcp(10.151.3.79:6446)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		glog.Fatal(err)
	}
	// 禁用表名复数
	db.SingularTable(true)

	// 把结构体和数据表进行对应
	db.AutoMigrate(&User4{}) // 只传入结构体的定义，用于表结构定义
	defer db.Close()

	// 创建几条数据
	//u1 := User4{Name: "rock", Age: 28, Hobby: "吃饭"}
	//u2 := User4{Name: "lss", Age: 26, Hobby: "睡觉"}
	//u3 := User4{Name: "nova", Age: 24, Hobby: "打豆豆"}
	//u4 := User4{Name: "mark", Age: 29, Hobby: "打豆豆"}
	//db.Debug().Create(&u1)  //Rock: 这里总是忘记 & 指针传入
	//db.Debug().Create(&u2)
	//db.Debug().Create(&u3)
	//db.Debug().Create(&u4)

	// 1.删除
	//var u User4
	//u.ID = 2  // 删除id为2的数据，这里id为主键，这是对的
	//db.Debug().Delete(&u)  // 软删除，只是把deleted_at状态修改掉
	//UPDATE `user4` SET `deleted_at`='2020-07-15 19:10:22'  WHERE `user4`.`deleted_at` IS NULL AND `user4`.`id` = 2

	// 2.如果不用主键删除，会导致所有数据全被删除
	//var u2 User4
	//u2.Name = "rock"
	//db.Debug().Delete(&u2)
	// UPDATE `user4` SET `deleted_at`='2020-07-15 19:34:14'  WHERE `user4`.`deleted_at` IS NULL
	// 发现where后面没有接 name = 4的条件，所以所有的数据全被删掉了。

	// 3.使用where进行批量删除，就不需要接主键了
	db.Debug().Where("name = ?", "rock").Delete(User4{})
	// UPDATE `user4` SET `deleted_at`='2020-07-15 20:05:47'  WHERE `user4`.`deleted_at` IS NULL AND ((name = 'rock'))
}


