package models

type User struct {
    ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role"`
    Username string `json:"username"`
    Password string `json:"password"` // hashed
}
