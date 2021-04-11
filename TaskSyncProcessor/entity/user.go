package entity

type User struct {
	FirestoreID   string
	DisplayName   string `firestore:"displayName"`
	EmailAddress  string `firestore:"emailAddress"`
	EmailVerified bool   `firestore:"emailVerified"`
	GoogleUserUID string `firestore:"googleUserUID"`
	// GoogleUser    GoogleUser
}
