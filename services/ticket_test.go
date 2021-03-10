package services

import (
	"testing"

	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTickets(t *testing.T) {
	assert := assert.New(t)
	req := &serializers.TicketReq{
		Status: string(models.TicketStatusOpen),
	}

	ticketSrv := NewTicketSrv(testOrganizationDAO, testUserDAO, testTickerDAO)

	tickets, err := ticketSrv.GetAllTicket(req)
	assert.Nil(err)
	assert.NotNil(tickets)
	if tickets != nil {
		for _, t := range tickets {
			assert.Equal(models.TicketStatusOpen, t.Status)
		}
	}

}
