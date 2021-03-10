package utilities

import (
	"testing"

	"github.com/go-test-challenge/models"
	"github.com/stretchr/testify/assert"
)

// Test utilities.Switch2Struct
func TestSwitch2Struct(t *testing.T) {
	assert := assert.New(t)
	var struct1 = map[string]interface{}{
		"name": "Pham Hoang Son",
		"_id":  123,
	}

	var struct2 *models.User

	err := Switch2Struct(&struct1, &struct2)
	assert.Nil(err, err)

	assert.Equal(struct2.Name, struct1["name"])
	assert.Equal(int(struct2.ID), int(123))

	var struct3 *models.Ticket
	err = Switch2Struct(&struct1, &struct3)
	assert.NotNil(err)
}
