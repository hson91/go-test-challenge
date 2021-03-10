package command

import (
	"testing"

	"github.com/go-test-challenge/daos"
	"github.com/go-test-challenge/services"
	"github.com/stretchr/testify/assert"
)

var (
	organizationDAOTest = daos.NewOrganizationDAO("../data/organizations.json")
	userDAOTest         = daos.NewUserDAO("../data/users.json")
	tickerDAOTest       = daos.NewTicketDAO("../data/tickets.json")

	organizationSrvTest = services.NewOrganizationSrv(organizationDAOTest, userDAOTest, tickerDAOTest)
	userSrvTest         = services.NewUserSrv(organizationDAOTest, userDAOTest, tickerDAOTest)
	ticketSrvTest       = services.NewTicketSrv(organizationDAOTest, userDAOTest, tickerDAOTest)

	cliTest = NewCLI(userSrvTest, organizationSrvTest, ticketSrvTest)
)

func TestParseCommandFromString(t *testing.T) {
	assert := assert.New(t)
	s := "user find id=1"

	cmd := cliTest.ParseCommand(s)
	assert.NotNil(cmd)

	assert.Equal(cmd.Action, "find")
	assert.Equal(cmd.Object, "user")

	s = "abc def adfg"

	cmd = cliTest.ParseCommand(s)
	assert.NotNil(cmd)

	assert.Equal(cmd.Action, "def")
	assert.Equal(cmd.Object, "abc")

	s = "organization find id=1 name='pham hoang son' "
	cmd = cliTest.ParseCommand(s)
	assert.NotNil(cmd)

	assert.Equal(cmd.Action, "find")
	assert.Equal(cmd.Object, "organization")

	assert.True(len(cmd.Filters) == 2)
}
