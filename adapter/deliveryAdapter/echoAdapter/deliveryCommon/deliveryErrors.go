package deliveryCommon

import "net/http"

type RestErr struct {
	Message string
	Status  int
	Error   string
}

func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "BadRequest",
	}
}

func ConflictRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusConflict,
		Error:   "ConflictRequest",
	}
}
