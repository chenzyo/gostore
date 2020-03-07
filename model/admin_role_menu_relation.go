package model

// AdminMenu 菜单模型
type AdminRoleMenuRelation struct {
	Model
	RoleId uint `gorm:"not null;default:'0';comment:'权限id'"`
	MenuId uint `gorm:"not null;default:'0';comment:'菜单id'"`
}

func (c *AdminRoleMenuRelation) TableName() string {
	return "admin_role_menu_relations"
}
