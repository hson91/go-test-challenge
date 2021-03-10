package errors

// Define error code for command
const (
	ECCommandEmpty ErrorCode = iota + IOTAECCommandBegin
	ECCommandSyntaxError
	ECCommandActionNotFound
	ECCommandObjectNotFound
)

// define error string for tiket
var (
	CommandEmpty          = getError(ECCommandEmpty, "Syntax is empty.")
	CommandSyntaxError    = getError(ECTicketInfoInvalid, "Syntax error.")
	CommandActionNotFound = getError(ECCommandActionNotFound, "Action not found.")
	CommandObjectNotFound = getError(ECCommandObjectNotFound, "Object not found.")
)
