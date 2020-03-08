package commodity_category

import (
	"fmt"
	"singo/model"
	"singo/serializer"
	"singo/serializer/site"

	"github.com/gin-gonic/gin"
)

type CreateCommodityCategoryService struct {
	Name     string `form:"name" json:"name" binding:"required,min=1,max=64"`
	ParentId uint   `form:"parent_id" json:"parent_id" binding:"numeric"`
}

// valid 验证表单
func (service *CreateCommodityCategoryService) valid() *serializer.Response {
	commodityCategory := model.CommodityCategory{}
	fmt.Println(service.ParentId)
	if service.ParentId != 0 {
		if err := model.DB.Where("id = ?", service.ParentId).First(&commodityCategory).Error; err != nil {
			return &serializer.Response{
				Code: 40001,
				Msg:  "父级id分类不存在",
			}
		}
	}
	return nil
}

// Login 用户登录函数
func (service *CreateCommodityCategoryService) Create(c *gin.Context) serializer.Response {

	commodityCategory := model.CommodityCategory{
		Name:     service.Name,
		ParentId: service.ParentId,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	//开启事务
	tx := model.DB.Begin()
	// 创建用户
	if err := tx.Create(&commodityCategory).Error; err != nil {
		tx.Rollback()
		return serializer.DBErr("创建失败", err)
	}

	tx.Commit()
	return site.BuildCommodityCategoryResponse(commodityCategory)
}
