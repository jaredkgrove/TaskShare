package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"google.golang.org/api/iterator"
	googleTasks "google.golang.org/api/tasks/v1"
)

//BookMySQL mysql repo
type TaskListFirestore struct {
	Client *firestore.Client
}

//NewTaskFirestore create new repository
func NewTaskListFirestore(client *firestore.Client) *TaskListFirestore {
	return &TaskListFirestore{
		Client: client,
	}
}

func (r *TaskListFirestore) Get(ctx context.Context, id entity.ID) (*entity.TaskList, error) {
	dsnap, err := r.Client.Collection("tasks").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var t entity.TaskList
	dsnap.DataTo(&t)

	return &t, nil
}

func (r *TaskListFirestore) FindByGoogleTaskListAndUser(ctx context.Context, googleTaskList *googleTasks.TaskList, userID string) (*entity.TaskList, error) {
	iter := r.Client.Collection("taskLists").Where(fmt.Sprint("users.", googleTaskList.Id), "==", userID).Limit(1).Documents(ctx)
	defer iter.Stop()
	var taskList entity.TaskList

	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if err := doc.DataTo(&taskList); err != nil {
		return nil, err
	}

	return &taskList, nil
}

func (r *TaskListFirestore) List(ctx context.Context, userId entity.ID) (*[]entity.TaskList, error) {
	iter := r.Client.Collection("taskLists").Documents(ctx)
	defer iter.Stop()
	var taskLists []entity.TaskList
	for {
		var tl entity.TaskList
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
			fmt.Println(err)
		}

		if err := doc.DataTo(&tl); err != nil {
			fmt.Println(err)
			continue
		}
		// tl.FirestoreID = doc.Ref.ID
		taskLists = append(taskLists, tl)
	}

	return &taskLists, nil
}

//Create a task
func (r *TaskListFirestore) Create(ctx context.Context, e *entity.TaskList) (entity.ID, error) {
	doc, result, err := r.Client.Collection("taskLists").Add(ctx, e)
	if err != nil {
		fmt.Println(result)
		return "", err
	}
	return doc.ID, nil
}

func (r *TaskListFirestore) CreateFromGoogleTaskList(ctx context.Context, googleTaskList *googleTasks.TaskList, userID string) (entity.ID, error) {
	doc, wr, err := r.Client.Collection("taskLists").Add(ctx, map[string]interface{}{
		"title":  googleTaskList.Title,
		"users": map[string]interface{}{googleTaskList.Id: userID},
	})
	if err != nil {
		fmt.Println(wr)
		return "", err
	}
	return doc.ID, nil
}


