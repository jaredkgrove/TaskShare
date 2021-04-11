package taskList

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

func (s *Service) Get(ctx context.Context, id entity.ID) (*entity.TaskList, error) {
	t, err := s.repo.Get(ctx, id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) Create(ctx context.Context, e *entity.TaskList) (*entity.ID, error) {
	t, err := s.repo.Create(ctx, e)
	// if t == nil {
	// 	return nil, entity.ErrNotFound
	// }
	// if err != nil {
	// 	return nil, err
	// }

	return &t, err
}

func (s *Service) List(ctx context.Context, userId entity.ID) (*[]entity.TaskList, error) {
	t, err := s.repo.List(ctx, userId)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}
