package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"rockwang.com/rock-second-demo/demo/model"
	"rockwang.com/rock-second-demo/demo/response"
	"rockwang.com/rock-second-demo/demo/vo"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	DB := model.GetDB()
	//DB.AutoMigrate(&model.Post{})
	return &PostController{DB: DB}
}

func (pc *PostController) Create(ctx *gin.Context) {
	// 获取用户传入数据
	var pv vo.PostVo
	if err := ctx.ShouldBind(&pv); err != nil {
		response.Response(ctx, 422, 422016, "数据不完整", nil)
		logrus.Infof("context should bind failed, err = ", err)
		return
	}

	// 获取上下文中保存的user
	user, exists := ctx.Get("user")
	if ! exists {
		response.Response(ctx, 403, 403001, "权限不足", nil)
		logrus.Infof("context get user key is not exists")
		return
	}

	userId := user.(model.UserInfo).ID
	if userId == 0 {
		response.Response(ctx, 403, 403001, "权限不足", nil)
		logrus.Infof("context user id equal 0")
		return
	}

	// 生成数据
	var postData = model.Post{
		UserId:     userId,
		CategoryId: pv.CategoryID,
		Title:      pv.Title,
		HeadImg:    pv.HeadImg,
		Content:    pv.Content,
	}

	// 创建信息
	if err := pc.DB.Create(&postData).Error; err != nil {
		response.Response(ctx, 400, 400003, "创建失败,请重试", nil)
		logrus.Errorf("db create value failed, err = ", err)
		return
	}

	response.Response(ctx, 200, 200001, "创建成功", gin.H{"post_info": postData})
}

func (pc *PostController) Update(ctx *gin.Context) {
	// 获取用户传入数据
	var pv vo.PostVo
	if err := ctx.ShouldBind(&pv); err != nil {
		response.Response(ctx, 422, 422010, "数据不完整", nil)
		return
	}

	// 获取uuid
	uuid := ctx.Params.ByName("uuid")
	if len(uuid) != 36 {
		response.Response(ctx, 422, 422011, "uuid必须为36位", nil)
		return
	}

	// 获取上下文中保存的user
	user, exists := ctx.Get("user")
	userId := user.(model.UserInfo).ID
	if !exists || userId == 0 {
		response.Response(ctx, 403, 403001, "权限不足", nil)
		return
	}

	// 查询uuid是否存在
	var postData model.Post
	if pc.DB.Where("id = ?", uuid).First(&postData).RecordNotFound() {
		response.Response(ctx, 422, 422012, "该uuid不存在", nil)
		return
	}

	// 确认当前上下文中的user id和数据库的user id相同，不同则操作的用户非自己的文章(禁止操作其他作者的文章)
	if userId != postData.UserId {
		response.Response(ctx, 422, 422013, "禁止操作其他作者的文章", nil)
		return
	}

	// 更新数据
	if err := pc.DB.Model(&postData).Update(pv).Error; err != nil {
		response.Response(ctx, 400, 400004, "更新失败,请重试", nil)
		logrus.Warnf("db update failed, err = ", err)
		return
	}
	response.Response(ctx, 200, 200001, "更新成功", gin.H{"post_info": postData})
}

func (pc *PostController) Show(ctx *gin.Context) {
	// 获取uuid
	uuid := ctx.Params.ByName("uuid")
	if len(uuid) != 36 {
		response.Response(ctx, 422, 422011, "uuid必须为36位", nil)
		return
	}

	// 查询uuid是否存在
	var postData model.Post
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
	// 另外: Preload("Category")中必须是struct字段，所以Category是大写C开头
	if pc.DB.Preload("Category").Where("id = ?", uuid).First(&postData).RecordNotFound() {
		response.Response(ctx, 422, 422012, "该uuid不存在", nil)
		return
	}

	response.Response(ctx, 200, 200001, "查询成功", gin.H{"post_info": postData})
}

func (pc *PostController) Delete(ctx *gin.Context) {
	// 获取uuid
	uuid := ctx.Params.ByName("uuid")
	if len(uuid) != 36 {
		response.Response(ctx, 422, 422011, "uuid必须为36位", nil)
		return
	}

	// 获取上下文中保存的user
	user, exists := ctx.Get("user")
	userId := user.(model.UserInfo).ID
	if !exists || userId == 0 {
		response.Response(ctx, 403, 403001, "权限不足", nil)
		return
	}

	// 查询uuid是否存在
	var postData model.Post
	if pc.DB.Where("id = ?", uuid).First(&postData).RecordNotFound() {
		response.Response(ctx, 422, 422012, "该uuid不存在", nil)
		return
	}

	// 确认当前上下文中的user id和数据库的user id相同，不同则操作的用户非自己的文章(禁止操作其他作者的文章)
	if userId != postData.UserId {
		response.Response(ctx, 422, 422013, "禁止操作其他作者的文章", nil)
		return
	}

	// 删除数据
	if err := pc.DB.Where("id = ?", uuid).Delete(&model.Post{}).Error; err != nil {
		response.Response(ctx, 400, 400001, "删除失败,请重试", nil)
		logrus.Warnf("db delete failed, err = ", err)
		return
	}

	response.Response(ctx, 200, 200001, "删除成功", nil)
}

func (pc *PostController) PageList(ctx *gin.Context) {
	pageNum, err1 := strconv.Atoi(ctx.DefaultQuery("page_num", "1"))
	pageSize, err2 := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	if err1 != nil || err2 != nil {
		response.Response(ctx, 422, 422014, "page_num page_size必须为数字", nil)
		return
	}

	// 获取第pageNum页数据,当页限定展示pageSize条数据(如:获取第2页，当页限定展示4条数据)
	// Order根据创建时间排序，最近创建的放到最上面
	// offset 设置偏移量，即指定在开始返回记录之前要跳过的记录数量 ,所以要先跳过: (pageNum - 1) * pageSize 条数据
	var postsData []model.Post
	//if err := pc.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&postsData).Error; err != nil {
	// Rock: 多加载了Preload 联表Category的信息
	if err := pc.DB.Preload("Category").Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&postsData).Error; err != nil {
		logrus.Warnf("db order offset limit find failed, err = ", err)
		response.Response(ctx, 500, 500008, "查询失败，请重试", nil)
		return
	}

	// 获取总条数 ,前端渲染分页需要知道获取数据的总条数
	var total int
	if err := pc.DB.Model(&model.Post{}).Count(&total).Error; err != nil{
		logrus.Warnf("db count failed, err = ", err)
		response.Response(ctx, 500, 500008, "查询失败，请重试", nil)
		return
	}

	response.Response(ctx, 200, 200001, "查询成功", gin.H{
		"posts_info": postsData,
		"total":      total,
		"page_num":   pageNum,
		"page_size":  pageSize,
	})
}
