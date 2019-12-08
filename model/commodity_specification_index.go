package model

// CommoditySpecificationIndex 商品规格属性关联模型
type CommoditySpecificationIndex struct {
	Model
	CommodityId     uint `gorm:"not null;default:'0';comment:'商品id'"`
	PropertyId      uint `gorm:"not null;default:'0';comment:'商品属性id'"`
	PropertyValueId uint `gorm:"not null;default:'0';comment:'商品属性值id'"`
}

func (c *CommoditySpecificationIndex) TableName() string {
	return "commodity_specification_indexes"
}
