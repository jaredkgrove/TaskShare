package main

import (
	"database/sql"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	// "net/http"

	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// type taskResponse struct {
// 	Kind          string `json:"kind"`
// 	Etag          string `json:"etag"`
// 	NextPageToken string `json:"nextPageToken"`
// 	Items         []task `json:"items"`
// }

// type taskList struct {
// 	kind          string
// 	etag          string
// 	nextPageToken string
// 	items         []task
// }

// type task struct {
// 	Kind     string `json:"kind"`
// 	ID       string `json:"id"`
// 	Etag     string `json:"etag"`
// 	Title    string `json:"title"`
// 	Updated  string `json:"updated"`
// 	SelfLink string `json:"selfLink"`
// 	Position string `json:"position"`
// 	Status   string `json:"status"`
// 	Due      string `json:"due"`
// 	Parent   string `json:"parent"`
// }

func InitializeSqlliteDB() {
	// sql.Register("sqlite3", &SQLiteDriver{}) //Not sure this is needed
	database, _ := sql.Open("sqlite3", "./TaskShareLite.db")
	statement1, _ := database.Prepare("DROP TABLE IF EXISTS task")
	statement1.Exec()

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS task (taskId INTEGER PRIMARY KEY, kind TEXT, id TEXT, etag TEXT, title TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO task (kind, id, etag, title) VALUES (?, ?, ?, ?)")
	statement.Exec("tasks#task", "bEItaXpWUF9wVlpleXN5VA", "\"LTIxMjU0NzE1Nzc\"", "Do Weekly Goals")
	rows, _ := database.Query("SELECT taskId, kind, id, etag, title FROM task")
	var taskId int
	var kind string
	var id string
	var etag string
	var title string
	for rows.Next() {
		rows.Scan(&taskId, &kind, &id, &etag, &title)
		fmt.Println(strconv.Itoa(taskId) + ": " + kind + " " + id + " " + etag + " " + title)
	}
}

// {
//     "kind": "tasks#tasks",
//     "etag": "\"LTIxMjU0NzE1NDY\"",
//     "nextPageToken": "CgwI_6zo-wUQmO2w2gISDAj_rOj7BRDwjqe2AhoQRjQ3LWZUY3JsRk8zaVEtaw==",
//     "items": [
//         {
//             "kind": "tasks#task",
//             "id": "bEItaXpWUF9wVlpleXN5VA",
//             "etag": "\"LTIxMjU0NzE1Nzc\"",
//             "title": "Do Weekly Goals",
//             "updated": "2020-12-20T18:01:04.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/bEItaXpWUF9wVlpleXN5VA",
//             "position": "00000000000000000000",
//             "status": "needsAction",
//             "due": "2020-12-20T00:00:00.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "NzVsZl91MExSMkc3UFJPUg",
//             "etag": "\"MTk5NDgzNTE3Mg\"",
//             "title": "Clean Bathroom",
//             "updated": "2020-12-18T17:30:04.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/NzVsZl91MExSMkc3UFJPUg",
//             "position": "00000000000000000003",
//             "status": "needsAction",
//             "due": "2020-12-18T00:00:00.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "VkF0eG1mZG9jLV85aFNuSA",
//             "etag": "\"MTk5NDgzNTE3Mg\"",
//             "title": "Sweep/Mop/Vacuum",
//             "updated": "2020-12-18T17:30:04.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/VkF0eG1mZG9jLV85aFNuSA",
//             "position": "00000000000000000001",
//             "status": "needsAction",
//             "due": "2020-12-18T00:00:00.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "aERESzZtYW5FUzBzcXRxRQ",
//             "etag": "\"MTk5NDgzNTE3Mg\"",
//             "title": "Grocery Shopping",
//             "updated": "2020-12-18T17:30:04.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/aERESzZtYW5FUzBzcXRxRQ",
//             "position": "00000000000000000002",
//             "status": "needsAction",
//             "due": "2020-12-18T00:00:00.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "M3hqMjJDUlVBYWpkSVlXTw",
//             "etag": "\"LTE4NjExOTI3MQ\"",
//             "title": "Pumpkin pie spice",
//             "updated": "2020-10-04T18:38:01.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/M3hqMjJDUlVBYWpkSVlXTw",
//             "parent": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000000",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "MUVDR2M5WHZrWEpvbF9ZZg",
//             "etag": "\"LTE4NjExOTg4OQ\"",
//             "title": "Pumpkin",
//             "updated": "2020-10-04T18:38:01.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/MUVDR2M5WHZrWEpvbF9ZZg",
//             "parent": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000001",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "ek8yMnZwdV9rX0dqd2xoUw",
//             "etag": "\"LTE4NjEyMTAxNA\"",
//             "title": "Jam",
//             "updated": "2020-10-04T18:38:00.000Z",

//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/ek8yMnZwdV9rX0dqd2xoUw",
//             "parent": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000002",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "ZUhadG5PbzhTdkN3elVKRQ",
//             "etag": "\"LTE4NjEyMTMwNQ\"",
//             "title": "Butter",
//             "updated": "2020-10-04T18:37:59.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/ZUhadG5PbzhTdkN3elVKRQ",
//             "parent": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000003",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "d2I5V3VBaFo2MXJWcDVfRw",
//             "etag": "\"LTE4NjEyMTkxOA\"",
//             "title": "Meat (sales)",
//             "updated": "2020-10-04T18:37:59.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/d2I5V3VBaFo2MXJWcDVfRw",
//             "parent": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000004",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "NHdRc2QzZHFFMldjcXpkdg",
//             "etag": "\"LTE4NjEyMzAwOQ\"",
//             "title": "Mixed greens",
//             "updated": "2020-10-04T18:37:58.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/NHdRc2QzZHFFMldjcXpkdg",
//             "parent": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000005",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "NFhFYVpLS1BJdkMtSFNKaw",
//             "etag": "\"LTE4NjEyMzY1OQ\"",
//             "title": "Frozen diced sweet potato",
//             "updated": "2020-10-04T18:37:57.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/NFhFYVpLS1BJdkMtSFNKaw",
//             "parent": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000006",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "Tk9YOW5sT185X3N5UDVydA",
//             "etag": "\"LTE4NjEyNDM4OQ\"",
//             "title": "Chicken strips",
//             "updated": "2020-10-04T18:37:56.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/Tk9YOW5sT185X3N5UDVydA",
//             "parent": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000007",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "aDRwbTlRQmJWZmhZYzRGUw",
//             "etag": "\"LTE4NjEyNDk2Ng\"",
//             "title": "Clean rain gutters",
//             "updated": "2020-10-04T18:37:56.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/aDRwbTlRQmJWZmhZYzRGUw",
//             "position": "00000000000000000004",
//             "status": "needsAction",
//             "due": "2020-10-10T00:00:00.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "MkhFLVZxdXVXSGhJRzh3Vw",
//             "etag": "\"LTE4NjEyNjE0NQ\"",
//             "title": "Groceries",
//             "updated": "2020-10-04T18:37:55.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/MkhFLVZxdXVXSGhJRzh3Vw",
//             "position": "00000000000000000005",
//             "status": "needsAction"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "djEwX0Z0cEJNUkc3UlNmNg",
//             "etag": "\"LTE4NjEyNjgyMg\"",
//             "title": "English muffins",
//             "updated": "2020-10-04T18:37:54.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/djEwX0Z0cEJNUkc3UlNmNg",
//             "position": "00000000000000000006",
//             "status": "completed",
//             "completed": "2020-10-04T18:37:54.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "eXhiSmp4c0ZaRThHd0x2dg",
//             "etag": "\"LTE4NjEyNzI2MA\"",
//             "title": "Chicken/bone broth for Kedzie",
//             "updated": "2020-10-04T18:37:54.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/eXhiSmp4c0ZaRThHd0x2dg",
//             "position": "00000000000000000007",
//             "status": "completed",
//             "completed": "2020-10-04T18:37:54.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "X1pfbHVYLS1xNExhZmFEeQ",
//             "etag": "\"LTE4NjEyNzM4NA\"",
//             "title": "Cereal",
//             "updated": "2020-10-04T18:37:53.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/X1pfbHVYLS1xNExhZmFEeQ",
//             "position": "00000000000000000008",
//             "status": "completed",
//             "completed": "2020-10-04T18:37:53.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "U0EteU1CRDYzX1VyaW1fbg",
//             "etag": "\"LTE4NjEyODcwMg\"",
//             "title": "Eggs",
//             "updated": "2020-10-04T18:37:52.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/U0EteU1CRDYzX1VyaW1fbg",
//             "position": "00000000000000000009",
//             "status": "completed",
//             "completed": "2020-10-04T18:37:52.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "NWJuR0hNaTdwR2JzWVlMeQ",
//             "etag": "\"LTE4NjEyOTM2Mg\"",
//             "title": "Cheddar slices",
//             "updated": "2020-10-04T18:37:52.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/NWJuR0hNaTdwR2JzWVlMeQ",
//             "position": "00000000000000000010",
//             "status": "completed",
//             "completed": "2020-10-04T18:37:52.000Z"
//         },
//         {
//             "kind": "tasks#task",
//             "id": "RjQ3LWZUY3JsRk8zaVEtaw",
//             "etag": "\"LTE4NjEyOTUxMA\"",
//             "title": "Pepper pack",
//             "updated": "2020-10-04T18:37:51.000Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/lists/RW9yQXktOXBOZ09rQS1rUw/tasks/RjQ3LWZUY3JsRk8zaVEtaw",
//             "position": "00000000000000000011",
//             "status": "completed",
//             "completed": "2020-10-04T18:37:51.000Z"
//         }
//     ]
// }

//lists

// {
//     "kind": "tasks#taskLists",
//     "etag": "\"LTg5NTE4MTk1Nw\"",
//     "items": [
//         {
//             "kind": "tasks#taskList",
//             "id": "MDI2OTgyMzY0NDIzNjIxOTI2MTk6MDow",
//             "etag": "\"MTk5NDgzNTIwMw\"",
//             "title": "My Tasks",
//             "updated": "2020-12-18T17:30:04.050Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/users/@me/lists/MDI2OTgyMzY0NDIzNjIxOTI2MTk6MDow"
//         },
//         {
//             "kind": "tasks#taskList",
//             "id": "RW9yQXktOXBOZ09rQS1rUw",
//             "etag": "\"LTIxMjU0NzE1NDY\"",
//             "title": "Work",
//             "updated": "2020-12-20T18:01:04.721Z",
//             "selfLink": "https://www.googleapis.com/tasks/v1/users/@me/lists/RW9yQXktOXBOZ09rQS1rUw"
//         }
//     ]
// }
