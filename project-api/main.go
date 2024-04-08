package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/keweiLv/project-api/api"
	"github.com/keweiLv/project-api/config"
	"github.com/keweiLv/project-api/router"
	srv "github.com/keweiLv/project-common"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
