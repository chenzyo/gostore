package model

// CommodityPropertyRelation 商品属性关联模型
type CommodityPropertyRelation struct {
	Model
	CommodityId     uint `gorm:"not null;default:'0';comment:'商品id'"`
	ProductId       uint `gorm:"not null;default:'0';comment:'货品id'"`
	PropertyId      uint `gorm:"not null;default:'0';comment:'商品属性id'"`
	PropertyValueId uint `gorm:"not null;default:'0';comment:'商品属性值id'"`
}

func (c *CommodityPropertyRelation) TableName() string {
	return "commodity_property_relations"
}
