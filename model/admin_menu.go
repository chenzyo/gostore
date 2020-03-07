package model

// AdminMenu 菜单模型
type AdminMenu struct {
	Model
	ParentId    uint   `gorm:"not null;default:'0';comment:'父级id'"`
	RoleId      uint   `gorm:"not null;default:'0';comment:'权限id'"`
	MenuName    string `gorm:"size:64;not null;default:'';comment:'菜单名称'"`
	Description string `gorm:"size:64;not null;default:'';comment:'描述'"`
	MenuPath    string `gorm:"size:256;not null;default:'';comment:'菜单路径'"`
	Enabled     int8   `gorm:"not null;default:'1';comment:'是否启用【1:启用 2:禁用】'"`
}

func (c *AdminMenu) TableName() string {
	return "admin_menus"
}
