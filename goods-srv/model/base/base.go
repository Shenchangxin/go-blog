package base

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID          int32 `gorm:"primarykey"`
	CreatedBy   string
	CreatedTime time.Time
	updatedBy   string
	updatedTime time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool
}
