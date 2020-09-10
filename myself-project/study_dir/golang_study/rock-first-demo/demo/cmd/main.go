package main

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"rockwang.com/rock-first-demo/demo/model"
	"rockwang.com/rock-first-demo/demo/routers"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

/*
	1.具有登录注册功能的页面
	2.支持配置文件定义:服务端口、数据库基础信息、token过期时间
	3.支持登陆验证，token过期
	4.支持helm操作: 服务list、单实例部署、单实例删除、单实例信息查看
*/

var rootCmd = &cobra.Command{
	Use:        "demo",
	Short:      "This is a demo website",
	SuggestFor: []string{"demo"}, // 没理解啥用处
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Used to get the version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo version : v1.1.0")
	},
}

var newServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Used to start the demo server",
	Run:   startDemoServer, // start demo server
}

func startDemoServer(cmd *cobra.Command, args []string) {
	// connect database
	err := model.InitDB()
	if err != nil {
		glog.Fatal("Connect database failed, err = ", err)
		return
	}

	// close db connection
	defer model.DBClose()

	// router define
	router := routers.InitRouter()
	routers.RouterDefine(router)

	srvPort := viper.GetString("server.port")
	var srvPortAddr string
	if srvPort == "" {
		srvPortAddr = ":8080"
	} else {
		srvPortAddr = ":" + srvPort
	}

	// Run routers
	// https://github.com/skyhee/gin-doc-cn 介绍了http.ListenAndServer、router.Run()、signal.Notify、context.WithTimeout用法
	s := &http.Server{
		//Addr:           ":8088",
		Addr:           srvPortAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			glog.Error("http ListenAndServe failed :", err)
		}
	}()

	//if err = router.Run(":8088"); err != nil {  // 这种Run的方式太简单了，入门用可以，后期就不要用了。
	//	panic(err)
	//}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// 顺义 王鑫的用法
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	glog.Info("Shutdown Server ...")

	// 其他代码，等待关闭信号
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		glog.Fatal("Server Shutdown: ", err)
	}
	glog.Info("Server exiting")
}

// Add command to demo server
func init() { // init()函数自动执行
	rootCmd.AddCommand(versionCmd)   // get server version
	rootCmd.AddCommand(newServerCmd) // start demo server
}

// Use viper to read config
func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		glog.Fatal("os get path failed: ", err)
	}
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	pathDir, _ := filepath.Split(workDir) // 只有这一种取巧的方法，拿到上级目录
	//fmt.Printf("dir:[%v], file:[%v]\n", pathDir, pathFile) // dir:[E:\mygopath\src\rockwang.com\rock-first-demo\demo\], file:[cmd]
	configPath := path.Join(pathDir, "config")
	viper.AddConfigPath(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		glog.Fatal("viper read config failed: ", err)
	}
}

func InitLog(){
	logrus.SetLevel(logrus.InfoLevel) // 设置日志等级

	logrus.SetFormatter(&nested.Formatter{ // 格式化日志
		// HideKeys:        true,
		TimestampFormat: time.RFC3339,
		//FieldsOrder:     []string{"name", "age"},
	})
}

func main() {
	InitLog()
	InitConfig()
	fmt.Printf("server port: [%v]", viper.GetString("server.port"))
	err := rootCmd.Execute()
	if err != nil {
		glog.Fatal(err)
	}
	glog.Flush()
}

// 功能介绍及新功能增加:
// 1.login功能
// 2.优化gin.H{}返回，做成函数
// 3.jwt返回token
// 4.register和login获取内容，应该用ctx.bind方式放到 表结构体中才对
// 5.viper读取配置文件，定义server和数据库、token过期时间
// 6.router.run支持多端口配置文件传入，存在定义的端口，则用定义的，没有则用默认的。.
// 7.认证中间件，解决登录后返回给用户的内容
// 8.转换数据库create at等时间格式和时区问题
// 最后: 文章部分：文章数据库，增删改查功能，多态实现，数据库设计时间修改格式、增加时区到数据库模型

// 视频进程顺序
// 1.本地go环境配置
// 2.添加用户注册registry路由及函数
// 3.重构项目代码-涉及model、router、controller、util
// 4.用户登录login路由及函数
// 5.jwt中间件, 敏感信息处理dto（处理userInfo路由）
// 6.封装同意的请求返回格式 response
// 7.viper从文件中读取配置（config）,如果server port没有则使用默认端口启动。
// 12.跨域问题:中间件CORSMiddleware 浏览器同源http://developer.mozilla.org/zh-CN/docs/Web/Security/Same-origin_policy
//    解决获取前端传参接收不到问题,用ctx.Bind(&requestUser) ，将数据绑定到requestUser中。 requestUser是model.User类型
//    vue可以自己做到登录成功后，保存token，跳转到主页。
// 17.新增文章表Category,访问/category/的增删改查路由，进行文章表的增删改查操作。
//    另外:gorm时间格式化、多态形式增删改查功能。
// 18.A.使用validator，定义结构体替换原有的model.Category，解决name字段不为空问题。
//      解决方法: 结构体中tag为 `json: "name" binding:"required"`，让name字段不能为空，用binding required解决）；
//      使用ctx.ShouldBind代替ctx.Bind
//    B.文章分类（将数据库操作部分封装到repository中）
//      将category重复操作数据库的地方，写到repository中(create update show delete)--这里没有用，因为前面没有做成多态
//      将panic的err部分，用自带的recovery功能，做成中间件函数，return给用户。需要将中间件放到路由中Use。(方便debug)
// 19.A.对文章分类之下的单篇文章进行增删改查操作。
//      增加是通过提前生成uuid作为该篇文章的id，并从ctx.Get("user")中获取用户id作为作者的id，防止他人修改作者的id。
//      删除是通过url中传入uuid进行删除（注意增加非当前user_id禁止操作）。
//      修改是根据uuid进行修改（注意增加非当前user_id禁止操作）。
//      查看是根据uuid进行查看，另外增加外键连接到category表，将category的分类信息也在总信息中展示出来。
//    B.增加分页功能
//      当前显示第N页，每页显示M条数据。

// 需要优化部分:
// model.User表创建部分从dbOperator中移到controller的代码中，方便对应管理
// 时间格式要优化为 2020-08-02 16:22:22 - ok
// articleCategory改为多态方式

// 老师代码:
// https://github.com/haydenzhourepo/gin-vue-gin-essential
// https://github.com/haydenzhourepo/gin-vue-gin-essential-vue
