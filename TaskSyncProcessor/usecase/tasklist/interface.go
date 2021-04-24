package taskList

import (
	"context"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	googleTasks "google.golang.org/api/tasks/v1"
)

type Reader interface {
	Get(ctx context.Context, id entity.ID) (*entity.TaskList, error)
	FindByGoogleTaskListAndUser(ctx context.Context, googleTaskList *googleTasks.TaskList, userID string) (*entity.TaskList, error)
	List(ctx context.Context, userId entity.ID) (*[]entity.TaskList, error)
}

type Writer interface {
	Create(ctx context.Context, e *entity.TaskList) (entity.ID, error)
	CreateFromGoogleTaskList(ctx context.Context, googleTaskList *googleTasks.TaskList, userID string) (entity.ID, error)
	// Update(ctx context.Context, e *entity.TaskList) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get(ctx context.Context, id entity.ID) (*entity.TaskList, error)
	List(ctx context.Context, userId entity.ID) (*[]entity.TaskList, error)
	SaveFromGoogleTaskList(ctx context.Context, googleTaskList *googleTasks.TaskList, userID string) (entity.ID, error)
}
