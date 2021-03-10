package serializers

// TicketReq : struct
type TicketReq struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Subject string `json:"subject"`
	Type    string `json:"type"`
	Tag     string `json:"tag"`
	Via     string `json:"via"`

	SubmitterID    uint `json:"submitter_id,string"`
	AssigneeID     uint `json:"assignee_id,string"`
	OrganizationID uint `json:"organization_id,string"`

	IsReload bool `json:"is_reload"`
}
