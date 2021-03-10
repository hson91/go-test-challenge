package daos

import (
	"encoding/json"

	"github.com/go-test-challenge/config"
	"github.com/go-test-challenge/errors"
	"github.com/go-test-challenge/libs"
	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

// UserDAO : struct
type UserDAO struct {
	Data []*models.User
}

// NewUserDAO : new instance UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// GetAllUsers : get all organization
func (ud *UserDAO) GetAllUsers(filters *serializers.UserReq) ([]*models.User, error) {
	if ud.Data != nil && filters != nil && !filters.IsReload {
		return ud.Data, nil
	}

	byteValues, err := libs.ReadFile(config.PathFileUser)
	if err != nil {
		return nil, err

	}

	if err := json.Unmarshal(byteValues, &ud.Data); err != nil {
		return nil, errors.ErrorWithMessage(errors.UnmarshalError, err.Error())
	}

	return ud.Data, nil
}
