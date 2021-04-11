package entity

type GoogleUser struct {
	UserId       string //Firestore Id of user to which this subcollection belongs
	DisplayName  string `firestore:"displayName"`
	Email        string `firestore:"email"`
	Token        string `firestore:"tasksAccessToken"`
	RefreshToken string `firestore:"tasksRefreshToken"`
}
