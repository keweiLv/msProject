package router

import (
	"github.com/gin-gonic/gin"
	"github.com/keweiLv/project-user/api/user"
)

type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

func InitRouter(r *gin.Engine) {
	rg := New()
	rg.Route(&user.RouterUser{}, r)
}
