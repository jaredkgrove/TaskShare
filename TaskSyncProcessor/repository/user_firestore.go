package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"google.golang.org/api/iterator"
)

//BookMySQL mysql repo
type UserFirestore struct {
	Client *firestore.Client
}

//NewTaskFirestore create new repository
func NewUserFirestore(client *firestore.Client) *UserFirestore {
	return &UserFirestore{
		Client: client,
	}
}

//Create a task
func (r *UserFirestore) Create(e *entity.User) (entity.ID, error) {
	return "not implemented", nil
}

//Get a task
func (r *UserFirestore) GetUsers(ctx context.Context) (*[]entity.User, error) {
	iter := r.Client.Collection("users").Documents(ctx)
	defer iter.Stop()
	var users []entity.User
	for {
		var u entity.User
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
			fmt.Println(err)
		}

		if err := doc.DataTo(&u); err != nil {
			fmt.Println(err)
			continue
		}
		u.FirestoreID = doc.Ref.ID
		users = append(users, u)
	}

	return &users, nil
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
