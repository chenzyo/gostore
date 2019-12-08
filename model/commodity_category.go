package model

// CommodityCategory 商品分类模型
type CommodityCategory struct {
	Model
	Name     string `gorm:"size:64;not null;default:'';comment:'分类名称'"`
	ParentId uint   `gorm:"not null;default:'0';comment:'分类父级id'"`
}

func (c *CommodityCategory) TableName() string {
	return "commodity_categories"
}
