package models

import "time"

// BaseModel 基础模型
type BaseModel struct {
	ID uint64 `gorm:"column:id;primary_key;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:created_at;index;" json:"updated_at,omitempty"`
}