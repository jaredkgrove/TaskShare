package task

import (
	// "strings"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Task
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[entity.ID]*entity.Task{}
	return &inmem{
		m: m,
	}
}

func (r *inmem) Create(e *entity.Task) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *inmem) Get(id entity.ID) (*entity.Task, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a book
// func (r *inmem) Update(e *entity.Task) error {
// 	_, err := r.Get(e.ID)
// 	if err != nil {
// 		return err
// 	}
// 	r.m[e.ID] = e
// 	return nil
// }

//Search books
// func (r *inmem) Search(query string) ([]*entity.Book, error) {
// 	var d []*entity.Book
// 	for _, j := range r.m {
// 		if strings.Contains(strings.ToLower(j.Title), query) {
// 			d = append(d, j)
// 		}
// 	}
// 	return d, nil
// }

//List books
// func (r *inmem) List() ([]*entity.Book, error) {
// 	var d []*entity.Book
// 	for _, j := range r.m {
// 		d = append(d, j)
// 	}
// 	return d, nil
// }

//Delete a book
// func (r *inmem) Delete(id entity.ID) error {
// 	if r.m[id] == nil {
// 		return entity.ErrNotFound
// 	}
// 	r.m[id] = nil
// 	return nil
// }