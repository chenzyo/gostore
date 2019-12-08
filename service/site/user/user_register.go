package user

import (
	"singo/model"
	"singo/serializer"
	"time"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	UserName string `form:"userName" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	//PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		UserName:   service.UserName,
		Enabled:    1,
		Locked:     2,
		VerifiedAt: time.Now(),
		//Status:   model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	//开启事务
	tx := model.DB.Begin()
	// 创建用户
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return serializer.ParamErr("注册失败", err)
	}

	userInfo := model.UserInfo{
		UserId:   user.ID,
		Mobile:   service.UserName,
		NickName: service.UserName,
		Gender:   3,
		//Status:   model.Active,
	}

	if err := tx.Create(&userInfo).Error; err != nil {
		tx.Rollback()
		return serializer.ParamErr("注册失败", err)
	}

	tx.Commit()
	return serializer.BuildUserResponse(user)
}
