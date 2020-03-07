package model

// AdminUser 角色模型
type AdminRole struct {
	Model
	RoleName    string `gorm:"size:64;not null;default:'';comment:'用户名账号'"`
	Description string `gorm:"size:256;not null;default:'';comment:'描述'"`
	Enabled     int8   `gorm:"not null;default:'1';comment:'是否启用【1:启用 2:禁用】'"`
}

func (c *AdminRole) TableName() string {
	return "admin_roles"
}
