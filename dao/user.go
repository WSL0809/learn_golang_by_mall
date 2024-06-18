package dao

import (
	"context"
	"init_golang/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{DB: newDBClient(ctx),}
}
func NewUserDaoByDB(db *gorm.DB) *UserDao {
    return &UserDao{DB: db}
}
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(model.User{}).Where("user_name = ?", userName).First(&user).Error
	
	if user == nil || err == gorm.ErrRecordNotFound {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(model.User{}).Create(user).Error
	return
}