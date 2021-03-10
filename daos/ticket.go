package daos

import (
	"encoding/json"

	"github.com/go-test-challenge/config"
	"github.com/go-test-challenge/errors"
	"github.com/go-test-challenge/libs"
	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

// TicketDAO : struct
type TicketDAO struct {
	Data []*models.Ticket
}

// NewTicketDAO : new instance UserDAO
func NewTicketDAO() *TicketDAO {
	return &TicketDAO{}
}

// GetAllTickets : get all organization
func (td *TicketDAO) GetAllTickets(filters *serializers.TicketReq) ([]*models.Ticket, error) {
	if td.Data != nil && filters != nil && !filters.IsReload {
		return td.Data, nil
	}

	byteValues, err := libs.ReadFile(config.PathFileTicket)
	if err != nil {
		return nil, err

	}

	if err := json.Unmarshal(byteValues, &td.Data); err != nil {
		return nil, errors.ErrorWithMessage(errors.UnmarshalError, err.Error())
	}

	return td.Data, nil
}
