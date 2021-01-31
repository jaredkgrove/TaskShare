package tasklist

type TaskList struct {
	kind          string
	etag          string
	nextPageToken string
	items         []Task
}