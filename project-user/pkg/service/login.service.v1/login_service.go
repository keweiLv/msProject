package login_service_v1

import (
	"context"
	"fmt"
	common "github.com/keweiLv/project-common"
	"github.com/keweiLv/project-common/errs"
	"github.com/keweiLv/project-user/pkg/dao"
	"github.com/keweiLv/project-user/pkg/model"
	"github.com/keweiLv/project-user/pkg/repo"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, req *CaptchaRequest) (*CaptchaResponse, error) {
	mobile := req.Mobile
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	// 生成验证码
	code := generateVerificationCode()
	go func() {
		zap.L().Info("短信平台调用成功,发送短信")
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := ls.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			zap.L().Error("验证码存入redis出错:REGISTER_%s:%s", zap.String("mobile", mobile), zap.String("code", code))
		}
		zap.L().Debug("将手机号和验证码存入redis成功:REGISTER_%s:%s", zap.String("mobile", mobile), zap.String("code", code))

	}()
	return &CaptchaResponse{Code: code}, nil
}

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间的纳秒部分作为随机种子

	// 生成6位随机数字字符串
	code := ""
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10)) // 生成0-9的随机数字
	}
	return code
}
