package services

import (
	"testing"

	"github.com/go-test-challenge/serializers"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	assert := assert.New(t)
	isTrue := true

	var req = &serializers.UserReq{IsReload: false, Shared: &isTrue}

	userSrv := NewUserSrv(testOrganizationDAO, testUserDAO, testTickerDAO)

	users, err := userSrv.GetAllUser(req)
	assert.Nil(err)
	assert.NotNil(users)

	for _, u := range users {
		assert.True(u.Shared)
	}
}
