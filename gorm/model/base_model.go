package model

import "time"

type BaseModel struct {
	ID        uint      `gorm:"column:id;primary_key"  json:"id"`
	CreatedAt time.Time `gorm:"column:created_at"      json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"      json:"updated_at"`
}
