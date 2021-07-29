package setdata_users

type CreateUserCommand struct {
	SuperUserId string `json:"super_user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

func (cmd *CreateUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserService).CreateUser(cmd)
}

type UpdateUserCommand struct {
	SuperUserId string `json:"super_user_id"`
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

func (cmd *UpdateUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(UserService).UpdateUser(cmd)
}

type DeleteUserCommand struct {
	SuperUserId string `json:"super_user_id"`
	Id          string `json:"id"`
}

type GetUserCommand struct {
	SuperUserId string `json:"super_user_id"`
	Id          string `json:"id"`
}

type ListUserCommand struct {
	SuperUserId string `json:"super_user_id"`
}
