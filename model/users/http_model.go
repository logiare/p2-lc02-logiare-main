package users

// Register is the request body for POST /users/register
type Register struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

// Login is the request body for POST /users/login
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ResponseRegister is the success response for register
type ResponseRegister struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

// ResponseLogin is the success response for login
type ResponseLogin struct {
	Token string `json:"token"`
}
