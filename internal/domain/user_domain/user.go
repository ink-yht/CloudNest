package user_domain

import "time"

// User 用户
type User struct {
	ID         int64     `json:"id"`          // ID
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	Email      string    `json:"email"`       // 邮箱，使用唯一键
	UserID     int64     `json:"user_id"`     // 用户ID 使用唯一键
	Password   string    `json:"password"`    // 加密密码
	UserInfo   UserInfo  `json:"user_info"`
	UserRole   UserRole  `json:"user_role"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID                int64     `json:"id"`                  // ID
	CreateTime        time.Time `json:"create_time"`         // 创建时间
	UpdateTime        time.Time `json:"update_time"`         // 更新时间
	UserID            int64     `json:"user_id"`             // 用户ID
	Nickname          string    `json:"nickname"`            // 昵称
	Signature         string    `json:"signature"`           // 个性签名
	Avatar            string    `json:"avatar"`              // 头像
	Birthday          time.Time `json:"birthday"`            //生日 默认2025-01-26
	Sex               int8      `json:"sex"`                 // 性别 1 男，2 女，3 未选择 默认未选择(3)
	PointsRemaining   int32     `json:"points_remaining"`    // 剩余积分
	NumberOfResources int16     `json:"number_of_resources"` // 资源数
	NumberOfPosts     int16     `json:"number_of_posts"`     // 帖子数
}

// UserRole 用户角色
type UserRole struct {
	ID         int64     `json:"id"`          // ID
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
	UserID     int64     `json:"user_id"`     // 用户ID
	IsAdmin    int8      `json:"is_admin"`    // 是否管理员 1 是管理员，2 不是管理员 默认不是(2)
	IsDisabled int8      `json:"is_disabled"` // 是否禁用 1 禁用，2 未禁用 默认未禁用(2)
	IsOrder    int8      `json:"is_order"`    // 是否有订单 1 有订单，2 无订单 默认无订单(2)
}
