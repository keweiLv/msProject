package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/keweiLv/project-common"
)

func main() {
	r := gin.Default()
	srv.Run(r, "project-user", ":80")
}
