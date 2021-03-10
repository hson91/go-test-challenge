package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCommand(t *testing.T) {
	assert := assert.New(t)
	cmd := &Command{}

	err := cliTest.validateCommand(cmd)
	assert.NotNil(err)

	cmd.Action = "abc"
	err = cliTest.validateCommand(cmd)
	assert.NotNil(err)

	cmd.Action = ActionFind

	cmd.Object = "bcd"
	err = cliTest.validateCommand(cmd)
	assert.NotNil(err)

	cmd.Object = ObjectTicket
	err = cliTest.validateCommand(cmd)
	assert.Nil(err)

}
