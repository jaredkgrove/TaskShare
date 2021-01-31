package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
)

//BookMySQL mysql repo
type TaskFirestore struct {
	Client *firestore.Client
}

//NewBookMySQL create new repository
func NewTaskFirestore(client *firestore.Client) *TaskFirestore {
	return &TaskFirestore{
		Client: client,
	}
}

//Create a book
func (r *TaskFirestore) Create(e *entity.Task) (entity.ID, error) {
	return "not implemented", nil
}

//Get a book
func (r *TaskFirestore) Get(ctx context.Context, id entity.ID) (*entity.Task, error) {
	dsnap, err := r.Client.Collection("tasks").Doc("GZhp0EXbLLLFWpzJRRHX").Get(ctx)
	if err != nil {
		return nil, err
	}
	var t entity.Task
	dsnap.DataTo(&t)

	return &t, nil
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
