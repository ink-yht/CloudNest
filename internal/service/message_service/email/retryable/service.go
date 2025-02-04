package retryable

//type Service struct {
//	msgSvc *tencent.Service
//	// 重试
//	retryCnt int
//}
//
//func (s Service) Send(ctx context.Context, email string, subject, body string) error {
//	err := s.msgSvc.Send(ctx, email, subject, body)
//	for err != nil && s.retryCnt < 10 {
//		err = s.msgSvc.Send(ctx, email, subject, body)
//		s.retryCnt++
//		return
//	}
//}
