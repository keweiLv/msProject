package user

import (
	"context"
	"github.com/gin-gonic/gin"
	common "github.com/keweiLv/project-common"
	loginservicev1 "github.com/keweiLv/project-user/pkg/service/login.service.v1"
	"net/http"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := LoginServiceClient.GetCaptcha(c, &loginservicev1.CaptchaRequest{Mobile: mobile})
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(2001, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(resp.Code))
}
