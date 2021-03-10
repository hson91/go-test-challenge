package services

import (
	"testing"

	"github.com/go-test-challenge/serializers"
	"github.com/stretchr/testify/assert"
)

func TestGetAllOrganizations(t *testing.T) {
	assert := assert.New(t)
	isTrue := true
	req := &serializers.OrganizationReq{
		SharedTickets: &isTrue,
	}

	organizationSrv := NewOrganizationSrv(testOrganizationDAO, testUserDAO, testTickerDAO)

	organizations, err := organizationSrv.GetAllOrganization(req)
	assert.Nil(err)
	assert.NotNil(organizations)

	for _, o := range organizations {
		assert.True(o.SharedTickets)
	}
}
