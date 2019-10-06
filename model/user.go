package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	Model
	UserName       string    `gorm:"size:64;not null;default:'';comment:'用户名账号'"`
	PasswordDigest string    `gorm:"size:256;not null;default:'';comment:'密码'"`
	Enabled        int8      `gorm:"not null;default:'1';comment:'是否启用【1:启用 2:禁用】'"`
	Locked         int8      `gorm:"not null;default:'2';comment:'是否启用【1:锁定 2:非锁定】'"`
	VerifiedAt     time.Time `gorm:"not null;default:'0000-00-00 00:00:00';comment:'验证时间'"`
}

//
type LoginReq struct {
	Phone string `json:"mobile"`
	Pwd   string `json:"pwd"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"

	Sign string = "tokenstring"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
