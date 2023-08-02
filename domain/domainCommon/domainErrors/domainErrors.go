package domainErrors

import (
	"fmt"
	"net/http"
)

type ErrorModel struct {
	Code          int    `json:"code,omitempty"`
	Message       string `json:"message,omitempty"`
	Result        string `json:"result,omitempty"`
	ReferenceCode string `json:"referenceCode,omitempty"`
	HttpResp      int    `json:"-"`
}

func (e *ErrorModel) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

var (
	BadRequest          = ErrorModel{Code: 4001, Message: "BAD_REQUEST", HttpResp: http.StatusBadRequest}
	Unauthorized        = ErrorModel{Code: 4011, Message: "UNAUTHORIZED", HttpResp: http.StatusUnauthorized}
	PermissionDenied    = ErrorModel{Code: 4031, Message: "PERMISSION_DENIED", HttpResp: http.StatusForbidden}
	MethodNotAllowed    = ErrorModel{Code: 4051, Message: "METHOD_NOT_ALLOWED", HttpResp: http.StatusMethodNotAllowed}
	AccountNotVerified  = ErrorModel{Code: 4012, Message: "ACCOUNT_NOT_VERIFIED", HttpResp: http.StatusUnauthorized}
	SessionNotActive    = ErrorModel{Code: 4013, Message: "SESSION_NOT_ACTIVE", HttpResp: http.StatusUnauthorized}
	InvalidRefreshToken = ErrorModel{Code: 4014, Message: "INVALID_REFRESH_TOKEN", HttpResp: http.StatusUnauthorized}
	InternalError       = ErrorModel{Code: 5001, Message: "INTERNAL_ERROR", HttpResp: http.StatusInternalServerError}
	ExpiredRequestToken = ErrorModel{Code: 4002, Message: "EXPIRED_REQUEST_TOKEN", HttpResp: http.StatusBadRequest}
	InvalidRequestToken = ErrorModel{Code: 4003, Message: "INVALID_REQUEST_TOKEN", HttpResp: http.StatusBadRequest}
	ValidationFailed    = ErrorModel{Code: 4004, Message: "VALIDATION_FAILED", HttpResp: http.StatusBadRequest}
	ResourceNotFound    = ErrorModel{Code: 4041, Message: "NOT FOUND", HttpResp: http.StatusNotFound}
)

func CheckResponseStatus(statusCode int, result string) error {
	var errorModel *ErrorModel
	switch statusCode {
	case 200, 201, 204:
		return nil
	case 400, 4001:
		errorModel = &BadRequest
	case 401, 4011:
		errorModel = &Unauthorized
	case 403, 4031:
		errorModel = &PermissionDenied
	case 404, 4041:
		errorModel = &ResourceNotFound
	case 405, 4051:
		errorModel = &MethodNotAllowed
	case 4012:
		errorModel = &AccountNotVerified
	case 4013:
		errorModel = &SessionNotActive
	case 4014:
		errorModel = &InvalidRefreshToken
	case 4002:
		errorModel = &ExpiredRequestToken
	case 4003:
		errorModel = &InvalidRequestToken
	case 4004:
		errorModel = &ValidationFailed
	default:
		errorModel = &InternalError
	}

	if result != "" {
		errorModel.Result = result
	}
	return errorModel
}

type InsufficientArgument struct {
	ArgumentName string
}

func (ia *InsufficientArgument) Error() string {
	return fmt.Sprintf("Error: insufficient argument '%s'", ia.ArgumentName)
}

type InvalidArgument struct {
	ArgumentName string
}

func (ia *InvalidArgument) Error() string {
	return fmt.Sprintf("Error: invalid argument '%s'", ia.ArgumentName)
}

func IsDomainError() string {
	return "*domainErrors.ErrorModel"
}
