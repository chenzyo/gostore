package model

// Product 货品模型
type Product struct {
	Model
	CommodityId    uint    `gorm:"not null;default:'0';comment:'商品id'"`
	Name           string  `gorm:"size:64;not null;default:'';comment:'商品名称'"`
	SpecInfo       string  `gorm:"size:256;not null;default:'';comment:'货品规格信息'"`
	Price          float64 `gorm:"type:decimal(10,2) unsigned;size:255;not null;default:'0';comment:'销售价'"`
	MarketPrice    float64 `gorm:"type:decimal(10,2) unsigned;size:255;not null;default:'0';comment:'市场价'"`
	CostPrice      float64 `gorm:"type:decimal(10,2) unsigned;size:255;not null;default:'0';comment:'成本价'"`
	Stock          uint    `gorm:"not null;default:'0';comment:'库存'"`
	SkuNo          string  `gorm:"size:64;not null;default:'';comment:'货品编号'"`
	DefaultImageId uint    `gorm:"not null;default:'0';comment:'默认商品图id'"`
	SellingStatus  int8    `gorm:"not null;default:'1';comment:'销售状态【1:上架 2:下架】'"`
	CommodityType  int8    `gorm:"not null;default:'1';comment:'商品类型【1:实物 2:虚拟商品 3:卡券 4:服务类商品】'"`
	SellingType    int8    `gorm:"not null;default:'1';comment:'售卖类型【1:普通类型 2:组合商品 3:试用商品】'"`
	LimitQuantity  uint    `gorm:"not null;default:'0';comment:'限购量 (0表示不限)'"`
}
