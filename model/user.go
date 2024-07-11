package model

import "time"
import "gorm.io/gorm"

type BaseModel struct {
	ID          int32 `gorm:"primarykey"`
	CreatedBy   string
	CreatedTime time.Time
	updatedBy   string
	updatedTime time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool
}
type User struct {
	BaseModel
	Phone    string `gorm:"index:idx_phone;unique;type;varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	NickName string `gorm:"type:varchar(20)"`
	Sex      string `gorm:"column:sex;default:1;type:varchar(6) comment '0女1男'"`
}
