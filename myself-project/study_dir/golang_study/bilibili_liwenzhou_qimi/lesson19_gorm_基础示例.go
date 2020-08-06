package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 3.表结构定义
type User struct {
	gorm.Model // gorm.Model包含id CreateAt UpdateAt DeleteAt 这4个字段
	Name       string
	Age        int64
	Hobby      string
}

func main() {
	// 1.连接数据库
	// utf8mb4 是支持表情的
	//db, err := gorm.Open("mysql", "root:UVlY88m9suHLsthK@tcp(10.151.3.79:6446)/userinfo?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:UVlY88m9suHLsthK@tcp(10.151.3.79:6446)/userinfo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		glog.Fatal(err)
	}
	// 2.禁用表名复数
	db.SingularTable(true)

	// 4.创建表(自动迁移，把结构体和数据表进行对应)
	db.AutoMigrate(&User{}) // 只传入结构体的定义，用于表结构定义
	defer db.Close()

	// 5.创建数据行
	//u1 := User{Name: "rock", Age: 28, Hobby: "吃饭"}
	//u2 := User{Name: "lss", Age: 27, Hobby: "睡觉"}
	//u3 := User{Name: "hh", Age: 29, Hobby: "打豆豆"}
	//db.Debug().NewRecord(u1)
	//db.Debug().Create(&u1)  // 正常只需要create，不需要NewRecord
	//db.Debug().Create(&u2)
	//db.Debug().Create(&u3)
	//db.Debug().NewRecord(u1)

	// 6.查询
	var u User
	//  SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` ASC LIMIT 1
	//db.Debug().First(&u)  // 查第一条

	//   执行db.Debug().First(&u)，则为: SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND `user`.`id` = 1
	// 不执行db.Debug().First(&u)，则为: SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL
	db.Debug().Find(&u)  // 不执行db.Debug().First(&u)，貌似不会影响第一条数据，不知道为啥。(测试时只有两条数据)
	fmt.Printf("u: %v\n",u)
	fmt.Printf("u: %#v\n",u)

	// 7.更新
	//  UPDATE `user` SET `hobby` = '羽毛球', `updated_at` = '2020-07-14 16:13:35'  WHERE `user`.`deleted_at` IS NULL
	//db.Debug().Model(&u).Update("hobby", "羽毛球")

	// 8.删除
	//db.Debug().Delete(&u)

	// Rock总结:
	// 第一个: 不管是增删改查，都会受到前面的操作影响
	//         例如前面做过 db.First(&u)，那么之后查、删、更新都只会对第一条数据进行操作。
	//         原因是: gorm对通一个变量 &u操作，将就会受影响。
	// 第二个: 删除数据不是真删除，而是在deleted_at字段添加了当次删除的操作时间。
	//        以后查询是不会查到的，因为查询是要求 deleted_at == NULL 才可以的。
}
