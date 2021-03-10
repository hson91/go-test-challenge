package errors

// Define error code for user
const (
	ECUserInfoInvalid ErrorCode = iota + IOTAECUserBegin
)

// define error string for user
var (
	UserIinfoInvalid = getError(ECUserInfoInvalid, "User info invalid.")
)
