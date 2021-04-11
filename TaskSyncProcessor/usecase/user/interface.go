package user

import (
	"context"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
)

type Reader interface {
	GetUsers(ctx context.Context) (*[]entity.User, error)
}

type Writer interface {
	Create(e *entity.User) (entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUsers() (*[]entity.User, error)
	// CreateTaskList(kind, id, etag, title, updated, selfLink, position, status, due string) (entity.ID, error)
}
