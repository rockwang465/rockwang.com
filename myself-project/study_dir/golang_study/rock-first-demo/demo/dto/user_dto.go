package dto

import (
	"rockwang.com/rock-first-demo/demo/model"
)

type UserResDto struct {
	ID        uint
	Name      string
	Telephone string
	//CreateAt  time.Time
}

// 处理用户的表数据中的敏感信息，返回不敏感数据给用户
func UserInfoResponse(value model.User) *UserResDto {
	newModel := &UserResDto{
		ID:        value.ID,
		Name:      value.Name,
		Telephone: value.Telephone,
		//CreateAt:  value.CreatedAt,
	}
	return newModel
}
