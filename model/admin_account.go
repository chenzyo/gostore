package model

import (
	"golang.org/x/crypto/bcrypt"
)

// AdminAccount 管理员账户模型
type AdminAccount struct {
	Model
	UserName       string `gorm:"size:64;not null;default:'';comment:'用户名账号'"`
	PasswordDigest string `gorm:"size:256;not null;default:'';comment:'密码'"`
	RealName       string `gorm:"size:64;not null;default:'';comment:'真实名称'"`
	Gender         int8   `gorm:"not null;default:'1';comment:'性别【1:女 2:男 3:未知】'"`
	RoleId         uint   `gorm:"not null;default:'0';comment:'权限id'"`
	Enabled        int8   `gorm:"not null;default:'1';comment:'是否启用【1:启用 2:禁用】'"`
	IsSuper        int8   `gorm:"not null;default:'2';comment:'是否管理员【1:是 2:否】'"`
}

func (adminAccount *AdminAccount) TableName() string {
	return "admin_users"
}

// GetAdminUser 用ID获取用户
func GetAdminUser(ID interface{}) (AdminAccount, error) {
	var adminUser AdminAccount
	result := DB.Select("id,user_name,role_id,enabled,is_super").First(&adminUser, ID)
	return adminUser, result.Error
}

// CheckPassword 校验密码
func (adminAccount *AdminAccount) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(adminAccount.PasswordDigest), []byte(password))
	return err == nil
}
