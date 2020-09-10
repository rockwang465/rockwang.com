package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"net/http"
	"rockwang.com/rock-first-demo/demo/model"
	"rockwang.com/rock-first-demo/demo/response"
	"strconv"
)

// 处理文章分类，进行文章的增删改查

// 需要转为多态功能

type Category struct {
	DB *gorm.DB
}

// 构造函数
func NewCategory() *Category {
	DB := model.GetDB() // 获取DB客户端
	return &Category{DB: DB}  // 这里是不使用repository中的构造函数
}

// 基于传入的id进行创建此id的文章分类
//func (c *Category)Create()gin.HandlerFunc{
//	return func(ctx *gin.Context){
//	}
//}

// 通过name进行数据创建
func (c *Category) Create(ctx *gin.Context) {
	var category = model.Category{}
	// 这里postman raw中传入json字符串，不可以有逗号 {"name":"rock123"}
	// 这里应该用ctx.ShouldBind才对,我写错了
	err := ctx.Bind(&category) // ctx.Bind()使用postman raw中传入json字符串进行绑定
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500004, "ctx bind err:"+err.Error(), nil)
		return
	}
	// 注意: model.Category中的id在grom中是自动自增的，所以不需要获取id。
	//name := ctx.PostForm("name")

	if category.Name == "" { // 判断是否为空
		response.Response(ctx, http.StatusUnprocessableEntity, 422006, "名称不能为空", nil)
		return
	}

	var cgRes model.Category
	c.DB.Where("name = ?", category.Name).First(&cgRes)
	if cgRes.Name != "" { // 非空表示数据库中已存此name
		response.Response(ctx, http.StatusUnprocessableEntity, 422007, "此名称已存在，请换个名字", nil)
		return
	}

	if err = c.DB.Create(&category).Error; err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500005, "创建文章分类失败", nil)
		glog.Error("db create category failed, err = ", err)
		return
	}
	response.Response(ctx, 200, 200011, "", gin.H{
		"msg":       "文章分类创建成功",
		"id":        category.ID,
		"name":      category.Name,
		"create_at": category.CreatedAt,
	})
}

// 查找id对应的值
func (c *Category) Show(ctx *gin.Context) {
	idStr := ctx.Param("id") // 获取url的参数
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422008, "id必须为数字", nil)
		return
	}
	var category = model.Category{ID: uint(id),}
	c.DB.First(&category)
	if category.Name == "" {
		response.Response(ctx, 400, 400002, "无此id的文章分类", nil)
		return
	}

	response.SuccessResp(ctx, "", gin.H{
		"id":        category.ID,
		"name":      category.Name,
		"create_at": category.CreatedAt,
	})
}

// 通过id查找数据，更新name对应的值
func (c *Category) Update(ctx *gin.Context) {
	// ByName返回键与给定名称匹配的第一个参数的值。
	// 如果没有找到匹配的参数，则返回空字符串。
	// Rock: Params.ByName 表示: 根据传入的参数名称来获取对应名称的值。如果没有找到，则返回空字符串。
	id, err := strconv.Atoi(ctx.Params.ByName("id")) //获取url的传参id
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422011, err.Error(), nil)
		return
	}

	var categoryArgs = model.Category{}
	err = ctx.Bind(&categoryArgs) //获取body中的传参(这里主要获取name值)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500004, "ctx bind err:"+err.Error(), nil)
		return
	}

	if categoryArgs.Name == "" {
		response.Response(ctx, http.StatusUnprocessableEntity, 422010, "name不能为空", nil)
		return
	}

	var updateCategory model.Category
	// 先根据id查询到第一条数据，并放到updateCategory struct中保存。
	// 如果查询不到数据，则return
	if c.DB.Where("id = ?", id).First(&updateCategory).RecordNotFound() {
		response.Response(ctx, http.StatusNotFound, 404001, "此分类id不存在", nil)
		return
	}

	// 修改 name=新的name
	c.DB.Model(&updateCategory).Update("name", categoryArgs.Name) // 修改完成
	response.SuccessResp(ctx, "分类name修改成功", gin.H{
		"id":        updateCategory.ID,
		"name":      updateCategory.Name,
		"create_at": updateCategory.CreatedAt,
	})
}

// 通过id进行删除数据
func (c *Category) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id")) // 获取要删除的id
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422011, err.Error(), nil)
		return
	}

	var category model.Category
	if c.DB.Where("id = ?", id).First(&category).RecordNotFound() { // 确认有此id的数据可供删除
		response.Response(ctx, http.StatusNotFound, 404001, "此id不存在", nil)
		return
	}

	//c.DB.Where("id = ?", id).Delete(&category) // 删除数据
	if err := c.DB.Delete(model.Category{}, id).Error; err != nil { // 老师的写法，加上error更安全
		response.FailResp(ctx, "删除失败", nil)
		return
	} // 删除数据
	response.SuccessResp(ctx, "分类id删除成功", nil)
}
