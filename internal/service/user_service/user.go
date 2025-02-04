package user_service

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/ink-yht/CloudNest/internal/domain/user_domain"
	"github.com/ink-yht/CloudNest/internal/repository/user_repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService 定义了用户服务的接口
type UserService interface {
	Signup(ctx context.Context, req user_domain.EmailRegisterRequest) error
}

// UserServiceImpl 实现了 UserService 接口
type UserServiceImpl struct {
	repo user_repository.UserRepository
}

func NewUserService(repo user_repository.UserRepository) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (svc *UserServiceImpl) Signup(ctx context.Context, req user_domain.EmailRegisterRequest) error {
	// 校验请求
	if err := req.Validate(); err != nil {
		return err
	}

	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		// 密码加密失败
		return err
	}

	// 生成唯一ID
	userId, err := svc.GenerateUniqueID()
	if err != nil {
		// 密码加密失败
		return err
	}

	// 初始化用户
	user := user_domain.DefaultUser(req.Email, string(hash), req.Nickname, userId)

	// 插入用户数据
	err = svc.repo.Create(ctx, user_domain.User{
		Email:    user.Email,
		UserID:   user.UserID,
		Password: user.Password,
		UserInfo: user_domain.UserInfo{
			Nickname:          user.UserInfo.Nickname,
			UserID:            user.UserInfo.UserID,
			Signature:         user.UserInfo.Signature,
			Avatar:            user.UserInfo.Avatar,
			Birthday:          user.UserInfo.Birthday,
			Sex:               user.UserInfo.Sex,
			PointsRemaining:   user.UserInfo.PointsRemaining,
			NumberOfPosts:     user.UserInfo.NumberOfPosts,
			NumberOfResources: user.UserInfo.NumberOfResources,
		},
		UserRole: user_domain.UserRole{
			UserID:     user.UserRole.UserID,
			IsAdmin:    user.UserRole.IsAdmin,
			IsOrder:    user.UserRole.IsOrder,
			IsDisabled: user.UserRole.IsDisabled,
		},
	})
	if err != nil {
		return err
	}
	return nil

}

// GenerateUniqueID 用于每次调用时生成一个新的唯一ID。
func (svc *UserServiceImpl) GenerateUniqueID() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println("Error creating snowflake node:", err)
		return 0, err
	}

	userID := node.Generate()
	fmt.Println("Generated Snowflake ID:", userID)
	return int64(userID), nil
}
