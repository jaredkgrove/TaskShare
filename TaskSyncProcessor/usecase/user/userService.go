package user

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

func (s *Service) GetUsers(ctx context.Context) (*[]entity.User, error) {
	t, err := s.repo.GetUsers(ctx)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}
