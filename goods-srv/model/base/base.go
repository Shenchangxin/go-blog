package base

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现sql.Scanner接口，Scan将value扫描至Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	ID          int32 `gorm:"primarykey"`
	CreatedBy   string
	CreatedTime time.Time
	updatedBy   string
	updatedTime time.Time
}
