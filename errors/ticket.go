package errors

// Define error code for ticket
const (
	ECTicketInfoInvalid ErrorCode = iota + IOTAECTicketBegin
)

// define error string for tiket
var (
	TicketInfoInvalid = getError(ECTicketInfoInvalid, "Tiket info invalid.")
)
