package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	ErrCodeSendTooMany        = errors.New("发送验证码太频繁")
	ErrCodeVerifyTooManyTimes = errors.New("验证次数太多")
	ErrUnknownForCode         = errors.New("我也不知发生什么了，反正是跟 code 有关")
)

//go:embed lua/set_code.lua
var luaSetCode string

//go:embed lua/verify_code.lua
var luaVerifyCode string

type CodeCache struct {
	client redis.Cmdable
}

func NewCodeCache(client redis.Cmdable) *CodeCache {
	return &CodeCache{
		client: client,
	}
}

func (c *CodeCache) Set(ctx context.Context, biz, email, code string) error {
	res, err := c.client.Eval(ctx, luaSetCode, []string{c.key(biz, email)}, code).Int()
	if err != nil {
		return err
	}
	switch res {
	case 0:
		// 毫无问题
		return nil
	case -1:
		// 发送太频繁
		return ErrCodeSendTooMany
	default:
		// 系统错误
		return errors.New("系统错误")
	}
}

func (c *CodeCache) Verify(ctx context.Context, biz, email, inputCode string) (bool, error) {
	res, err := c.client.Eval(ctx, luaVerifyCode, []string{c.key(biz, email)}, inputCode).Int()
	if err != nil {
		return false, err
	}

	switch res {
	case 0:
		// 毫无问题
		return true, nil
	case -1:
		// 发送太频繁
		return false, ErrCodeVerifyTooManyTimes
	default:
		// 系统错误
		return false, ErrUnknownForCode
	}
}

func (c *CodeCache) key(biz, email string) string {
	return fmt.Sprintf("email_code:%s:%s", biz, email)
}
