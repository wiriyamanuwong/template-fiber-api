package schemas

// NewError create new error
func NewError(code int, message string) *APIResponse {
	return &APIResponse{
		OK:      false,
		Code:    code,
		Message: message,
	}
}
