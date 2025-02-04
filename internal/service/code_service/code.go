package code_service

import (
	"fmt"
	"github.com/ink-yht/CloudNest/internal/repository/code_repository"
	"github.com/ink-yht/CloudNest/internal/service/message_service/email/tencent"
	"golang.org/x/net/context"
	"math/rand"
)

type CodeService interface {
	Send(ctx context.Context, email string, biz string) error
	Verify(ctx context.Context, biz string, email string, inputCode string) (bool, error)
}

type CodeServiceImpl struct {
	repo   *code_repository.CodeRepository
	msgSvc *tencent.Service
}

func NewCodeService(repo *code_repository.CodeRepository, msgSvc *tencent.Service) CodeService {
	return &CodeServiceImpl{
		repo:   repo,
		msgSvc: msgSvc,
	}
}

// Send 发送验证码 biz: 区别业务场景
func (svc *CodeServiceImpl) Send(ctx context.Context, email string, biz string) error {
	// 生成验证码
	code := svc.generateCode()
	// 塞进去 redis
	err := svc.repo.Store(ctx, biz, email, code)
	if err != nil {
		// 有问题
		return err
	}
	// 发送出去
	err = svc.msgSvc.Send(ctx, email, "验证码", code)
	return err
}

func (svc *CodeServiceImpl) Verify(ctx context.Context, biz string, email string, inputCode string) (bool, error) {
	return svc.repo.Verify(ctx, biz, email, inputCode)
}

func (svc *CodeServiceImpl) generateCode() string {
	// 六位数，num 在 0, 999999 之间，包含 0 和 999999
	num := rand.Intn(1000000)
	// 不够六位的，加上前导 0
	// 000001
	return fmt.Sprintf("%06d", num)
}
