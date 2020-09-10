package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rockwang.com/rock-second-demo/demo/commonTools"
	"rockwang.com/rock-second-demo/demo/repository"
	"rockwang.com/rock-second-demo/demo/response"
	"rockwang.com/rock-second-demo/demo/vo"
)

// 多态用法
type IArticleCategory interface {
	RestController
	AllInfo(ctx *gin.Context)
}

// 主结构体
type ArticleCategory struct {
	//DB *gorm.DB
	Repository repository.ICategoryResp
}

// 构建函数
func NewArticleCategory() IArticleCategory {
	//ac := &ArticleCategory{DB: model.GetDB()} // 老方法
	ac := &ArticleCategory{Repository: repository.NewCategoryRepository()} // 使用repository的多态中数据库操作方法

	return ac
}

// 新增一个文章分类
func (ac *ArticleCategory) Create(ctx *gin.Context) {
	var crv vo.CategoryRequestVo // name字段必须存在
	if err := ctx.ShouldBind(&crv); err != nil {
		errStr := fmt.Sprint("ctx should bind failed ,err = ", err.Error())
		response.Response(ctx, 422, 422006, errStr, nil)
		return
	}

	if len(crv.Name) == 0 {
		response.Response(ctx, 422, 422009, "name不可为空", nil)
		return
	}

	// 新增分类(name)到数据库 - 老方法
	//var info model.ArticleCategoryInfo
	//info.Name = crv.Name
	//ac.DB.Create(&info)

	// 新增分类(name)到数据库 - 使用repository的多态中数据库操作方法
	info := ac.Repository.CreateByName(crv.Name)

	response.Response(ctx, 200, 200001, "", gin.H{"Create_info": info})
}

// 更新一个文章分类
func (ac *ArticleCategory) Update(ctx *gin.Context) {
	var crv vo.CategoryRequestVo // name字段必须存在
	if err := ctx.ShouldBind(&crv); err != nil {
		errStr := fmt.Sprint("ctx should bind failed ,err = ", err.Error())
		response.Response(ctx, 422, 422006, errStr, nil)
		return
	}

	idStr := ctx.Params.ByName("id")

	id, err := commonTools.StrTransferInt(idStr)
	if err != nil {
		response.Response(ctx, 422, 422007, "id必须为数字", nil)
		return
	}

	// 先检查id是否存在,再更新 - 老方法
	//var info model.ArticleCategoryInfo
	//ac.DB.Where("id = ?", id).First(&info).
	//if info.ID == 0 {
	//	//errStr := fmt.Sprint("DB query id failed, err = ", err)
	//	response.Response(ctx, 422, 422008, "此id不存在", nil)
	//	return
	//}
	//
	//if err = ac.DB.Model(&info).Update("name", crv.Name).Error; err != nil {
	//	errStr := fmt.Sprint("DB update failed, err = ", err)
	//	response.InternalErrResp(ctx, 500002, errStr, nil)
	//	return
	//}

	// 先检查id是否存在,再更新 - 使用repository的多态中数据库操作方法
	err, httpCode, returnCode, info := ac.Repository.UpdateByID(id, crv.Name)
	if err != nil {
		response.Response(ctx, httpCode, returnCode, err.Error(), nil)
		return
	}

	response.Response(ctx, 200, 200, "", gin.H{"Update_info": info})
}

// 展示一个文章分类
func (ac *ArticleCategory) Show(ctx *gin.Context) {
	idStr := ctx.Params.ByName("id")

	id, err := commonTools.StrTransferInt(idStr)
	if err != nil {
		response.Response(ctx, 422, 422007, "id必须为数字", nil)
		return
	}

	// 检查id是否存在 - 老方法
	//var info model.ArticleCategoryInfo
	////if ac.DB.First(&info, id).RecordNotFound() {
	//ac.DB.First(&info, id)
	//if info.ID == 0 {
	//	response.Response(ctx, 422, 422008, "此id不存在", nil)
	//	return
	//}

	// 检查id是否存在 - 使用repository的多态中数据库操作方法
	err, httpCode, returnCode, info := ac.Repository.ShowByID(id)
	if err != nil {
		response.Response(ctx, httpCode, returnCode, err.Error(), nil)
		return
	}

	response.Response(ctx, 200, 200001, "", gin.H{"Show_info": info})

}

// 删除一个文章分类
func (ac *ArticleCategory) Delete(ctx *gin.Context) {
	idStr := ctx.Params.ByName("id")
	id, err := commonTools.StrTransferInt(idStr)
	if err != nil {
		response.Response(ctx, 422, 422007, "id必须为数字", nil)
		return
	}

	// 先查询,再删除 - 老方法
	//var aci model.ArticleCategoryInfo
	//ac.DB.Where("id = ?", id).First(&aci)
	//if aci.ID == 0 {
	//	response.Response(ctx, 422, 422008, "此id不存在", nil)
	//	return
	//}
	//
	//if err = ac.DB.Delete(model.ArticleCategoryInfo{}, id).Error; err != nil {
	//	response.InternalErrResp(ctx, 400001, "删除失败,请重试", nil)
	//	glog.Error("delete id failed , err = ", err)
	//	return
	//}

	// 先查询,再删除 - 使用repository的多态中数据库操作方法
	err, httpCode, returnCode := ac.Repository.DeleteByID(id)
	if err != nil {
		response.Response(ctx, httpCode, returnCode, err.Error(), nil)
		return
	}

	response.Response(ctx, 200, 200001, "删除成功", nil)
}

// 展示所有文章分类
func (ac *ArticleCategory) AllInfo(ctx *gin.Context) {
	// 查询所有信息 - 老方法
	//var allInfo []model.ArticleCategoryInfo
	//if err := ac.DB.Find(&allInfo).Error; err != nil {
	//	response.InternalErrResp(ctx, 400002, "查询失败,请重试", nil)
	//	return
	//}

	// 查询所有信息 - 使用repository的多态中数据库操作方法
	err, httpCode, returnCode, allInfo := ac.Repository.ShowAllInfo()
	if err != nil {
		response.Response(ctx, httpCode, returnCode, err.Error(), nil)
		return
	}

	response.Response(ctx, 200, 200001, "", gin.H{"All_info": allInfo})
}
