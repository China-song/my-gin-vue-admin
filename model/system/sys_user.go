package system

import (
	"github.com/China-song/my-gin-vue-admin/global"
	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                // 用户UUID
	Username string    `json:"userName" gorm:"index;comment:用户登录名"`             // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`                        // 用户登录密码
	NickName string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`       // 用户昵称
	Enable   int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}
