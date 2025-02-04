package user_web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ink-yht/CloudNest/internal/domain/user_domain"
	"github.com/ink-yht/CloudNest/internal/service/code_service"
	"github.com/ink-yht/CloudNest/internal/service/user_service"
	"github.com/ink-yht/CloudNest/internal/web"
	"github.com/ink-yht/CloudNest/pkg/logger"
	"net/http"
)

const biz = "signup"

// 确保 UserHandler 上实现了 Handler 接口
var _ web.Handler = (*UserHandler)(nil)

type UserHandler struct {
	svc     user_service.UserService
	codeSvc code_service.CodeService
	l       logger.Logger
}

func NewUserHandler(svc user_service.UserService, codeSvc code_service.CodeService, l logger.Logger) *UserHandler {
	return &UserHandler{
		svc:     svc,
		codeSvc: codeSvc,
		l:       l,
	}
}

// RegisterRoutes 路由注册
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	router := server.Group("/api/user")
	//router.POST("/signup", u.SignUp)                        // 用户注册
	router.PUT("/signup/email/code", u.SendSignupEmailCode) // 发验证码
	router.POST("/signup/email/code", u.SignupEmail)        // 校验验证码
	//router.POST("/login", u.Login)    // 用户登录
	//router.PUT("/users/:id", u.Login) // 用户修改个人信息
	//router.GET("/users/:id", u.Login) // 用户信息获取
	//router.GET("/users/:id", u.Login) // 用户退出登录
}

func (u *UserHandler) SendSignupEmailCode(ctx *gin.Context) {
	type SignupEmailReq struct {
		Email string `json:"email"`
	}
	var req SignupEmailReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	err := u.codeSvc.Send(ctx, req.Email, biz)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 2,
			"msg":  "系统异常",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "发送成功",
	})
}

func (u *UserHandler) SignupEmail(ctx *gin.Context) {
	var req user_domain.EmailRegisterRequest
	err := ctx.Bind(&req)
	if err != nil {
		// 出错会返回 400 错误
		return
	}
	ok, err := u.codeSvc.Verify(ctx, biz, req.Email, req.Code)
	if err != nil {
		ctx.JSON(200, web.Result{
			Code: 2,
			Msg:  "系统错误",
		})
		return
	}
	if !ok {
		ctx.JSON(200, web.Result{
			Code: 1,
			Msg:  "验证码错误",
		})
		return
	}

	err = u.svc.Signup(ctx, req)
	if errors.Is(err, user_domain.ErrTheMailboxIsNotInTheRightFormat) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "电子邮件格式无效",
			Data: nil,
		})
		u.l.Warn("电子邮件格式无效", logger.String("email", req.Email))
		return
	}
	if errors.Is(err, user_domain.ErrThePasswordIsNotInTheRightFormat) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "密码长度必须为 8-20 个字符，并包含字母、数字和特殊字符",
			Data: nil,
		})
		u.l.Warn("密码格式不对", logger.String("email", req.Email))
		return
	}
	if errors.Is(err, user_domain.ErrThePasswordIsInconsistentTwice) {
		ctx.JSON(http.StatusOK, web.Result{
			Code: 1,
			Msg:  "两次密码不一致",
			Data: nil,
		})
		u.l.Warn("两次密码不一致", logger.String("email", req.Email))
		return
	}
}
