package setdata_users

import (
	"fmt"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"testing"
)

func TestUsersStore_Create(t *testing.T) {
	config := PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "setdatauser",
		Password: "123456789",
		Database: "setdata",
		Params:   "sslmode=disable",
	}
	store, err := NewPostgresUsersStore(config)
	if err != nil {
		panic(err)
	}
	newUser, err := store.Create(&User{
		Id:        "123",
		Username:  "456",
		Email:     "678",
		LoginType: setdata_common.Email,
		FirstName: "123",
		LastName:  "423",
	})
	fmt.Println(newUser)
}
