package site

import (
	"singo/model"
	"singo/serializer"
)

// User 用户序列化器
type CommodityCategory struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ParentId  uint   `json:"parent_id"`
	CreatedAt int64  `json:"created_at"`
}

// BuildCommodityCategory 商品分类序列化
func BuildCommodityCategory(commodityCategory model.CommodityCategory) CommodityCategory {
	return CommodityCategory{
		ID:        commodityCategory.ID,
		Name:      commodityCategory.Name,
		ParentId:  commodityCategory.ParentId,
		CreatedAt: commodityCategory.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildCommodityCategoryResponse(commodityCategory model.CommodityCategory) serializer.Response {
	return serializer.Response{
		Data: BuildCommodityCategory(commodityCategory),
	}
}
