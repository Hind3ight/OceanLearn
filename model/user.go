package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20); not null"`
	Telephone string `gorm:"type:varchar(110); not null; unique"`
	Password  string `gorm:"size:255; not null"`
}

type UserReq struct {
}

type UserResp struct {
	Name      string
	Telephone string
}

func ToUserResp(user User) UserResp {
	return UserResp{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
