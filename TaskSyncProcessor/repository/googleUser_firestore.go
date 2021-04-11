package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"google.golang.org/api/iterator"
)

//BookMySQL mysql repo
type GoogleUserFirestore struct {
	Client *firestore.Client
}

//NewTaskFirestore create new repository
func NewGoogleUserFirestore(client *firestore.Client) *GoogleUserFirestore {
	return &GoogleUserFirestore{
		Client: client,
	}
}

//Create a task
func (r *GoogleUserFirestore) Create(e *entity.GoogleUser) (entity.ID, error) {
	return "not implemented", nil
}

//Get a task
func (r *GoogleUserFirestore) GetGoogleUsers(ctx context.Context) (*[]entity.GoogleUser, error) {
	iter := r.Client.CollectionGroup("googleUsers").Documents(ctx)
	defer iter.Stop()
	var googleUsers []entity.GoogleUser
	for {
		var gu entity.GoogleUser
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
			fmt.Println(err)
		}

		if err := doc.DataTo(&gu); err != nil {
			fmt.Println(err)
			continue
		}
		gu.UserId = doc.Ref.Parent.Parent.ID
		googleUsers = append(googleUsers, gu)
	}

	return &googleUsers, nil
}
