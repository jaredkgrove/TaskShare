package taskList

import (
	"context"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
)

type Reader interface {
	Get(ctx context.Context, id entity.ID) (*entity.TaskList, error)
	List(ctx context.Context, userId entity.ID) (*[]entity.TaskList, error)
}

type Writer interface {
	Create(ctx context.Context, e *entity.TaskList) (entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get(ctx context.Context, id entity.ID) (*entity.TaskList, error)
	List(ctx context.Context, userId entity.ID) (*[]entity.TaskList, error)
	Create(ctx context.Context, e *entity.TaskList) (entity.ID, error)

	// CreateTaskList(kind, id, etag, title, updated, selfLink, position, status, due string) (entity.ID, error)
}
