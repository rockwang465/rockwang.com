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
	"rockwang.com/rock-second-demo/demo/model"
	"rockwang.com/rock-second-demo/demo/router"
	"time"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:     "demo",
	Short:   "This is a demo website.",
	Example: "demo server",
	Version: "v1.0.0",
}

var startServerCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start demo server",
	Example: "demo server",
	Version: "v1.0.0",
	Run:     startDemoServer,
}

func startDemoServer(cmd *cobra.Command, args []string) {
	// 这里的cmd 就是 startServerCmd
	// load config file
	//cmd.Flags().StringVar(&configFile, "application", "/etc/demo/application.yaml", "demo server config file")
	cmd.Flags().StringVar(&configFile, "application", "E:\\mygopath\\src\\rockwang.com\\rock-second-demo\\demo\\cmd\\demo\\application.yml", "demo server config file")
	viper.SetConfigType("yml")
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		glog.Error("viper read in config failed , err = ", err)
	}

	fmt.Println("Run demo server ...")
	err := model.ConnectDB() // 连接数据库，生成DB
	if err != nil {
		glog.Fatal("gorm open failed , err = ", err)
	}
	defer model.Close()
	model.InitDB() // create DB

	routers := router.NewRouter()
	routers.RegisterAndLogin()
	routers.UserInfo()
	routers.ArticleCategory()
	routers.Post()

	serverPort := viper.GetString("server.port")
	if serverPort == "" {
		serverPort = "8080"
	}
	svrPort := ":" + serverPort
	fmt.Printf("Demo server : http://127.0.0.1%v\n", svrPort)

	// listen server
	server := http.Server{
		Addr:           svrPort,
		Handler:        routers.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			glog.Fatal("Demo serve listen failed , err = ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	glog.Info("Shutdown Server ...")

	// 等待关闭信号
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		glog.Fatal("Server shutdown : ", err)
	}
	glog.Info("Server exiting")
}

// 下面初始化config功能写到startDemoServer函数里, 所以这里注释了
//func InitConfig() {
//	configFilePath := path.Join("../../", "config")
//	viper.SetConfigType("yml")
//	viper.SetConfigName("application")
//	viper.AddConfigPath(configFilePath)
//	err := viper.ReadInConfig()
//	if err != nil {
//		glog.Fatal("viper read in config failed ,err = ", err)
//	}
//}

func init() { // init()函数自动执行
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startServerCmd)
}

func main() {
	//InitConfig()  // load config file-之前老师用的,这里不用.因为要仿照sophon,将初始化config做到server启动部分中去.
	err := rootCmd.Execute()
	if err != nil {
		glog.Fatal("Execute root cmd error = ", err)
	}
	glog.Flush()
}

// demo2:
// 视频进程顺序
// 1.本地go环境配置
// 2.添加用户注册registry路由及函数
// 3.重构项目代码-涉及model、router、controller、util
// 4.用户登录login路由及函数
// 5.jwt中间件, 敏感信息处理dto（处理userInfo路由）
// 6.封装统一的请求返回格式 response
// 7.viper从文件中读取配置（config）,如果server port没有则使用默认端口启动。
// 12.跨域问题:中间件CORSMiddleware 浏览器同源http://developer.mozilla.org/zh-CN/docs/Web/Security/Same-origin_policy
//    registry和login两个controller模块中,解决获取前端传参接收不到问题,用ctx.Bind(&requestUser) ，将数据绑定到requestUser中。 requestUser是model.User类型
//    这样解决当前端没有传参就可以报错.
//    vue可以自己做到登录成功后，保存token，跳转到主页。
// 17.新增文章表Category,访问/category/的增删改查路由，进行文章表的增删改查操作。
//    另外:gorm时间格式化、多态形式增删改查功能。
// 18.A.使用validator，定义结构体替换原有的model.Category，解决name字段不为空问题。
//      解决方法: 结构体中tag为 `json: "name" binding:"required"`，让name字段不能为空，用binding required解决）；
//      使用ctx.ShouldBind代替ctx.Bind
//    B.文章分类（将数据库操作部分封装到repository中）
//      将controller/articleCategory.go文章分类的数据库操作部分放到repository/categoryRepository.go中,并做成多态方式
//      将category重复操作数据库的地方，写到repository中(create update show delete)--这里没有用，因为前面没有做成多态
//      将panic的err部分，用自带的recovery功能，做成中间件函数，return给用户。需要将中间件放到路由中Use。(方便debug)
// 19.A.对文章分类之下的单篇文章进行增删改查操作。
//      增加是通过提前生成uuid作为该篇文章的id，并从ctx.Get("user")中获取用户id作为作者的id，防止他人修改作者的id。
//      删除是通过url中传入uuid进行删除（注意增加非当前user_id禁止操作）。
//      修改是根据uuid进行修改（注意增加非当前user_id禁止操作）。
//      查看是根据uuid进行查看，另外增加外键连接到category表，将category的分类信息也在总信息中展示出来。
//    B.增加分页功能
//      当前显示第N页，每页显示M条数据，计算数据的总条数。

// 需要优化部分:
// model.User表创建部分从dbOperator中移到controller的代码中，方便对应管理
// 时间格式要优化为 2020-08-02 16:22:22 - ok
// articleCategory改为多态方式

// demo2需要优化新增功能:(基于demo1额外自己加的)
// 1.支持修改用户名、密码、手机号 -- ok
// 2.将config文件存放指定位置读取,linux上的需求(例如:/etc/demo/config.yaml) -- ok
// 3.将demo version 功能移到demo/version.go中. -- ok
// 4.glog弃用，转用logrus
// 5.recovery.go中要优化各种panic返回值给出来，顺义的优化代码: infra-console-service中的error_handler.go (/gitlab.sz.sensetime.com/galaxias/infra-console-service/console/server/middleware)

// 未测试通过的问题
// token expire问题

// 老师代码:
// https://github.com/haydenzhourepo/gin-vue-gin-essential
// https://github.com/haydenzhourepo/gin-vue-gin-essential-vue
