package services

import (
	"fmt"
	"strings"

	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

func (usrv *UserSrv) itemMatchedFilters(user *models.User, filters *serializers.UserReq) bool {
	if filters.ID > 0 && filters.ID == user.ID {
		return true
	}

	if filters.Alias != "" && filters.Alias == user.Alias {
		return true
	}

	if filters.Email != "" && filters.Email == user.Email {
		return true
	}

	if filters.Name != "" && filters.Name == user.Name {
		return true
	}

	if filters.Phone != "" && filters.Phone == user.Phone {
		return true
	}

	if filters.OrganizationID > 0 && filters.OrganizationID == user.OrganizationID {
		return true
	}

	if role, ok := models.UserRoleValue[filters.Role]; ok && role == user.Role {
		return true
	}

	if filters.Active != nil && user.Active == *filters.Active {
		return true
	}

	if filters.Shared != nil && user.Shared == *filters.Shared {
		return true
	}

	if filters.Suspended != nil && user.Suspended == *filters.Suspended {
		return true
	}

	if filters.Verified != nil && user.Verified == *filters.Verified {
		return true
	}

	if filters.Tag != "" {
		for _, t := range user.Tags {
			if filters.Tag == t {
				return true
			}
		}
	}

	return false
}

// search : search users by filters
func (usrv *UserSrv) search(users []*models.User, filters *serializers.UserReq) []*models.User {
	if filters == nil {
		return users
	}
	var results []*models.User
	for _, item := range users {
		if usrv.itemMatchedFilters(item, filters) {
			results = append(results, item)
			if filters.ID > 0 {
				return results
			}
		}
	}

	return results
}

func (usrv *UserSrv) includeTicketsAndOrganization(users []*models.User) error {
	if users == nil {
		return nil
	}

	organizations, err := usrv.getOrganizationToDict()
	if err != nil {
		return err
	}

	assignees, submitters, err := usrv.getTicketsToDictGroupByUser()
	if err != nil {
		return err
	}

	for _, item := range users {
		if o, ok := organizations[item.OrganizationID]; ok {
			item.Organization = o
		}

		if tickets, ok := assignees[item.ID]; ok {
			item.TicketsAssignee = tickets
		}

		if tickets, ok := submitters[item.ID]; ok {
			item.TicketsSubmited = tickets
		}
	}
	return nil
}

func (usrv *UserSrv) print(u *models.User) string {
	result := ""
	result += fmt.Sprintf(
		"\n\n**--------------------------------------------------** \n"+
			"** ID: %d \n"+
			"ExternalID \t: %s \n"+
			"Name \t\t: %s \n"+
			"Alias \t\t: %s \n"+
			"Phone \t\t: %s \n"+
			"Role \t\t: %s \n"+
			"URL \t\t: %s \n"+
			"Signature \t: %s \n"+
			"Active \t: %t \n"+
			"Shared \t: %t \n"+
			"Suspended \t: %t \n"+
			"Verified \t: %t \n"+
			"Tags \t\t: %s \n"+
			"CreatedAt \t: %s"+
			"\n",
		u.ID,
		u.ExternalID,
		u.Name,
		u.Alias,
		u.Phone,
		u.Role,
		u.URL,
		u.Signature,
		u.Active,
		u.Shared,
		u.Suspended,
		u.Verified,
		strings.Join(u.Tags, ", "),
		u.CreatedAt,
	)

	if u.Organization != nil {
		result += fmt.Sprintf("+ Organization: %s \t ID: %d \n", u.Organization.Name, u.Organization.ID)
	}

	if u.TicketsAssignee != nil {
		result += "\n---- Ticket Assignee ----  \n"
		for i, u := range u.TicketsAssignee {
			result += fmt.Sprintf("Ticket %d: Subject: %s \n", i, u.Subject)
		}
	}

	if u.TicketsSubmited != nil {
		result += "\n---- Ticket Submited ----  \n"
		for i, u := range u.TicketsSubmited {
			result += fmt.Sprintf("Ticket %d: Subject: %s \n", i, u.Subject)
		}
	}

	return result
}

// PrintData :
func (usrv *UserSrv) PrintData(data []*models.User) string {
	results := ""

	for _, u := range data {
		results += usrv.print(u)
	}
	return results
}

// convert get all tickets and convert to map[userID][]*models.Ticket
// return TicketGroupSubmiter, TicketGroupAssignees, error
func (usrv *UserSrv) getTicketsToDictGroupByUser() (map[uint][]*models.Ticket, map[uint][]*models.Ticket, error) {
	tickets, err := usrv.td.GetAllTickets(nil)
	if err != nil {
		return nil, nil, err
	}

	var (
		assignees  = map[uint][]*models.Ticket{}
		submitters = map[uint][]*models.Ticket{}
	)
	for _, u := range tickets {
		if _, ok := assignees[u.AssigneeID]; !ok {
			assignees[u.AssigneeID] = []*models.Ticket{}
		}

		if _, ok := submitters[u.SubmitterID]; !ok {
			submitters[u.SubmitterID] = []*models.Ticket{}
		}

		assignees[u.AssigneeID] = append(assignees[u.AssigneeID], u)
		submitters[u.SubmitterID] = append(submitters[u.SubmitterID], u)

	}
	return assignees, submitters, nil
}

// get all organizations and convert to map[OrganizationID]*models.Organization
func (usrv *UserSrv) getOrganizationToDict() (map[uint]*models.Organization, error) {
	organizations, err := usrv.od.GetAllOrganization(nil)
	if err != nil {
		return nil, err
	}

	var result = map[uint]*models.Organization{}
	for _, u := range organizations {
		result[u.ID] = u
	}
	return result, nil
}
