package task

import (
	"context"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	googleTasks "google.golang.org/api/tasks/v1"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateTask(ctx context.Context, kind, id, etag, title, updated, selfLink, position, status, due string) (entity.ID, error) {
	t, err := entity.NewTask(kind, id, etag, title, updated, selfLink, position, status, due)
	if err != nil {
		return t.ID, err
	}
	return s.repo.Create(ctx, t)
}

func (s *Service) GetTask(ctx context.Context, id entity.ID) (*entity.Task, error) {
	t, err := s.repo.Get(ctx, id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) SaveFromGoogleTask(ctx context.Context, googleTask *googleTasks.Task, taskList *entity.TaskList, userID string) error {
	t, err := s.repo.FindByTaskListGoogleTaskIDAndUserID(ctx, taskList, googleTask.Id, userID)
	if err != nil {
		return err
	}
	if t == nil {
		s.repo.CreateFromGoogleTask(ctx, googleTask, taskList, userID) //TODO
	}

	return nil
}
