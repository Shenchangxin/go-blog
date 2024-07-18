package model

import "go-blog/goods-srv/model/base"

type User struct {
	base.BaseModel
	UserName string `gorm:"type:varchar(20)"`
	Phone    string `gorm:"index:idx_phone;unique;type;varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	NickName string `gorm:"type:varchar(20)"`
	Sex      string `gorm:"column:sex;default:1;type:varchar(6) comment '0女1男'"`
}
