package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	assert := assert.New(t)

	b, err := ReadFile("abc/def.json")

	assert.Nil(b)
	assert.NotNil(err)

	b, err = ReadFile("../data/users.json")

	assert.NotNil(b)
	assert.Nil(err)

}
