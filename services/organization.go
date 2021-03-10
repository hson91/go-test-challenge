package services

import (
	"github.com/go-test-challenge/daos"
	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

// OrganizationSrv : struct
type OrganizationSrv struct {
	od *daos.OrganizationDAO
	ud *daos.UserDAO
	td *daos.TicketDAO
}

// NewOrganizationSrv : new instance OrganizationSrv
func NewOrganizationSrv(od *daos.OrganizationDAO, ud *daos.UserDAO, td *daos.TicketDAO) *OrganizationSrv {
	return &OrganizationSrv{
		od: od,
		ud: ud,
		td: td,
	}
}

// GetAllOrganization : get all organization
func (osrv *OrganizationSrv) GetAllOrganization(filters *serializers.OrganizationReq) ([]*models.Organization, error) {
	organizations, err := osrv.od.GetAllOrganization(filters)
	if err != nil {
		return nil, err
	}

	results, err := osrv.search(organizations, filters), nil
	if err != nil {
		return nil, err
	}

	if err := osrv.includeUsers(results); err != nil {
		return nil, err
	}

	if err := osrv.includeTickets(results); err != nil {
		return nil, err
	}

	return results, nil
}
