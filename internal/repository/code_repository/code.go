package code_repository

import (
	"context"
	"github.com/ink-yht/CloudNest/internal/repository/cache"
)

var (
	ErrCodeSendTooMany        = cache.ErrCodeSendTooMany
	ErrCodeVerifyTooManyTimes = cache.ErrCodeVerifyTooManyTimes
)

type CodeRepository struct {
	cache *cache.CodeCache
}

func NewCodeRepository(c *cache.CodeCache) *CodeRepository {
	return &CodeRepository{
		cache: c,
	}
}

func (repo *CodeRepository) Store(ctx context.Context, biz string, email string, code string) error {
	return repo.cache.Set(ctx, biz, email, code)
}

func (repo *CodeRepository) Verify(ctx context.Context, biz, email, inputCode string) (bool, error) {
	return repo.cache.Verify(ctx, biz, email, inputCode)
}
