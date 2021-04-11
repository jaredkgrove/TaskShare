package entity

type TaskList struct {
	// Kind           string
	// ID             string
	// Due            string
	Etag           string
	googleID       string
	googleTaskLink string
	Title          string
	UserID         string
	// Updated        string
	// Position       string
}

// func NewTaskList(etag, title, selfLink, position, status, due string) (*TaskList, error) {
// 	t := &TaskList{

// 		Etag:  etag,
// 		Title: title,
// 		// Updated: updated,
// 		googleTaskLink: selfLink,
// 		// Position:       position,
// 	}
// 	err := t.Validate()
// 	return t, err
// }

// func (t *TaskList) Validate() error {
// 	if false {
// 		return ErrInvalidEntity
// 	}
// 	return nil
// }
