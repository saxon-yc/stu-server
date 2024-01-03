package dbsvc

import (
	"student-server/internal/model"
	"time"
)

func (b *basicService) FindUser(user model.DupuserRequest) (result model.UserDb, err error) {
	err = b.gdb.Where("username = ?", user.UserName).First(&result).Error
	return
}

func (b *basicService) CreateUser(user model.LoginRequest) (err error) {
	err = b.gdb.Create(&model.UserDb{
		UserName: user.UserName, Password: user.Password, Menus: []string{"class", "student", "tag", "notice"}, CreateTime: time.Now(), UpdateTime: time.Now()}).Error

	return
}

func (b *basicService) UpdatePassword(user model.LoginRequest, password string) (result model.UserDb, err error) {
	err = b.gdb.Model(result).Where("username = ?", user.UserName).Updates(&model.UserDb{Password: password, UpdateTime: time.Now()}).Error
	return
}

func (b *basicService) RefreshToken(user model.LoginRequest, token string) (result model.UserDb, err error) {
	err = b.gdb.Model(result).Where("username = ?", user.UserName).Update("token", token).Error
	return
}
