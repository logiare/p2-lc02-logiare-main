package users

type Register struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseRegister struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}
