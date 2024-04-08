package user

import (
	loginserviceV1 "github.com/keweiLv/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var LoginServiceClient loginserviceV1.LoginServiceClient

func InitRpcUserClient() {
	conn, err := grpc.Dial("127.0.0.1:8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	LoginServiceClient = loginserviceV1.NewLoginServiceClient(conn)
}
