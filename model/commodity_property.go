package model

// CommodityProperty 商品属性模型
type CommodityProperty struct {
	Model
	PropertyName string `gorm:"size:64;not null;default:'';comment:'属性名称'"`
	IsFilter     int8   `gorm:"not null;default:'1';comment:'是否筛选项【1:是 2:否】'"`
	InputType    int8   `gorm:"not null;default:'1';comment:'属性值类型【1:单选项 2:多选项】'"`
	Order        uint   `gorm:"not null;default:'0';comment:'排序'"`
}

func (c *CommodityProperty) TableName() string {
	return "commodity_properties"
}
