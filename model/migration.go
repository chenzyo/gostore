package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserInfo{})
	DB.AutoMigrate(&Commodity{})
	DB.AutoMigrate(&CommodityProperty{})
	DB.AutoMigrate(&CommodityPropertyIndex{})
	DB.AutoMigrate(&CommodityPropertyValue{})
	DB.AutoMigrate(&CommoditySpecificationIndex{})
	DB.AutoMigrate(&Product{})
}
