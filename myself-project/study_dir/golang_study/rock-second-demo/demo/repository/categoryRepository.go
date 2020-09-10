package repository

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"rockwang.com/rock-second-demo/demo/model"
)

/*
对文章分类 controller/articleCategory.go中数据库操作部分放到这里来
这里以多态方式展现
*/

type ICategoryResp interface {
	CreateByName(string) *model.ArticleCategoryInfo
	DeleteByID(int) (error, int, int)
	UpdateByID(int, string) (error, int, int, *model.ArticleCategoryInfo)
	ShowByID(int) (error, int, int, *model.ArticleCategoryInfo)
	ShowAllInfo() (error, int, int, []*model.ArticleCategoryInfo)
}

type CategoryResp struct {
	DB *gorm.DB
}

func NewCategoryRepository() ICategoryResp { // 这里interface类型不需要加*
	return &CategoryResp{DB: model.GetDB()}
}

func (cr *CategoryResp) CreateByName(name string) *model.ArticleCategoryInfo {
	var info model.ArticleCategoryInfo
	info.Name = name
	cr.DB.Create(&info)
	return &info
}

func (cr *CategoryResp) DeleteByID(id int) (error, int, int) {
	var aci model.ArticleCategoryInfo
	cr.DB.Where("id = ?", id).First(&aci)
	if aci.ID == 0 {
		//	response.Response(ctx, 422, 422008, "此id不存在", nil)
		return fmt.Errorf("%s", "此id不存在"), 422, 422008
	}

	if err := cr.DB.Delete(model.ArticleCategoryInfo{}, id).Error; err != nil {
		//	response.InternalErrResp(ctx, 400001, "删除失败,请重试", nil)
		glog.Error("delete id failed , err = ", err)
		return fmt.Errorf("%s", "删除失败,请重试"), 400, 400001
	}
	return nil, 0, 0
}

func (cr *CategoryResp) UpdateByID(id int, name string) (error, int, int, *model.ArticleCategoryInfo) {
	var info model.ArticleCategoryInfo
	cr.DB.Where("id = ?", id).First(&info)
	if info.ID == 0 {
		return fmt.Errorf("%s", "此id不存在"), 422, 422008, nil
	}

	if err := cr.DB.Model(&info).Update("name", name).Error; err != nil {
		errStr := fmt.Errorf("DB update failed, err = %s", err)
		return errStr, 400, 400001, nil
	}

	return nil, 0, 0, &info

}

func (cr *CategoryResp) ShowByID(id int) (error, int, int, *model.ArticleCategoryInfo) {
	var info model.ArticleCategoryInfo
	cr.DB.First(&info, id)
	if info.ID == 0 {
		return fmt.Errorf("%s", "此id不存在"), 422, 422008, nil
	}
	return nil, 0, 0, &info

}

func (cr *CategoryResp) ShowAllInfo() (error, int, int, []*model.ArticleCategoryInfo) {
	var allInfo []*model.ArticleCategoryInfo
	if err := cr.DB.Find(&allInfo).Error; err != nil {
		glog.Error("db exec find failed, err = ", err)
		return fmt.Errorf("%s", "查询失败,请重试"), 400, 400002, nil
	}
	return nil, 0, 0, allInfo
}
