package model

import "time"

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null;default:'0000-00-00 00:00:00';comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"not null;default:'0000-00-00 00:00:00';comment:'更新时间'"`
}
