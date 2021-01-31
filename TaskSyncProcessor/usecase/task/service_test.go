package task_test

import (
	"testing"
	// "time"

	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/usecase/task"
	"github.com/stretchr/testify/assert"
)

func newFixtureTask() *entity.Task {
	return &entity.Task{
		Kind:		"A Kind",
		ID:       	"An ID",
		Etag:     	"An Etag",
		Title:    	"A Title",
		Updated:  	"An Updated",
		SelfLink: 	"A SelfLink",
		Position: 	"A Position",
		Status:   	"A Status",
		Due:      	"A Due",
	}
}

func Test_Create(t *testing.T) {
	repo := task.NewInmem()
	m := task.NewService(repo)
	u := newFixtureTask()
	_, err := m.CreateTask(u.Kind, u.ID, u.Etag, u.Title, u.Updated, u.SelfLink, u.Position, u.Status, u.Due)
	assert.Nil(t, err)
	assert.NotNil(t, u.Updated)
}

func Test_SearchAndFind(t *testing.T) {
	repo := task.NewInmem()
	m := task.NewService(repo)
	
	u1 := newFixtureTask()
	// u2 := newFixtureBook()
	// u2.Title = "Lemmy: Biography"

	uID, _ := m.CreateTask(u1.Kind, u1.ID, u1.Etag, u1.Title, u1.Updated, u1.SelfLink, u1.Position, u1.Status, u1.Due)
	// _, _ = m.CreateBook(u2.Title, u2.Author, u2.Pages, u2.Quantity)

	// t.Run("search", func(t *testing.T) {
	// 	c, err := m.SearchBooks("ozzy")
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, 1, len(c))
	// 	assert.Equal(t, "I Am Ozzy", c[0].Title)

	// 	c, err = m.SearchBooks("dio")
	// 	assert.Equal(t, entity.ErrNotFound, err)
	// 	assert.Nil(t, c)
	// })
	// t.Run("list all", func(t *testing.T) {
	// 	all, err := m.ListBooks()
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, 2, len(all))
	// })

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetTask(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Title, saved.Title)
	})
}

// func Test_Update(t *testing.T) {
// 	repo := newInmem()
// 	m := NewService(repo)
// 	u := newFixtureBook()
// 	id, err := m.CreateBook(u.Title, u.Author, u.Pages, u.Quantity)
// 	assert.Nil(t, err)
// 	saved, _ := m.GetBook(id)
// 	saved.Title = "Lemmy: Biography"
// 	assert.Nil(t, m.UpdateBook(saved))
// 	updated, err := m.GetBook(id)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "Lemmy: Biography", updated.Title)
// }

// func TestDelete(t *testing.T) {
// 	repo := newInmem()
// 	m := NewService(repo)
// 	u1 := newFixtureBook()
// 	u2 := newFixtureBook()
// 	u2ID, _ := m.CreateBook(u2.Title, u2.Author, u2.Pages, u2.Quantity)

// 	err := m.DeleteBook(u1.ID)
// 	assert.Equal(t, entity.ErrNotFound, err)

// 	err = m.DeleteBook(u2ID)
// 	assert.Nil(t, err)
// 	_, err = m.GetBook(u2ID)
// 	assert.Equal(t, entity.ErrNotFound, err)
// }