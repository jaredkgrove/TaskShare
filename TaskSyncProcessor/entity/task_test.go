package entity_test

import (
	"testing"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	tsk, err := entity.NewTask("testKind", "1", "etag1", "Test Task", "updated date", "self link", "position", "status", "a date that it is due")
	assert.Nil(t, err)
	assert.Equal(t, tsk.Title, "Test Task")
	assert.NotNil(t, tsk.ID)
}


// func TestTaskValidate(t *testing.T) {
// 	type test struct {
// 		title    string
// 		author   string
// 		pages    int
// 		quantity int
// 		want     error
// 	}

// 	tests := []test{
// 		{
// 			title:    "American Gods",
// 			author:   "Neil Gaiman",
// 			pages:    100,
// 			quantity: 1,
// 			want:     nil,
// 		},
// 		{
// 			title:    "American Gods",
// 			author:   "Neil Gaiman",
// 			pages:    100,
// 			quantity: 0,
// 			want:     entity.ErrInvalidEntity,
// 		},
// 		{
// 			title:    "",
// 			author:   "Neil Gaiman",
// 			pages:    100,
// 			quantity: 1,
// 			want:     entity.ErrInvalidEntity,
// 		},
// 		{
// 			title:    "American Gods",
// 			author:   "",
// 			pages:    100,
// 			quantity: 1,
// 			want:     entity.ErrInvalidEntity,
// 		},
// 		{
// 			title:    "American Gods",
// 			author:   "Neil Gaiman",
// 			pages:    0,
// 			quantity: 1,
// 			want:     entity.ErrInvalidEntity,
// 		},
// 	}
// 	for _, tc := range tests {

// 		_, err := entity.NewBook(tc.title, tc.author, tc.pages, tc.quantity)
// 		assert.Equal(t, err, tc.want)
// 	}

// }