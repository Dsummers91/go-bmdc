package user

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserProfile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Partner  string `json:"partner"`
	City     string `json:"city"`
	State    string `json:"state"`
}
