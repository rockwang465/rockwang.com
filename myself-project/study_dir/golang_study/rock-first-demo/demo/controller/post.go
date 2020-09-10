package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
	"rockwang.com/rock-first-demo/demo/model"
	"rockwang.com/rock-first-demo/demo/response"
	"rockwang.com/rock-first-demo/demo/vo"
	"strconv"
)

/*
单篇文章增删改查操作
postman post请求:
url: 127.0.0.1:8088/post
{
	"category_id": 1,
	"title": "开发语言专题1-Golang",
	"head_image": "http://www.baidu.com/image/1.jpg",
	"content":"Go（又称 Golang）是 Google 的 Robert Griesemer，Rob Pike 及 Ken Thompson 开发的一种静态强类型、编译型语言。Go 语言语法与 C 相近，但功能上有：内存安全，GC（垃圾回收），结构形态及 CSP-style 并发计算。Go的语法接近C语言，但对于变量的声明有所不同。"
}

postman put请求:
根据UUID来修改内容
url: 127.0.0.1:8088/post/ee491ff3-4209-45b9-91b8-029839b9add7
{
	"category_id":1,
	"title": "开发语言专题1-java1",
	"head_image": "http://www.baidu.com/image/java1.jpg",
	"content":"ava 教程 Java 是由Sun Microsystems公司于1995年5月推出的高级程序设计语言。 Java可运行于多个平台,如Windows, Mac OS,及其他多种UNIX版本的系统。"
}
*/

type IPostController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
	PageList(ctx *gin.Context)
}

type postController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	DB := model.GetDB()
	return &postController{DB: DB}
}

func (pc *postController) Create(ctx *gin.Context) {
	// get user id from gin.Context
	user, exists := ctx.Get("user")
	if ! exists {
		response.Response(ctx, http.StatusNonAuthoritativeInfo, 203001, "权限不足", nil)
		return
	}
	userId := user.(model.User).ID // 强制转为User类型

	var pv vo.PostVo
	// 获取postman数据
	if err := ctx.ShouldBind(&pv); err != nil {
		response.Response(ctx, 422, 422012, "传入数据无效", nil)
		logrus.Info("传入数据异常, err = ", err)
		return
	}

	// 填入获取的数据到struct中，用于存储到数据库
	postData := model.Post{
		UserID:     userId,
		CategoryID: pv.CategoryID,
		Title:      pv.Title,
		HeadImage:  pv.HeadImage,
		Content:    pv.Content,
	}
	if err := pc.DB.Create(&postData).Error; err != nil {
		response.FailResp(ctx, "创建失败", nil)
		logrus.Warning("post data create failed , err = ", err)
		return
	}

	response.Response(ctx, 200, 200001, "创建成功", gin.H{"post_info": postData})
}

func (pc *postController) Update(ctx *gin.Context) {
	uuid := ctx.Params.ByName("uuid")
	if len(uuid) != 36 {
		response.Response(ctx, 422, 422013, "uuid长度必须为36", nil)
		return
	}

	var postData model.Post
	if pc.DB.Where("id = ?", uuid).First(&postData).RecordNotFound() {
		response.Response(ctx, 422, 422014, "此uuid不存在", nil)
		return
	}

	// 判断当前的user id和查询的postData中的user_id是否相同。如果不同，说明当前用户在操作其他用户的文章，是不允许的。
	user, exists := ctx.Get("user")
	if !exists {
		response.Response(ctx, 203, 203001, "权限不足", nil)
		return
	}
	userId := user.(model.User).ID // 强行转user类型
	logrus.Infof("context user id [%v], post user id [%v]", userId, postData.UserID)
	if userId != postData.UserID {
		response.Response(ctx, 422, 422013, "禁止操作其他用户文章", nil)
		return
	}

	var pv vo.PostVo
	if err := ctx.ShouldBind(&pv); err != nil {
		response.Response(ctx, 422, 422015, "post数据不完整", nil)
		return
	}

	//if err := pc.DB.Model(&postData).Update(map[string]interface{}{"category_id": pv.CategoryID, "title": pv.Title, "head_image": pv.HeadImage, "content": pv.Content}).Error; err != nil {
	if err := pc.DB.Model(&postData).Update(pv).Error; err != nil { // 这里直接使用struct比较方便，但gorm不推荐用struct，因为部分为0时是不操作的。
		response.Response(ctx, 500, 500006, "更新失败，请重试", nil)
		logrus.Errorf("update post failed, err = ", err)
		return
	}

	response.Response(ctx, 200, 200001, "更新成功", gin.H{"post_info": postData})
}

