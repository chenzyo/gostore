package model

// CommoditySpecificationRelation 商品规格属性关联模型
type CommoditySpecificationRelation struct {
	Model
	CommodityId     uint `gorm:"not null;default:'0';comment:'商品id'"`
	PropertyId      uint `gorm:"not null;default:'0';comment:'商品属性id'"`
	PropertyValueId uint `gorm:"not null;default:'0';comment:'商品属性值id'"`
}

func (c *CommoditySpecificationRelation) TableName() string {
	return "commodity_specification_relation"
}
