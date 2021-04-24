package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/repository"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/usecase/googleUser"
	// "github.com/jaredkgrove/TaskShare/TaskSyncProcessor/usecase/task"
	"github.com/jaredkgrove/TaskShare/TaskSyncProcessor/usecase/taskList"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	googleTasks "google.golang.org/api/tasks/v1"
)

func main() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "taskshare")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	googleUserRepo := repository.NewGoogleUserFirestore(client)
	googleUserService := googleUser.NewService(googleUserRepo)
	googleUsers, err := googleUserService.GetGoogleUsers(ctx)

	taskRepo := repository.NewTaskFirestore(client)
	// taskService := task.NewService(taskRepo)

	taskListRepo := repository.NewTaskListFirestore(client)
	taskListService := taskList.NewService(taskListRepo, taskRepo)


	if err != nil {
		fmt.Printf("Error getting googleUsers: %v", err)
	}
	fmt.Println(googleUsers)

	a := `{
		"web":{
			"client_id":"203481644879-kphv934udt5fnj31emn1fd1ffp81fr5r.apps.googleusercontent.com",
			"project_id":"taskshare",
			"auth_uri":"https://accounts.google.com/o/oauth2/auth",
			"token_uri":"https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs",
			"client_secret":"Bf4COmUSKhArN3vPlztgcca0",
			"redirect_uris":["https://taskshare-35739.firebaseapp.com/__/auth/handler"],
			"javascript_origins":["http://localhost","http://localhost:5000","https://taskshare-35739.firebaseapp.com"]}
		}`
	conf, err := google.ConfigFromJSON([]byte(a), "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/tasks")
	if err != nil {
		fmt.Printf("Error creating Oauth Config: %v", err)
	}

	for _, googleUser := range *googleUsers {
		token := new(oauth2.Token)
		token.AccessToken = googleUser.Token.AccessToken
		token.RefreshToken = googleUser.Token.RefreshToken
		token.Expiry = googleUser.Token.Expiry
		token.TokenType = googleUser.Token.TokenType

		tokenSource := conf.TokenSource(ctx, token)

		newToken, err := tokenSource.Token()
		if err != nil {
			log.Fatalln(err)
		}

		if newToken.AccessToken != token.AccessToken {
			googleUser.Token = newToken
			err = googleUserService.Update(ctx, &googleUser)
			if err != nil {
				log.Println("Error Saving googleUser: ", err)
			}
			log.Println("Old token:", googleUser.Token.AccessToken)
			log.Println("Saved new token:", newToken.AccessToken)
		} else {
			log.Println("Same Old Token: ", newToken.AccessToken)
		}
		tasksAPIService, err := googleTasks.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, token)))
		if err != nil {
			log.Println("Error Getting Task Lists: ", err)
		}
		taskListsAPIResponse, err := tasksAPIService.Tasklists.List().Do()

		if err != nil {
			log.Println("Error Getting Task Lists: ", err)
		}

		for _, googleTaskList := range taskListsAPIResponse.Items {

			// tasksResponse, err := tasksAPIService.Tasks.List(taskList).Do()
			if err != nil {
				log.Println("Error Getting Tasks: ", err)
			}
			err := taskListService.SaveFromGoogleTaskList(ctx, googleTaskList, googleUser.Ref.Parent.Parent.ID)
			if(err != nil){
				log.Println(err)
			}
			// fmt.Println(tasksResponse)
		}
		// client := oauth2.NewClient(oauth2.NoContext, tokenSource)
		// resp, err := client.Get(...)
	}
}
