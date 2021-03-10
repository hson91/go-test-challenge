package services

import (
	"github.com/go-test-challenge/daos"
	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

// UserSrv struct {}
type UserSrv struct {
	od *daos.OrganizationDAO
	ud *daos.UserDAO
	td *daos.TicketDAO
}

// NewUserSrv : new instance UserSrv
func NewUserSrv(od *daos.OrganizationDAO, ud *daos.UserDAO, td *daos.TicketDAO) *UserSrv {
	return &UserSrv{
		od: od,
		ud: ud,
		td: td,
	}
}

// GetAllUser : get all users
func (usrv *UserSrv) GetAllUser(filters *serializers.UserReq) ([]*models.User, error) {
	tickets, err := usrv.ud.GetAllUsers(filters)
	if err != nil {
		return nil, err
	}

	results := usrv.search(tickets, filters)

	if err := usrv.includeTicketsAndOrganization(results); err != nil {
		return nil, err
	}

	return results, nil
}
