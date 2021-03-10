package daos

import (
	"encoding/json"

	"github.com/go-test-challenge/errors"
	"github.com/go-test-challenge/libs"
	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

// TicketDAO : struct
type TicketDAO struct {
	Data     []*models.Ticket
	filePath string
}

// NewTicketDAO : new instance UserDAO
func NewTicketDAO(filePath string) *TicketDAO {
	return &TicketDAO{
		filePath: filePath,
	}
}

// GetAllTickets : get all organization
func (td *TicketDAO) GetAllTickets(filters *serializers.TicketReq) ([]*models.Ticket, error) {
	if td.Data != nil && filters != nil && !filters.IsReload {
		return td.Data, nil
	}

	byteValues, err := libs.ReadFile(td.filePath)
	if err != nil {
		return nil, err

	}

	if err := json.Unmarshal(byteValues, &td.Data); err != nil {
		return nil, errors.ErrorWithMessage(errors.UnmarshalError, err.Error())
	}

	return td.Data, nil
}
