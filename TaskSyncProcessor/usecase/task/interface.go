package task

import (
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"context"
)

type Reader interface{
	Get(ctx context.Context, id entity.ID)(*entity.Task, error)
}

type Writer interface {
	Create(e *entity.Task) (entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetTask(id entity.ID) (*entity.Task, error)
	CreateTask(kind, id, etag, title, updated, selfLink, position, status, due string) (entity.ID, error)
}
