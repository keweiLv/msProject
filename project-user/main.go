package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/keweiLv/project-common"
	"github.com/keweiLv/project-user/config"
	"github.com/keweiLv/project-user/router"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)
	// grpc注册
	gc := router.RegisterGrpc()
	stop := func() {
		gc.Stop()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
