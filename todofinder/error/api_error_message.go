package error

//API Error structure
type Error struct {
	ErrorCode string
	Error     error
}

//Return a verbose message corresponding to the api error
func (error *Error) ToString() string {
	return ErrorMessages[error.ErrorCode].ErrorMessage
}

//API Error Codes
const (
	NOT_FOUND           = "not_found"
	METHOD_NOT_ALLOWED  = "method_not_allowed"
	BAD_BODY            = "bad_body"
	BAD_PARAMETER       = "bad_parameter"
	SERVER_ERROR        = "internal_error"
	PACKAGE_NOT_FOUND   = "package_not_found"
	NO_SOURCE           = "no_source"
	SOURCE_NOT_READABLE = "source_not_readable"
)

//API HTTP Error object
type errorHttpMessage struct {
	ErrorMessage string
	HttpStatus   int
}

//Map between API error code and HTTP Error object
var ErrorMessages = map[string]*errorHttpMessage{
	NOT_FOUND:           {"Resource not found", 404},
	METHOD_NOT_ALLOWED:  {"Unexpected method", 405},
	BAD_BODY:            {"Incorrect request body", 400},
	BAD_PARAMETER:       {"Wrong parameters", 400},
	SERVER_ERROR:        {"Internal server error", 500},
	PACKAGE_NOT_FOUND:   {"Cannot find Package", 404},
	NO_SOURCE:           {"No Buildable go source file", 404},
	SOURCE_NOT_READABLE: {"Cannot read Source", 500},
}
