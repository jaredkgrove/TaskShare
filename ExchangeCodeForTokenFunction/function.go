package p

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/functions/metadata"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type FirestoreEvent struct {
	OldValue FirestoreValue `json:"oldValue"`
	Value    FirestoreValue `json:"value"`
}

type FirestoreValue struct {
	CreateTime time.Time  `json:"createTime"`
	Fields     GoogleUser `json:"fields"`
	Name       string     `json:"name"`
	UpdateTime time.Time  `json:"updateTime"`
}

type GoogleUser struct {
	ServerAuthCode struct {
		StringValue string `json:"stringValue"`
	} `json:"serverAuthCode"`
}

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	a := fmt.Sprintln(`{
		"web":{
			"client_id":%v,
			"project_id":"taskshare",
			"auth_uri":"https://accounts.google.com/o/oauth2/auth",
			"token_uri":"https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs",
			"client_secret":%v,
			"redirect_uris":["https://taskshare-35739.firebaseapp.com/__/auth/handler"],
			"javascript_origins":["http://localhost","http://localhost:5000","https://taskshare-35739.firebaseapp.com"]}
		}`, os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"))
	var err error
	googleOauthConfig, err = google.ConfigFromJSON([]byte(a), "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/tasks")
	if err != nil {
		log.Println("Error creating Oauth Config: %v", err)
	}
}

func ExchangeCodeForToken(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function triggered by change to: %v", meta.Resource)

	fullPath := strings.Split(e.Value.Name, "/documents/")[1]

	token, err := googleOauthConfig.Exchange(ctx, e.Value.Fields.ServerAuthCode.StringValue)
	if err != nil {
		log.Println(err)
		return err
	}

	client, err := firestore.NewClient(ctx, "taskshare")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = client.Doc(fullPath).Set(ctx, map[string]interface{}{
		"token": token,
	}, firestore.MergeAll)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
