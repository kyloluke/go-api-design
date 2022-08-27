package models

import (
	"github.com/spf13/cast"
	"time"
)

type BaseModel struct {
	// tag 之间需要输入一个空格
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

func (a *BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
