package userprofile

type userProfile struct {
	ID    string `json:"Id"`
	FName string `json:"fname"`
	LName string `json:"lname"`
	Email string `json:"email"`
}
