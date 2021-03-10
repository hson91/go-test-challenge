package services

import (
	"fmt"
	"strings"

	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
)

func (tsrv *TicketSrv) itemMatchedFilters(ticket *models.Ticket, filters *serializers.TicketReq) bool {
	if filters.ID != "" && filters.ID == ticket.ID {
		return true
	}

	if filters.SubmitterID > 0 && filters.SubmitterID == ticket.SubmitterID {
		return true
	}

	if filters.AssigneeID > 0 && filters.AssigneeID == ticket.AssigneeID {
		return true
	}

	if filters.OrganizationID > 0 && filters.OrganizationID == ticket.OrganizationID {
		return true
	}

	if filters.Subject != "" && filters.Subject == ticket.Subject {
		return true
	}

	if status, ok := models.TicketStatusValue[filters.Status]; ok && status == ticket.Status {
		return true
	}

	if ticketType, ok := models.TicketTypeValue[filters.Type]; ok && ticketType == ticket.Type {
		return true
	}

	if ticketVia, ok := models.TicketViaValue[filters.Via]; ok && ticketVia == ticket.Via {
		return true
	}

	if filters.Tag != "" {
		for _, t := range ticket.Tags {
			if filters.Tag == t {
				return true
			}
		}
	}

	return false
}

// search : get all Ticket
func (tsrv *TicketSrv) search(tickets []*models.Ticket, filters *serializers.TicketReq) []*models.Ticket {
	if filters == nil {
		return tickets
	}

	var results []*models.Ticket
	for _, item := range tickets {
		if tsrv.itemMatchedFilters(item, filters) {
			results = append(results, item)
			if filters.ID != "" {
				return results
			}
		}
	}

	return results
}

func (tsrv *TicketSrv) includeUsersAndOrganization(tickets []*models.Ticket) error {
	if tickets == nil {
		return nil
	}

	users, err := tsrv.getUsersToDict()
	if err != nil {
		return err
	}

	organizations, err := tsrv.getOrganizationToDict()
	if err != nil {
		return err
	}

	for _, item := range tickets {
		if user, ok := users[item.SubmitterID]; ok {
			item.Submitter = user
		}

		if user, ok := users[item.AssigneeID]; ok {
			item.Assignee = user
		}

		if o, ok := organizations[item.OrganizationID]; ok {
			item.Organization = o
		}
	}
	return nil
}

func (tsrv *TicketSrv) getUsersToDict() (map[uint]*models.User, error) {
	users, err := tsrv.ud.GetAllUsers(nil)
	if err != nil {
		return nil, err
	}

	var result = map[uint]*models.User{}
	for _, u := range users {
		result[u.ID] = u
	}
	return result, nil
}

func (tsrv *TicketSrv) getOrganizationToDict() (map[uint]*models.Organization, error) {
	organizations, err := tsrv.od.GetAllOrganization(nil)
	if err != nil {
		return nil, err
	}

	var result = map[uint]*models.Organization{}
	for _, t := range organizations {
		result[t.ID] = t
	}
	return result, nil
}

func (tsrv *TicketSrv) print(t *models.Ticket) string {
	if t == nil {
		return ""
	}

	result := fmt.Sprintf(
		"\n\n**--------------------------------------------------** \n"+
			"** ID: %s \n"+
			"ExternalID \t: %s \n"+
			"Subject \t\t: %s \n"+
			"URL \t\t: %s \n"+
			"Description \t\t: %s \n"+
			"DueAt \t: %s \n"+
			"HasIncidents \t: %t \n"+
			"Priority \t: %s \n"+
			"Status \t: %s \n"+
			"Tags \t\t: %s \n"+
			"Type \t\t: %s \n"+
			"Via \t\t: %s \n"+
			"CreatedAt \t: %s"+
			"\n",
		t.ID,
		t.ExternalID,
		t.Subject,
		t.URL,
		t.Description,
		t.DueAt,
		t.HasIncidents,
		t.Priority,
		t.Status,
		strings.Join(t.Tags, ", "),
		t.Type,
		t.Via,
		t.CreatedAt,
	)

	if t.Assignee != nil {
		result += fmt.Sprintf("+ Assignee: %s \t ID: %d \n", t.Assignee.Name, t.Assignee.ID)
	}

	if t.Submitter != nil {
		result += fmt.Sprintf("+ Submitter: %s \t ID: %d \n", t.Submitter.Name, t.Submitter.ID)
	}

	if t.Organization != nil {
		result += fmt.Sprintf("+ Organization: %s \t ID: %d \n", t.Organization.Name, t.Organization.ID)
	}

	return result
}

// PrintData :
func (tsrv *TicketSrv) PrintData(data []*models.Ticket) string {
	if data == nil {
		return ""
	}

	results := ""

	for _, o := range data {
		results += tsrv.print(o)
	}
	return results
}
