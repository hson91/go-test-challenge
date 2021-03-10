package models

// UserRole : string
type UserRole string

// define role for user
const (
	UserRoleAgent   UserRole = "agent"
	UserRoleAdmin   UserRole = "admin"
	UserRoleEndUser UserRole = "end-user"
)

// UserRoleValue : map[string]UserRole
var UserRoleValue = map[string]UserRole{
	"agent":    UserRoleAgent,
	"admin":    UserRoleAdmin,
	"end-user": UserRoleEndUser,
}

// User : struct
type User struct {
	Base `json:",inline"`

	Alias       string `json:"alias"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Locale      string `json:"locale"`
	LastLoginAt string `json:"last_login_at"`
	Phone       string `json:"phone"`
	Signature   string `json:"signature"`
	URL         string `json:"url"`
	Timezone    string `json:"timezone"`

	Tags []string `json:"tags"`

	OrganizationID uint          `json:"organization_id"`
	Organization   *Organization `json:"organization"`

	Role UserRole `json:"role"`

	Active          bool      `json:"active"`
	Shared          bool      `json:"shared"`
	Suspended       bool      `json:"suspended"`
	Verified        bool      `json:"verified"`
	TicketsAssignee []*Ticket `json:"-"`
	TicketsSubmited []*Ticket `json:"-"`
}
