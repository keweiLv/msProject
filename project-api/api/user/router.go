package user

import (
	"github.com/gin-gonic/gin"
	"github.com/keweiLv/project-api/router"
	"log"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Route(r *gin.Engine) {
	// 初始化grpc客户端连接

	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)

}
