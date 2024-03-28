package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/keweiLv/project-common"
	"github.com/keweiLv/project-common/logs"
	"github.com/keweiLv/project-user/config"
	"github.com/keweiLv/project-user/router"
	"log"
)

func main() {
	r := gin.Default()

	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "/Users/kezi/GolandProjects/msProject/logs/debug/project-debug.log",
		InfoFileName:  "/Users/kezi/GolandProjects/msProject/logs/info/project-info.log",
		WarnFileName:  "/Users/kezi/GolandProjects/msProject/logs/project-error.log",
		MaxSize:       500,
		MaxAge:        3,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}

	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr)
}
