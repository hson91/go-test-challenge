package services

import (
	"fmt"
	"strings"

	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

func (osrv *OrganizationSrv) itemMatchedFilters(organization *models.Organization, filters *serializers.OrganizationReq) bool {
	if filters.ID > 0 && filters.ID == organization.ID {
		return true
	}

	if filters.Name != "" && filters.Name == organization.Name {
		return true
	}

	if filters.Domain != "" {
		for _, d := range organization.DomainNames {
			if filters.Domain == d {
				return true
			}
		}
	}

	if filters.Tag != "" {
		for _, t := range organization.Tags {
			if filters.Tag == t {
				return true
			}
		}
	}

	if filters.SharedTickets != nil && organization.SharedTickets == *filters.SharedTickets {
		return true
	}

	return false
}

// search : get all organization
func (osrv *OrganizationSrv) search(organizations []*models.Organization, filters *serializers.OrganizationReq) []*models.Organization {
	if filters == nil {
		return organizations
	}
	var results []*models.Organization
	for _, item := range organizations {
		if osrv.itemMatchedFilters(item, filters) {
			results = append(results, item)
			if filters.ID > 0 {
				return results
			}
		}
	}

	return results
}

func (osrv *OrganizationSrv) includeUsers(organizations []*models.Organization) error {
	if organizations == nil {
		return nil
	}

	users, err := osrv.getUsersToDict()
	if err != nil {
		return err
	}

	for _, item := range organizations {
		if _, ok := users[item.ID]; ok {
			item.Users = users[item.ID]
		}
	}
	return nil
}

func (osrv *OrganizationSrv) includeTickets(organizations []*models.Organization) error {
	if organizations == nil {
		return nil
	}

	tickets, err := osrv.getTicketsToDict()
	if err != nil {
		return err
	}

	for _, item := range organizations {
		if _, ok := tickets[item.ID]; ok {
			item.Tickets = tickets[item.ID]
		}
	}
	return nil
}

func (osrv *OrganizationSrv) getUsersToDict() (map[uint][]*models.User, error) {
	users, err := osrv.ud.GetAllUsers(nil)
	if err != nil {
		return nil, err
	}

	var result = map[uint][]*models.User{}
	for _, u := range users {
		if _, ok := result[u.OrganizationID]; !ok {
			result[u.OrganizationID] = []*models.User{}
		}
		result[u.OrganizationID] = append(result[u.OrganizationID], u)
	}
	return result, nil
}

func (osrv *OrganizationSrv) getTicketsToDict() (map[uint][]*models.Ticket, error) {
	tickets, err := osrv.td.GetAllTickets(nil)
	if err != nil {
		return nil, err
	}

	var result = map[uint][]*models.Ticket{}
	for _, t := range tickets {
		if _, ok := result[t.OrganizationID]; !ok {
			result[t.OrganizationID] = []*models.Ticket{}
		}
		result[t.OrganizationID] = append(result[t.OrganizationID], t)
	}
	return result, nil
}

func (osrv *OrganizationSrv) print(o *models.Organization) string {
	result := ""
	result += fmt.Sprintf(
		"\n\n**--------------------------------------------------** \n"+
			"** ID: %d \n"+
			"ExternalID \t: %s \n"+
			"Name \t\t: %s \n"+
			"URL \t\t: %s \n"+
			"SharedTickets \t: %t \n"+
			"DomainName \t: %s \n"+
			"Tags \t\t: %s \n"+
			"CreatedAt \t: %s"+
			"\n",
		o.ID,
		o.ExternalID,
		o.Name,
		o.URL,
		o.SharedTickets,
		strings.Join(o.DomainNames, ", "),
		strings.Join(o.Tags, ", "),
		o.CreatedAt,
	)

	if o.Users != nil {
		result += "\n---- Users List ----  \n"
		for i, u := range o.Users {
			result += fmt.Sprintf("User %d: Name: %s \t ID: %d \n", i, u.Name, u.ID)
		}
	}

	if o.Tickets != nil {
		result += "\n---- Ticket List ----  \n"
		for i, t := range o.Tickets {
			result += fmt.Sprintf("Ticket %d: Subject: %s \n", i, t.Subject)
		}
	}

	return result
}

// PrintData :
func (osrv *OrganizationSrv) PrintData(data []*models.Organization) string {
	results := ""

	for _, o := range data {
		results += osrv.print(o)
	}
	return results
}
