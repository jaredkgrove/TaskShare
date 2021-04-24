package entity

import (
	"cloud.google.com/go/firestore"
	"golang.org/x/oauth2"
)

type GoogleUser struct {
	Ref         *firestore.DocumentRef //I don't like this
	DisplayName string                 `firestore:"displayName"`
	Email       string                 `firestore:"email"`
	Token       *oauth2.Token          `firestore:"token"`
}

// type googleToken struct {
// 	AccessToken  string    `firestore:"AccessToken"`
// 	RefreshToken string    `firestore:"RefreshToken"`
// 	Expiry       time.Time `firestore:"Expiry"`
// 	TokenType    string    `firestore:"TokenType"`
// }
