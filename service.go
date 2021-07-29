package setdata_users

import (
	"github.com/google/uuid"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type UserService interface {
	CreateUser(cmd *CreateUserCommand) (*User, error)
	UpdateUser(cmd *UpdateUserCommand) (*User, error)
	DeleteUser(cmd *DeleteUserCommand) error
	GetUser(cmd *GetUserCommand) (*User, error)
	ListUser(cmd *ListUserCommand) ([]User, error)
}

type userService struct {
	userStore UsersStore
}

func NewUserService(uStore UsersStore) UserService {
	return &userService{userStore: uStore}
}

func (u *userService) CreateUser(cmd *CreateUserCommand) (*User, error) {
	user := &User{
		Id:        uuid.New().String(),
		Username:  cmd.Username,
		Password:  cmd.Password,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Email:     cmd.Email,
		LoginType: setdata_common.Email,
	}
	return u.userStore.Create(user)
}

func (u *userService) UpdateUser(cmd *UpdateUserCommand) (*User, error) {
	return nil, nil
}

func (u *userService) DeleteUser(cmd *DeleteUserCommand) error {
	return nil
}

func (u *userService) GetUser(cmd *GetUserCommand) (*User, error) {
	return nil, nil
}

func (u *userService) ListUser(cmd *ListUserCommand) ([]User, error) {
	return nil, nil
}
