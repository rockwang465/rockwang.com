package repository

// 用于数据库的增删改查操作，给其他函数直接调用，相当于工具
// 此处工具是专门为articleCategory.go 这里提供的
// 另外必须articleCategory中有多态，否则功能不能放进去

//type CategoryRepository struct {
//	DB *gorm.DB
//}
//
//func NewCategoryRepository() *CategoryRepository {
//	return &CategoryRepository{DB: model.GetDB()}
//}
//
//func (c *CategoryRepository) Create(id int, name string) error {
//
//}
//
//func (c *CategoryRepository) SelectById(id int, name string) error {
//
//}
