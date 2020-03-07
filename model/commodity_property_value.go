package model

// CommodityPropertyValue 商品属性值模型
type CommodityPropertyValue struct {
	Model
	CommodityId     uint `gorm:"not null;default:'0';comment:'商品id'"`
	ProductId       uint `gorm:"not null;default:'0';comment:'货品id'"`
	PropertyId      uint `gorm:"not null;default:'0';comment:'商品属性id'"`
	PropertyValueId uint `gorm:"not null;default:'0';comment:'商品属性值id'"`
}

func (c *CommodityPropertyValue) TableName() string {
	return "commodity_property_value"
}
