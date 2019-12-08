package model

import "time"

// UserInfo 用户明细模型
type UserInfo struct {
	Model
	User     User
	UserId   uint      `gorm:"not null;default:0;comment:'会员id'"`
	Mobile   string    `gorm:"size:32;not null;default:'';comment:'手机号'"`
	Email    string    `gorm:"size:128;not null;default:'';comment:'邮箱'"`
	Birthday time.Time `gorm:"not null;default:'0000-00-00 00:00:00';comment:'生日年月'"`
	RealName string    `gorm:"size:64;not null;default:'';comment:'真实姓名'"`
	NickName string    `gorm:"size:64;not null;default:'';comment:'昵称'"`
	Gender   int8      `gorm:"not null;default:'1';comment:'是否启用【1:女 2:男 3:未知】'"`
}