func (pc *postController) Show(ctx *gin.Context) {
	uuid := ctx.Params.ByName("uuid")
	if len(uuid) != 36 {
		response.Response(ctx, 422, 422013, "uuid长度必须为36", nil)
		return
	}

	var postData model.Post
	//if pc.DB.Where("id = ?", uuid).First(&postData).RecordNotFound() {

	// 将model.Category信息也查出来，放到 model.Post.Category中
	// 所以这里就需要关联model.Category.id
	// 如果model.Category的ID字段不叫ID，而是叫其他的，如categoryid,则需要model.Post的Category中gorm声明部分声明Category的ID(外键)为: categoryid
	// 示例: type Post struct {
	//	ID uuid.UUID `json:"id" gorm:"type:char(36); primary_key"`
	//	UserID     uint `json:"user_id" gorm:"not null"`
	//	CategoryID uint `json:"category_id" gorm:"not null"`
	//	Category   *Category `gorm:"foreignkey:categoryid"`  // Rock理解: 这里要声明清楚才行
	//  ... ...
	//}
	if pc.DB.Preload("Category").Where("id = ?", uuid).First(&postData).RecordNotFound() {
		response.Response(ctx, 422, 422014, "此uuid不存在", nil)
		return
	}
	// 查询效果如下
	// {
	//    "code": 200001,
	//    "data": {
	//        "post_info": {
	//            "id": "e9e36ad2-00d2-46cb-aa13-664d3d1be54d",
	//            "user_id": 2,
	//            "category_id": 3,
	//            "Category": { // 此处为关联category表查询到的信息
	//                "id": 3,
	//                "name": "c++",
	//                "create_at": "2020-08-04 19:41:42",
	//                "update_at": "2020-08-04 19:41:42"
	//            },
	//            "title": "开发语言专题1-C++2",
	//            "head_image": "http://www.baidu.com/image/c++1.jpg",
	//            "content": " C++是C语言的继承,它既可以进行C语言的过程化程序设计,又可以进行以抽象数据类型为特点的基于对象的程序设计,还可以进行以继承和多态为特点的面向对象...",
	//            "created_at": "2020-09-06 14:19:57",
	//            "updated_at": "2020-09-06 14:19:57"
	//        }
	//    },
	//    "msg": "查询成功"
	//}

	response.Response(ctx, 200, 200001, "查询成功", gin.H{"post_info": postData})
}

func (pc *postController) Delete(ctx *gin.Context) {
	uuid := ctx.Params.ByName("uuid")
	if len(uuid) != 36 {
		response.Response(ctx, 422, 422013, "uuid长度必须为36", nil)
		return
	}

	var postData model.Post
	if pc.DB.Where("id = ?", uuid).First(&postData).RecordNotFound() {
		response.Response(ctx, 422, 422014, "此uuid不存在", nil)
		return
	}

	// 判断当前的user id和查询的postData中的user_id是否相同。如果不同，说明当前用户在操作其他用户的文章，是不允许的。
	user, exists := ctx.Get("user")
	if !exists {
		response.Response(ctx, 203, 203001, "权限不足", nil)
		return
	}
	userId := user.(model.User).ID // 强行转user类型
	logrus.Infof("context user id [%v], post user id [%v]", userId, postData.UserID)
	if userId != postData.UserID {
		response.Response(ctx, 422, 422013, "禁止操作其他用户文章", nil)
		return
	}

	if err := pc.DB.Where("id = ?", uuid).Delete(&model.Post{}).Error; err != nil {
		response.Response(ctx, 500, 500007, "删除失败，请重试", nil)
		logrus.Error("delete uuid failed, err = ", err)
		return
	}

	response.Response(ctx, 200, 200001, "删除成功", nil)
}

func (pc *postController) PageList(ctx *gin.Context) {
	// 获取分页参数
	// 注意: ctx.DefaultQuery的page_num page_size要在postman的Params中输入才行
	pageNum, err := strconv.Atoi(ctx.DefaultQuery("page_num", "1")) // postman传入page_num值，没有则用默认值
	if err != nil {
		response.Response(ctx, 422, 422015, "page_num的值要为数字", nil)
		logrus.Info("strconv transfer to int failed, err = ", err)
		return
	}
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "20")) // postman传入page_size值，没有则用默认值
	if err != nil {
		response.Response(ctx, 422, 422015, "page_size的值要为数字", nil)
		logrus.Info("strconv transfer to int failed, err = ", err)
		return
	}

	// 分页
	var postsData []model.Post
	// 获取第pageNum页数据,当页限定展示pageSize条数据(如:获取第2页，当页限定展示4条数据)
	pc.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&postsData)
	// Order根据创建时间排序，最近创建的放到最上面
	// offset 设置偏移量，即指定在开始返回记录之前要跳过的记录数量
	// 所以要先跳过: (pageNum - 1) * pageSize 条数据

	// 前端渲染分页需要知道获取数据的总条数
	var total int
	//pc.DB.Model(&postsData).Count(&total)  // 我开始理解错了，以为只是当前需求的数据的总量
	pc.DB.Model(&model.Post{}).Count(&total) // 老师这里求的是所有数据的总量，这是对的。因为当前获取的数据只是所有表数据的一部分而已

	response.Response(ctx, 200, 200001, "成功获取分页信息",
		gin.H{
			"posts_info": postsData,
			"total":      total,
			"page_num":   pageNum,
			"page_size":  pageSize})
}






