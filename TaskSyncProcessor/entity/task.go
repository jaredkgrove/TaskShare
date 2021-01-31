package entity

type Task struct {
	Kind     string
	ID       string
	Etag     string
	Title    string
	Updated  string
	SelfLink string
	Position string
	Status   string
	Due      string
}

func NewTask(kind, id, etag, title, updated, selfLink, position, status, due string) (*Task, error) {
	t := &Task{
		Kind: kind,
		ID: id,
		Etag: etag,
		Title: title,
		Updated: updated,
		SelfLink: selfLink,
		Position: position,
		Status: status,
		Due: due,
	}
	err := t.Validate()
	return t, err
}

func (t *Task)Validate() error{
	if false{
		return ErrInvalidEntity
	}
	return nil
}