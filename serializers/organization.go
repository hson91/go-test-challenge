package serializers

// OrganizationReq : struct
type OrganizationReq struct {
	ID            uint   `json:"id,string"`
	IsReload      bool   `json:"is_reload"`
	SharedTickets *bool  `json:"shared_tickets,string"`
	Domain        string `json:"domain"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
}
