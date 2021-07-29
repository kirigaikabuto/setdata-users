package setdata_users

import (
	cm "github.com/kirigaikabuto/setdata-common"
)

type User struct {
	Id        string       `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	LoginType cm.LoginType `json:"login_type"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
}

type UserUpdate struct {
	Id        string        `json:"id"`
	Username  *string       `json:"username"`
	Email     *string       `json:"email"`
	LoginType *cm.LoginType `json:"login_type"`
	FirstName *string       `json:"first_name"`
	LastName  *string       `json:"last_name"`
}
