package models

// Organization : struct
type Organization struct {
	Base    `json:",inline"`
	Details string `json:"details"`
	Name    string `json:"name"`
	URL     string `json:"url"`

	DomainNames []string `json:"domain_names"`
	Tags        []string `json:"tags"`

	SharedTickets bool `json:"shared_tickets"`

	Tickets []*Ticket `json:"-"`
	Users   []*User   `json:"-"`
}
