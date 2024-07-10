package main

import (
	"github.com/China-song/my-gin-vue-admin/core"
	"github.com/China-song/my-gin-vue-admin/global"
	"github.com/China-song/my-gin-vue-admin/initialize"
	"go.uber.org/zap"
)

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     1.0
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		//
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
