package task

import (
	"context"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	googleTasks "google.golang.org/api/tasks/v1"
)

type Reader interface {
	Get(ctx context.Context, id entity.ID) (*entity.Task, error)
	FindByTaskListGoogleTaskIDAndUserID(ctx context.Context, taskList *entity.TaskList, googleTaskID string, userID string) (*entity.Task, error)
}

type Writer interface {
	Create(ctx context.Context, e *entity.Task) (entity.ID, error)
	CreateFromGoogleTask(ctx context.Context, googleTask *googleTasks.Task, taskList *entity.TaskList, userID string) (entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetTask(ctx context.Context, id entity.ID) (*entity.Task, error)
	CreateTask(ctx context.Context, kind, id, etag, title, updated, selfLink, position, status, due string) (entity.ID, error)
	SaveFromGoogleTask(ctx context.Context, googleTaskList *googleTasks.TaskList, taskList entity.TaskList, userID string) (entity.ID, error)
}
