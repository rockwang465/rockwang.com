package vo

// validator数据验证,所有要配合binding required参数
type CategoryRequestVo struct {
	Name string `json:"name" binding:"required"`
}