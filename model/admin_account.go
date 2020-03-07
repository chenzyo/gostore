package model

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

func (c *AdminAccount) TableName() string {
	return "admin_users"
}

// GetAdminUser 用ID获取用户
func GetAdminUser(ID interface{}) (AdminAccount, error) {
	var adminUser AdminAccount
	result := DB.First(&adminUser, ID)
	return adminUser, result.Error
}