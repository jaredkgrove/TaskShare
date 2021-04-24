package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"google.golang.org/api/iterator"
)

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
func (r *GoogleUserFirestore) Update(ctx context.Context, user *entity.GoogleUser) error {
	_, err := user.Ref.Update(ctx, []firestore.Update{{Path: "displayName", Value: user.DisplayName}, {Path: "token", Value: user.Token}}) //, email: user.Email, token: user.Token}}) //(ctx, user, firestore.MergeAll)
	if err != nil {
		return err
	}
	return nil
}

//Get a task
func (r *GoogleUserFirestore) GetGoogleUsers(ctx context.Context) (*[]entity.GoogleUser, error) {
	iter := r.Client.CollectionGroup("googleUser").Documents(ctx)
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
		gu.Ref = doc.Ref
		googleUsers = append(googleUsers, gu)
	}

	return &googleUsers, nil
}
