package model

//执行数据迁移

func migration() {
	//自动迁移模式
	//管理员相关
	DB.AutoMigrate(&AdminAccount{})
	DB.AutoMigrate(&AdminMenu{})
	DB.AutoMigrate(&AdminRole{})
	DB.AutoMigrate(&AdminRoleMenuRelation{})

	//会员相关
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserAttribute{})
	DB.AutoMigrate(&UserCart{})

	//商品相关
	DB.AutoMigrate(&Commodity{})
	DB.AutoMigrate(&CommodityCategory{})
	DB.AutoMigrate(&CommodityProperty{})
	DB.AutoMigrate(&CommodityPropertyRelation{})
	DB.AutoMigrate(&CommodityPropertyValue{})
	DB.AutoMigrate(&CommoditySpecificationRelation{})
	DB.AutoMigrate(&Product{})

	//订单相关

}
