package service

import (
	"k8s-server/config"

	"github.com/pkg/errors"

	"github.com/wonderivan/logger"
)

var Login login

type login struct{}

// 验证账号密码
func (l *login) Auth(username, password string) (err error) {
	if username == config.Config.GetString("User.adminuser") && password == config.Config.GetString("User.adminpwd") {
		return nil
	} else {
		logger.Error("登录失败, 用户名或密码错误")
		return errors.New("登录失败, 用户名或密码错误")
	}
	// return nil
}
