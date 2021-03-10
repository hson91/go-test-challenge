package errors

const (
	// IOTAECUserBegin : Error code for User will start from 1000 to 1999
	IOTAECUserBegin ErrorCode = 1000

	// IOTAECOrganizationBegin : Error code for Organization will start from 2000 to 2999
	IOTAECOrganizationBegin ErrorCode = 2000

	// IOTAECTicketBegin : Error code for Ticket will start from 3000 to 3999
	IOTAECTicketBegin ErrorCode = 3000

	// IOTAECCommandBegin : Error code for command will start from 3000
	IOTAECCommandBegin ErrorCode = 4000
)

// Define error code for system
const (
	ECUnknown          ErrorCode = -1
	ECOpenFileHasError ErrorCode = 1
	ECReadFileHasError ErrorCode = 2
	ERUnmarshalError   ErrorCode = 3
	ERMarshalError     ErrorCode = 4
)

// define error message
var (
	Unknow           = getError(ECUnknown, "Error Unknown on app")
	OpenFileHasError = getError(ECOpenFileHasError, "Has error when open file.")
	ReadFileHasError = getError(ECReadFileHasError, "Has error when read file.")
	UnmarshalError   = getError(ERUnmarshalError, "Unmarshal has error.")
	MarshalError     = getError(ERMarshalError, "Marshal has error.")
)
