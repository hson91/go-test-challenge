package errors

// Error : struct
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// ErrorCode : int
type ErrorCode int

// ErrorWithMessage : his function return custom error
// params: error, ...MessageString
func ErrorWithMessage(err error, messages ...string) error {
	if err == nil {
		return nil
	}

	e := ToError(err)
	message := e.Message

	for _, msg := range messages {
		message += " " + msg
	}

	return &Error{
		Code:    e.Code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}

// ToError : convert error to Error type
func ToError(err error) *Error {
	if _, ok := err.(*Error); ok {
		return err.(*Error)
	}

	return &Error{
		Code:    ECUnknown,
		Message: err.Error(),
	}
}

func getError(errorCode ErrorCode, message string) error {
	return &Error{
		Code:    errorCode,
		Message: message,
	}
}
