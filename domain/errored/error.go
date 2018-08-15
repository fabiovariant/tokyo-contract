package errored

import "encoding/json"

//JSONError Struct to return errors
type JSONError struct {
	ErrorID string `json:"error_id"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
}

func (e JSONError) String() string {
	b, _ := json.MarshalIndent(e, "", "    ")
	return string(b)
}

func (e JSONError) Error() string {
	return e.String()
}

//GetErrorMessage returns the error message by the code
func GetErrorMessage(code string) string {
	if code == "" {
		code = "E0000"
	}
	if _, ok := messageErros[code]; ok {
		return messageErros[code]
	}
	// If error code do not exists, return a generic message.
	return "Internal error"
}

// Knowed error codes
var messageErros = map[string]string{
	"REQ0001": "Incorrect body",
	"E0000":   "Internal server error",
	"E0001":   "Error sending e-mail",
	//general
	"A0000": "Database unknow error",
	"A0001": "E-mail not found",
	"A0002": "E-mail already registred",
	//house
	"HOU0001": "House trade invalid",
	"HOU0002": "House company name invalid",
	"HOU0003": "House document id invalid",
	"HOU0004": "House location invalid",
	"HOU0005": "House address invalid",
}

//GetErrorType return a error type builded
func GetErrorType(code string, stack error) (jErr JSONError) {
	jErr.ErrorID = code
	if stack != nil {
		jErr.Stack = stack.Error()
	}
	jErr.Message = GetErrorMessage(code)
	return
}
