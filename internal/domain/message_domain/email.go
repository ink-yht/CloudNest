package message_domain

import "time"

// VerificationCode 验证码结构体
type VerificationCode struct {
	Email      string    `json:"email"`       // 接收邮箱
	Code       string    `json:"code"`        // 验证码
	ExpiryTime time.Time `json:"expiry_time"` // 过期时间
}
