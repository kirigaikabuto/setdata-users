package setdata_users

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrCreateUserUnknown = com.NewMiddleError(errors.New("could not create user: unknown error"), 500, 1)
	ErrUserNotFound      = com.NewMiddleError(errors.New("user not found"), 404, 2)
	ErrNothingToUpdate   = com.NewMiddleError(errors.New("nothing to update"), 400, 3)
)
