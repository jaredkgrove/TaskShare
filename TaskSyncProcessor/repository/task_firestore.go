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
type TaskFirestore struct {
	Client *firestore.Client
}

//NewTaskFirestore create new repository
func NewTaskFirestore(client *firestore.Client) *TaskFirestore {
	return &TaskFirestore{
		Client: client,
	}
}

//Create a task
func (r *TaskFirestore) Create(ctx context.Context, t *entity.Task) (entity.ID, error) {
	return "not implemented", nil
}

//Get a task
func (r *TaskFirestore) Get(ctx context.Context, id entity.ID) (*entity.Task, error) {
	dsnap, err := r.Client.Collection("tasks").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var t entity.Task
	dsnap.DataTo(&t)

	return &t, nil
}

func (r *TaskFirestore) FindByTaskListGoogleTaskIDAndUserID(ctx context.Context, taskList *entity.TaskList, googleTaskID string, userID string) (*entity.Task, error) {

	iter := taskList.Ref.Collection("tasks").Where("googleTaskID", "==", googleTaskID).Limit(1).Documents(ctx)
	defer iter.Stop()
	var task entity.Task

	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if err := doc.DataTo(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskFirestore) CreateFromGoogleTask(ctx context.Context, googleTask *googleTasks.Task, taskList *entity.TaskList, userID string) (entity.ID, error) {

	//GooglTaskListID is not the firestore taskListID
	// iter := r.Client.Collection("taskLists").Where(fmt.Sprint("userGoogleMapping.", googleTaskListID), "==", userID).Limit(1).Documents(ctx)

	// defer iter.Stop()

	// taskListDsnap, err := iter.Next()

	// var taskList entity.TaskList

	// if err == iterator.Done {
	// 	return "", nil
	// }
	// if err != nil {
	// 	return "", err
	// }

	// if err := taskListDsnap.DataTo(&taskList); err != nil {
	// 	return "", err
	// }

	//FIND task where userTaskMapping key of userID == googleTaskID OR key of userID == nil AND title == googlTask.Title
	//If it exists then update (TODO: determine requirements)
	//If not then create
	// var task entity.Task

	iter := taskList.Ref.Collection("tasks").Where(fmt.Sprint("userTaskMapping.", userID), "==", googleTask.Id).Limit(1).Documents(ctx)
	taskDsnap, err := iter.Next()
	if err != iterator.Done {
		fmt.Println(taskDsnap.Ref.ID)
		//update and return if update criteria met
		return "", err
	}

	iter = taskList.Ref.Collection("tasks").Where("title", "==", googleTask.Title).Where(fmt.Sprint("userTaskMapping.", userID), "==", "").Limit(1).Documents(ctx)
	taskDsnap, err = iter.Next()
	if err != iterator.Done {
		fmt.Println(taskDsnap.Ref.ID)
		//update and return if update criteria met
		wr1, err := taskList.Ref.Collection("tasks").Doc(taskDsnap.Ref.ID).Set(ctx, map[string]interface{}{
			"title":           googleTask.Title,
			"userID":          userID,
			"userTaskMapping": map[string]interface{}{userID: googleTask.Id},
		}) //TODO merge true?
		if err != nil {
			fmt.Println(wr1)
			return "", err
		}

		return taskDsnap.Ref.ID, err
	}

	doc, wr, err := taskList.Ref.Collection("tasks").Add(ctx, map[string]interface{}{
		"title":           googleTask.Title,
		"userID":          userID,
		"userTaskMapping": map[string]interface{}{userID: googleTask.Id},
	})

	if err != nil {
		fmt.Println(wr)
		return "", err
	}
	return doc.ID, nil
}

//Update a book
// func (r *BookMySQL) Update(e *entity.Book) error {
// 	e.UpdatedAt = time.Now()
// 	_, err := r.db.Exec("update book set title = ?, author = ?, pages = ?, quantity = ?, updated_at = ? where id = ?", e.Title, e.Author, e.Pages, e.Quantity, e.UpdatedAt.Format("2006-01-02"), e.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//Search books
// func (r *BookMySQL) Search(query string) ([]*entity.Book, error) {
// 	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book where title like ?`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var books []*entity.Book
// 	rows, err := stmt.Query("%" + query + "%")
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var b entity.Book
// 		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		books = append(books, &b)
// 	}

// 	return books, nil
// }

//List books
// func (r *BookMySQL) List() ([]*entity.Book, error) {
// 	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var books []*entity.Book
// 	rows, err := stmt.Query()
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var b entity.Book
// 		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		books = append(books, &b)
// 	}
// 	return books, nil
// }

//Delete a book
// func (r *BookMySQL) Delete(id entity.ID) error {
// 	_, err := r.db.Exec("delete from book where id = ?", id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
