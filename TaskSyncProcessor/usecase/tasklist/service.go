package taskList

import (
	"context"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/usecase/task"
	googleTasks "google.golang.org/api/tasks/v1"
)

type Service struct {
	taskListRepo Repository
	taskRepo     task.Repository
}

func NewService(r Repository, tr task.Repository) *Service {
	return &Service{
		taskListRepo: r,
		taskRepo:     tr,
	}
}

func (s *Service) Get(ctx context.Context, id entity.ID) (*entity.TaskList, error) {
	t, err := s.taskListRepo.Get(ctx, id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) SaveFromGoogleTaskList(ctx context.Context, googleTaskList *googleTasks.TaskList, userID string) error {
	t, err := s.taskListRepo.FindByGoogleTaskListAndUser(ctx, googleTaskList, userID)
	if err != nil {
		return err
	}
	if (t == nil){
		s.taskListRepo.CreateFromGoogleTaskList(ctx, googleTaskList, userID)
	}

	return nil
}

func (s *Service) List(ctx context.Context,userId entity.ID) (*[]entity.TaskList, error) {
	t, err := s.taskListRepo.List(ctx, userId)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return t, nil
}
