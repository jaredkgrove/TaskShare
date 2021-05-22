package taskList

import (
	"context"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	googleTasks "google.golang.org/api/tasks/v1"
)

type Reader interface {
	Get(ctx context.Context, id entity.ID) (*entity.TaskList, error)
	FindByGoogleTaskListIDAndUserID(ctx context.Context, googleTaskListID string, userID string) (*entity.TaskList, error)
	List(ctx context.Context, userId entity.ID) (*[]entity.TaskList, error)
}

type Writer interface {
	Create(ctx context.Context, e *entity.TaskList) (entity.ID, error)
	CreateFromGoogleTaskList(ctx context.Context, googleTaskList *googleTasks.TaskList, userID string) (*entity.TaskList, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get(ctx context.Context, id entity.ID) (*entity.TaskList, error)
	List(ctx context.Context, userId entity.ID) (*[]entity.TaskList, error)
	SaveFromGoogleTaskList(ctx context.Context, googleTaskList *googleTasks.TaskList, userID string) (*entity.TaskList, error)
}
