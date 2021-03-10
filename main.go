package main

import (
	"github.com/go-test-challenge/command"
	"github.com/go-test-challenge/daos"
	"github.com/go-test-challenge/services"
)

func main() {
	var (
		organizationDAO = daos.NewOrganizationDAO()
		userDAO         = daos.NewUserDAO()
		tickerDAO       = daos.NewTicketDAO()

		organizationSrv = services.NewOrganizationSrv(organizationDAO, userDAO, tickerDAO)
		userSrv         = services.NewUserSrv(organizationDAO, userDAO, tickerDAO)
		ticketSrv       = services.NewTicketSrv(organizationDAO, userDAO, tickerDAO)
	)

	cli := command.NewCLI(userSrv, organizationSrv, ticketSrv)
	cli.Run()
}
