package error

import (
	"encoding/json"
)

//Error HTTP response for server
type ErrorHttpResponse struct {
	ErrorCode  string `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"status"`
}

//Create new instance of error
func NewError(err error, errorCode string) *Error {
	return &Error{errorCode, err}
}

//Convert an error object to JSON string
func (errorResponse *ErrorHttpResponse) ToJson() string {
	errorJson, err := json.Marshal(errorResponse)
	if err != nil {
		//Log.Log.WithFields(logrus.Fields{"err": err, "func": "toJson"}).Fatal("Couldn't encode JSON")
		return ""
	}

	return string(errorJson)
}