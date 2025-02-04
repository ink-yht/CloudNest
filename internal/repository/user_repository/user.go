package user_repository

import (
	"context"
	"github.com/ink-yht/CloudNest/internal/domain/user_domain"
	"github.com/ink-yht/CloudNest/internal/repository/dao/user_dao"
)

type UserRepository interface {
	Create(ctx context.Context, user user_domain.User) error
}

type UserRepositoryImpl struct {
	dao user_dao.UserDao
}

func NewUserRepository(dao user_dao.UserDao) UserRepository {
	return &UserRepositoryImpl{
		dao: dao,
	}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, user user_domain.User) error {
	return repo.dao.Insert(ctx, repo.domainToEntity(user))
}

func (repo *UserRepositoryImpl) domainToEntity(u user_domain.User) user_dao.User {
	return user_dao.User{
		CreateTime: u.CreateTime.UnixMilli(),
		UpdateTime: u.UpdateTime.UnixMilli(),
		Email:      u.Email,
		UserID:     u.UserID,
		Password:   u.Password,
		UserInfo: user_dao.UserInfo{
			CreateTime:        u.UserInfo.CreateTime.UnixMilli(),
			UpdateTime:        u.UserInfo.UpdateTime.UnixMilli(),
			UserID:            u.UserInfo.UserID,
			Nickname:          u.UserInfo.Nickname,
			Avatar:            u.UserInfo.Avatar,
			Signature:         u.UserInfo.Signature,
			Birthday:          u.UserInfo.Birthday.UnixMilli(),
			Sex:               u.UserInfo.Sex,
			PointsRemaining:   u.UserInfo.PointsRemaining,
			NumberOfPosts:     u.UserInfo.NumberOfPosts,
			NumberOfResources: u.UserInfo.NumberOfResources,
		},
		UserRole: user_dao.UserRole{
			CreateTime: u.UserRole.CreateTime.UnixMilli(),
			UpdateTime: u.UserRole.UpdateTime.UnixMilli(),
			UserID:     u.UserRole.UserID,
			IsAdmin:    u.UserRole.IsAdmin,
			IsOrder:    u.UserRole.IsOrder,
			IsDisabled: u.UserRole.IsDisabled,
		},
	}
}
