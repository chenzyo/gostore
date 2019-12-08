package model

// User 用户模型
type UserCart struct {
	Model
	CommodityId uint `gorm:"not null;default:'0';comment:'商品id'"`
	ProductId   uint `gorm:"not null;default:'0';comment:'货品id'"`
	Quantity    uint `gorm:"not null;default:'0';comment:'数量'"`
	UserId      uint `gorm:"not null;default:0;comment:'会员id'"`
	IsSelected  int8 `gorm:"not null;default:'1';comment:'是否选中【1:选中 2:非选中】'"`
}
