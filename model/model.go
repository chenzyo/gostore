package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt *time.Time `gorm:"comment:'创建时间'"`
	UpdatedAt *time.Time `gorm:"comment:'更新时间'"`
}
