package initialize

import (
	"github.com/China-song/my-gin-vue-admin/global"
	"gorm.io/gorm"
)

// Gorm
// 依据配置文件指定的数据库类型(如mysql), 连接到相应数据库, 并返回 gorm.DB 对象
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	//db := global.GVA_DB
	//err := db.AutoMigrate()
	//if err != nil {
	//
	//}
}
