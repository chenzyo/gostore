package model

// Commodity 商品模型
type Commodity struct {
	Model
	Name               string  `gorm:"size:64;not null;default:'';comment:'商品名称'"`
	Price              float64 `gorm:"type:decimal(10,2) unsigned;size:255;not null;default:'0';comment:'销售价'"`
	MarketPrice        float64 `gorm:"type:decimal(10,2) unsigned;size:255;not null;default:'0';comment:'市场价'"`
	CostPrice          float64 `gorm:"type:decimal(10,2) unsigned;size:255;not null;default:'0';comment:'成本价'"`
	Stock              uint    `gorm:"not null;default:'0';comment:'库存'"`
	ItemNo             string  `gorm:"size:64;not null;default:'';comment:'商品编号'"`
	DefaultImageId     uint    `gorm:"not null;default:'0';comment:'默认商品图id'"`
	AlbumImageIds      string  `gorm:"size:64;not null;default:'[]';comment:'商品相册'"`
	SellingStatus      int8    `gorm:"not null;default:'1';comment:'销售状态【1:上架 2:下架】'"`
	CategoryId         uint    `gorm:"not null;default:'0';comment:'分类id'"`
	BrandId            uint    `gorm:"not null;default:'0';comment:'品牌id'"`
	WarehouseId        uint    `gorm:"not null;default:'0';comment:'仓库id'"`
	SupplierId         uint    `gorm:"not null;default:'0';comment:'分类id'"`
	Detail             string  `gorm:"size:1024;not null;default:'';comment:'商品详情'"`
	CommodityType      int8    `gorm:"not null;default:'1';comment:'商品类型【1:实物 2:虚拟商品 3:卡券 4:服务类商品】'"`
	SellingType        int8    `gorm:"not null;default:'1';comment:'售卖类型【1:普通类型 2:组合商品 3:试用商品】'"`
	Order              uint    `gorm:"not null;default:'0';comment:'排序'"`
	SellingPoint       string  `gorm:"size:64;not null;default:'';comment:'商品卖点'"`
	SharingDescription string  `gorm:"size:256;not null;default:'';comment:'商品分享描述'"`
}

func (c *Commodity) TableName() string {
	return "commodities"
}
