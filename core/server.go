package core

import (
	"fmt"
	"github.com/China-song/my-gin-vue-admin/global"
	"github.com/China-song/my-gin-vue-admin/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
