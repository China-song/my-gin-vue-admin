package router

import "github.com/China-song/my-gin-vue-admin/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
