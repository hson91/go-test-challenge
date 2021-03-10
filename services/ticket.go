package services

import (
	"github.com/go-test-challenge/daos"
	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

// TicketSrv : struct
type TicketSrv struct {
	od *daos.OrganizationDAO
	ud *daos.UserDAO
	td *daos.TicketDAO
}

// NewTicketSrv : new instance TicketSrv
func NewTicketSrv(od *daos.OrganizationDAO, ud *daos.UserDAO, td *daos.TicketDAO) *TicketSrv {
	return &TicketSrv{
		od: od,
		ud: ud,
		td: td,
	}
}

// GetAllTicket : get all tickers
func (tsrv *TicketSrv) GetAllTicket(filters *serializers.TicketReq) ([]*models.Ticket, error) {
	tickets, err := tsrv.td.GetAllTickets(filters)
	if err != nil {
		return nil, err
	}

	results, err := tsrv.search(tickets, filters), nil
	if err != nil {
		return nil, err
	}

	if err := tsrv.includeUsersAndOrganization(results); err != nil {
		return nil, err
	}

	return results, nil
}
