package googleUser

import (
	"context"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// func (s *Service) CreateTask(kind, id, etag, title, updated, selfLink, position, status, due string) (entity.ID, error) {
// 	t, err := entity.NewTask(kind, id, etag, title, updated, selfLink, position, status, due)
// 	if err != nil {
// 		return t.ID, err
// 	}
// 	return s.repo.Create(t)
// }

func (s *Service) GetGoogleUsers(ctx context.Context) (*[]entity.GoogleUser, error) {
	gu, err := s.repo.GetGoogleUsers(ctx)
	if gu == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return gu, nil
}
