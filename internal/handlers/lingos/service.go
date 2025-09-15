package lingos

import (
	"context"
)

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) CreateLingo(ctx context.Context, l Lingo) (*Lingo, error) {
	return s.repo.Insert(ctx, l)
}

func (s *Service) GetLingo(ctx context.Context, lang Lang, key string) (*Lingo, error) {
	return s.repo.GetByLangAndKey(ctx, lang, key)
}

func (s *Service) DeleteLingo(ctx context.Context, lang Lang, key string) error {
	return s.repo.DeleteByLangAndKey(ctx, lang, key)
}
