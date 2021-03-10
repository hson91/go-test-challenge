package daos

import (
	"encoding/json"

	"github.com/go-test-challenge/errors"
	"github.com/go-test-challenge/libs"
	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

// OrganizationDAO : struct
type OrganizationDAO struct {
	Data     []*models.Organization
	filePath string
}

// NewOrganizationDAO : new instance OrganizationDAO
func NewOrganizationDAO(filePath string) *OrganizationDAO {
	return &OrganizationDAO{
		filePath: filePath,
	}
}

// GetAllOrganization : get all organization
func (od *OrganizationDAO) GetAllOrganization(filters *serializers.OrganizationReq) ([]*models.Organization, error) {
	if od.Data != nil && filters != nil && !filters.IsReload {
		return od.Data, nil
	}

	byteValues, err := libs.ReadFile(od.filePath)
	if err != nil {
		return nil, err

	}

	if err := json.Unmarshal(byteValues, &od.Data); err != nil {
		return nil, errors.ErrorWithMessage(errors.UnmarshalError, err.Error())
	}

	return od.Data, nil
}
