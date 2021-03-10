package serializers

// UserReq : struct
type UserReq struct {
	ID    uint   `json:"id,string"`
	Alias string `json:"alias"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	Tag   string `json:"tag"`

	Active    *bool `json:"active,string"`
	Shared    *bool `json:"shared,string"`
	Suspended *bool `json:"suspended,string"`
	Verified  *bool `json:"verified,string"`
	IsReload  bool  `json:"is_reload"`

	OrganizationID uint `json:"organization_id,string"`
}
