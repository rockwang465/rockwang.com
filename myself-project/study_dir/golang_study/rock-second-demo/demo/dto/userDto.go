package dto

import (
	"rockwang.com/rock-second-demo/demo/model"
)

type UserData struct {
	ID        uint
	Name      string
	Telephone string
}

// 过滤敏感数据
func UserDto(user *model.UserInfo) *UserData {
	var userData = &UserData{
		ID:        user.ID,
		Name:      user.Name,
		Telephone: user.Telephone,
	}
	return userData
}
