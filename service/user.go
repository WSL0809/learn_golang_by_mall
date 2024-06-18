package service

import (
	"context"
	"init_golang/dao"
	"init_golang/model"
	"init_golang/pkg/e"
	"init_golang/pkg/e/util"
	"init_golang/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"`
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.SUCCESS
	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "key长度必须是16位",
		}
	}
	util.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, _ :=  userDao.ExistOrNotByUserName(service.UserName)
	// if err != nil {
	// 	code = e.ERROR
	// 	return serializer.Response{
	// 		Status: code,
	// 		Msg:    e.GetMsg(code),
	// 		Error:  err.Error(),
	// 	}
	// }
	if exist {
		code = e.ERROR_EXIST_USER
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Status:   model.Active,
		Avatar:   "avatar.jpg",
		Money:    util.Encrypt.AesEncoding("10000"),	
	}

	if err := user.SetPassword(service.Password); err != nil {
		code = e.ERROR_FAIL_ENCRYPT
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err := userDao.CreateUser(&user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}