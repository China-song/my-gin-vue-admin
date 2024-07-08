package main

import (
	"github.com/China-song/my-gin-vue-admin/core"
	"github.com/China-song/my-gin-vue-admin/global"
	"github.com/China-song/my-gin-vue-admin/initialize"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	// TODO: 检验config 2024/07/05 18:05
	//initialize.OtherInit()
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		//
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
