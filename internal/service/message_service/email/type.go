package email

import "context"

type Service interface {
	Send(ctx context.Context, email string, subject, body string) error
}
