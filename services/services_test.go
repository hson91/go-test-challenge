package services

import "github.com/go-test-challenge/daos"

var (
	testOrganizationDAO = daos.NewOrganizationDAO("../data/organizations.json")
	testUserDAO         = daos.NewUserDAO("../data/users.json")
	testTickerDAO       = daos.NewTicketDAO("../data/tickets.json")
)
