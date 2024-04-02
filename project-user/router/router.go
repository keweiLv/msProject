package router

import (
	"github.com/gin-gonic/gin"
	"github.com/keweiLv/project-user/api/user"
	"github.com/keweiLv/project-user/config"
	loginserviceV1 "github.com/keweiLv/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"log"
	"net"
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

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			loginserviceV1.RegisterLoginServiceServer(g, loginserviceV1.New())
		},
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}
