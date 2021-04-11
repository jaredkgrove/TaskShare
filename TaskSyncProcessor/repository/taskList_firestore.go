package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"google.golang.org/api/iterator"
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
	fmt.Println(err)
	fmt.Println(result)
	fmt.Println(doc.ID)
	return doc.ID, nil
}

//Get a task
// func (r *TaskFirestore) Get(ctx context.Context, id entity.ID) (*entity.TaskList, error) {
// 	dsnap, err := r.Client.Collection("tasks").Doc(id).Get(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var t entity.Task
// 	dsnap.DataTo(&t)

// 	return &t, nil
// }

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
