package user_domain

import (
	"errors"
	"github.com/dlclark/regexp2"
	"time"
)

var (
	emailRegex                          = regexp2.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, regexp2.None)
	passwordRegex                       = regexp2.MustCompile(`^(?=.*[a-zA-Z])(?=.*[0-9])(?=.*[._~!@#$^&*])[A-Za-z0-9._~!@#$^&*]{8,20}$`, regexp2.None)
	ErrTheMailboxIsNotInTheRightFormat  = errors.New("电子邮件格式无效")
	ErrThePasswordIsNotInTheRightFormat = errors.New("密码长度必须为 8-20 个字符，并包含字母、数字和特殊字符")
	ErrThePasswordIsInconsistentTwice   = errors.New("两次密码不一致")
)

// EmailRegisterRequest 用户邮箱注册请求体
type EmailRegisterRequest struct {
	Email           string `json:"email"`
	Code            string `json:"code"`
	Nickname        string `json:"nickname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

// Validate 校验请求参数
func (req *EmailRegisterRequest) Validate() error {
	// 校验邮箱格式
	if match, _ := emailRegex.MatchString(req.Email); !match {
		return ErrTheMailboxIsNotInTheRightFormat
	}

	// 校验密码格式
	if match, _ := passwordRegex.MatchString(req.Password); !match {
		return ErrThePasswordIsNotInTheRightFormat
	}

	// 确认密码是否一致
	if req.Password != req.ConfirmPassword {
		return ErrThePasswordIsInconsistentTwice
	}

	return nil
}

// DefaultUser 初始化默认用户
func DefaultUser(email, password, nickname string, userId int64) *User {
	return &User{
		Email:    email,
		Password: password,
		UserID:   userId,
		UserInfo: UserInfo{
			Nickname:          nickname,
			UserID:            userId,
			Signature:         "即便总在期待后失望，即便总在孤独中疗伤。依然努力绽放、绝不投降",
			Avatar:            "/uploads/avatar/logo.png",
			Birthday:          time.Date(2025, time.January, 26, 0, 0, 0, 0, time.UTC),
			Sex:               3,
			PointsRemaining:   0,
			NumberOfResources: 0,
			NumberOfPosts:     0,
		},
		UserRole: UserRole{
			UserID:     userId,
			IsAdmin:    2,
			IsOrder:    2,
			IsDisabled: 2,
		},
	}
}
