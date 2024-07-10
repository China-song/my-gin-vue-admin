package system

import (
	"errors"
	"fmt"
	"github.com/China-song/my-gin-vue-admin/global"
	"github.com/China-song/my-gin-vue-admin/model/system"
	"github.com/China-song/my-gin-vue-admin/utils"
)

type UserService struct{}

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user system.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, nil
}
