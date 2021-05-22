package entity

import "cloud.google.com/go/firestore"

type TaskList struct {
	Ref               *firestore.DocumentRef
	Etag              string
	Title             string
	UserGoogleMapping map[string]interface{}
	Users             []string
}

// func NewTaskList(etag, title, selfLink, position, status, due string) (*TaskList, error) {
// 	t := &TaskList{

// 		Etag:  etag,
// 		Title: title,
// 		// Updated: updated,
// 		googleTaskLink: selfLink,
// 		// Position:       position,
// 	}
// }
// func NewFromGoogleTaskListAndUserID(googleTaskList googleTasks.TaskList, userID string) (*TaskList, error) {
// 	t := &TaskList{

// 		Etag:  googleTaskList.Etag,
// 		Title: googleTaskList.Title,
// 		user: googleTaskList.Id
// 	}
// }

// 	err := t.Validate()
// 	return t, err
// }

// func (t *TaskList) Validate() error {
// 	if false {
// 		return ErrInvalidEntity
// 	}
// 	return nil
// }
