package errors

// Define error code for organization
const (
	ECOrganizationInfoInvalid ErrorCode = iota + IOTAECOrganizationBegin
	ECOrganizationNotFound
)

// define error string for organization
var (
	OrganizationInfoInvalid = getError(ECOrganizationInfoInvalid, "Organization info invalid.")
	OrganizationNotFound    = getError(ECOrganizationNotFound, "Organization not found.")
)
