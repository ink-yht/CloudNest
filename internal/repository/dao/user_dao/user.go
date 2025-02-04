package user_dao

import (
	"context"
	"database/sql"
	"time"
)

type UserDao interface {
	Insert(ctx context.Context, u User) error
}

type UserDAOImpl struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) UserDao {
	return &UserDAOImpl{db: db}
}

// Insert 注册
func (dao *UserDAOImpl) Insert(ctx context.Context, u User) error {
	// 毫秒
	now := time.Now().UnixMilli()
	u.CreateTime = now
	u.UpdateTime = now
	u.UserInfo.CreateTime = now
	u.UserInfo.UpdateTime = now
	u.UserRole.CreateTime = now
	u.UserRole.UpdateTime = now

	//err := dao.db.WithContext(ctx).Preload("UserConf").Create(&u).Error
	//
	//// 如果错误是MySQL错误类型
	//var me *mysql.MySQLError
	//if errors.As(err, &me) {
	//	const duplicateErr uint16 = 1062
	//	if me.Number == duplicateErr {
	//		// 邮箱冲突
	//		return ErrDuplicate
	//	}
	//}

	// 系统错误
	return nil
}
