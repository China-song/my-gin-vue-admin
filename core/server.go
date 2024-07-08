package core

import (
	"fmt"
	"github.com/China-song/my-gin-vue-admin/global"
	"github.com/China-song/my-gin-vue-admin/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	s.ListenAndServe()
}
