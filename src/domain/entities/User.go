package entities


type User struct {
	ID int64 `json:"id"` 
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(id int64, username string, password string) *User {
	return &User{
		ID: id,
		Username: username,
		Password: password,
	}
}

