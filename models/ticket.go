package models

// TicketStatus : string
type TicketStatus string

// define ticket status
const (
	TicketStatusOpen    TicketStatus = "open"
	TicketStatusPending TicketStatus = "pending"
	TicketStatusHold    TicketStatus = "hold"
	TicketStatusSolved  TicketStatus = "solved"
	TicketStatusClosed  TicketStatus = "closed"
)

// TicketStatusValue : map[string]TicketStatus
var TicketStatusValue = map[string]TicketStatus{
	"open":    TicketStatusOpen,
	"pending": TicketStatusPending,
	"hold":    TicketStatusHold,
	"solved":  TicketStatusSolved,
	"closed":  TicketStatusClosed,
}

// TicketVia : string
type TicketVia string

// define ticket via
const (
	TicketViaWeb   TicketVia = "web"
	TicketViaChat  TicketVia = "chat"
	TicketViaVoice TicketVia = "voice"
)

// TicketViaValue : map[string]TicketVia
var TicketViaValue = map[string]TicketVia{
	"web":   TicketViaWeb,
	"chat":  TicketViaChat,
	"voice": TicketViaVoice,
}

// TicketType : string
type TicketType string

// define ticket type
const (
	TicketTypeIncident TicketType = "incident"
	TicketTypeProblem  TicketType = "problem"
	TicketTypeQuestion TicketType = "question"
	TicketTypeTask     TicketType = "task"
)

// TicketTypeValue : map[string]TicketType
var TicketTypeValue = map[string]TicketType{
	"incident": TicketTypeIncident,
	"problem":  TicketTypeProblem,
	"question": TicketTypeQuestion,
	"task":     TicketTypeTask,
}

// TicketPriority : string
type TicketPriority string

// define priority for ticket
const (
	TicketPriorityNormal TicketPriority = "normal"
	TicketPriorityLow    TicketPriority = "low"
	TicketPriorityHeight TicketPriority = "high"
	TicketPriorityUrgent TicketPriority = "urgent"
)

// TicketPriorityVallue : map[string]TicketPriority
var TicketPriorityVallue = map[string]TicketPriority{
	"normal": TicketPriorityNormal,
	"low":    TicketPriorityLow,
	"high":   TicketPriorityHeight,
	"urgent": TicketPriorityUrgent,
}

// Ticket : struct
type Ticket struct {
	ID         string `json:"_id"`
	ExternalID string `json:"external_id"`
	CreatedAt  string `json:"created_at"`

	AssigneeID     uint `json:"assignee_id"`
	OrganizationID uint `json:"organization_id"`
	SubmitterID    uint `json:"submitter_id"`

	Description string   `json:"description"`
	DueAt       string   `json:"due_at"`
	URL         string   `json:"url"`
	Subject     string   `json:"subject"`
	Tags        []string `json:"tags"`

	HasIncidents bool `json:"has_incidents"`

	Priority TicketPriority `json:"priority"`
	Status   TicketStatus   `json:"status"`
	Type     TicketType     `json:"type"`
	Via      TicketVia      `json:"via"`

	Assignee     *User         `json:"assignee"`
	Submitter    *User         `json:"submitter"`
	Organization *Organization `json:"organization"`
}
