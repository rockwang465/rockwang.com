时间 | 日期
--- | --- 
开始时间 | 2020-08
结束时间 | 2020-09

# 前言
## 1.关于所有代码来源
+ 当前课程来自于: `https://www.bilibili.com/video/BV1CE411H7bQ`
+ 课程名称: `【评论送书】Go语言 Gin+Vue 前后端分离实战 - OceanLearn`
```
老师在github上的代码:
https://github.com/haydenzhourepo/gin-vue-gin-essential
https://github.com/haydenzhourepo/gin-vue-gin-essential-vue
```

## 2.老师的源码
+ 老师的源码已经下载下来了，项目名称为:`gin-vue-gin-essential`

## 3.关于自己写的代码部分
+ 首先我是看了老师的视频，然后`rock-first-demo`一边看，一边写。
+ 而`rock-second-demo`算是自己默写，且进行了很多改进优化。
+ 所以以后来复习代码，建议以看`rock-second-demo`为主。部分感觉不对劲的可以去看`rock-first-demo`。
+ 至于自己的改进和还没有做好的优化内容，都写在了`main.go`中哦。

# 目录介绍
## `cmd`
+ main函数入口目录

## `controller`
+ router路由中的func操作函数

## `routers`
+ 路由、路由组定义

## `middleWare`
+ router路由中需要的中间件函数

## `config`
+ 当前服务的配置文件存放位置
+ 另外:后面放到`cmd/demo/application.yml`这里了

## `repository`
+ 对数据库进行操作的部分封装到这里了。后面操作直接调用这里的函数即可。

## `util`
+ 工具

## `model`
+ 数据库模型
+ 数据库操作

## `common`
+ 通用函数

## `dto`
+ 对请求或者响应的数据进行敏感数据处理
+ 也就是说，本身响应给用户context数据里可能密码、创建时间等不该返回的，
+ 这时候就要做一层数据封装，只返回部分数据给用户即可。

## `response`
+ 响应数据封装
+ 将http状态码、code、msg等都放入封装函数，这样不用每次都取写
```
正常需要这样:
ctx.JSON(http.StatusOK, gin.H{
			"code": "200", "msg": "can not get user info",
		})

有了封装函数只需要这样调用了:
response.Response(ctx, 200, 200001, msg, data)
对应函数定义的类型是:
func Response(ctx *gin.Context, httpCode , code int, msg string, data gin.H){
}
```

### `vo`目录
+ 存放`validator`
+ `validator`是用于前端或者postman发送请求a时,需要含有validator定义的字段
+ 例如,
```
type定义
type category struct{
    Name string json: "name" binding:"required"
}
required表示必须含有name字段是必须要有的

那么后面代码部分就可以用ShouldBind进行判断了
if err := ctx.ShouldBind(&category); err != nil{
    fmt.pring("name字段没有值,数据验证失效")
}
```