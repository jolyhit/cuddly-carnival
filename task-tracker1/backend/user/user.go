package user

/*
this struct gets data from inputs on authorization page and stores it in different params of struct called User
*/
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
