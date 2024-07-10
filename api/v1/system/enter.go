package system

import (
	"github.com/China-song/my-gin-vue-admin/service"
)

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
