package models

/*
User
Interface to represent basic user authentication objects.
*/
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
RegisterResponse
Interface to represent user registration response structure.
*/
type RegisterResponse struct {
	Message string `json:"message"`
}

/*
LoginResponse
Interface to represent user login response structure.
*/
type LoginResponse struct {
	Token string `json:"token"`
}
