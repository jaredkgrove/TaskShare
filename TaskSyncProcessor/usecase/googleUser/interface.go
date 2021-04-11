package googleUser

import (
	"context"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
)

type Reader interface {
	GetGoogleUsers(ctx context.Context) (*[]entity.GoogleUser, error)
}

type Writer interface {
	Create(e *entity.GoogleUser) (entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetGoogleUsers() (*[]entity.GoogleUser, error)
	// CreateTaskList(kind, id, etag, title, updated, selfLink, position, status, due string) (entity.ID, error)
}
