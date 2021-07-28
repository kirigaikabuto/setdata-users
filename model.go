package setdata_users

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	LoginType string `json:"login_type"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

