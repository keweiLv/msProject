package user

import "github.com/gin-gonic/gin"

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandleUser{}
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
