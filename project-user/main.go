package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/keweiLv/project-common"
	"github.com/keweiLv/project-user/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, "project-user", ":80")
}
