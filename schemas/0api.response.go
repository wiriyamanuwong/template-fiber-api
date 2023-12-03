// Package schemas for Request and Response Schema
package schemas

// APIResponse Default API Response
type APIResponse struct {
	OK      bool   `json:"ok"`
	Code    int    `json:"code"`
	Message string `json:"message"`
} // @name APIResponse

// NewAPIResponse create new APIResponse
func NewAPIResponse(code int, message string) *APIResponse {
	return &APIResponse{
		OK:      (code >= 200 && code < 300),
		Code:    code,
		Message: message,
	}
}

// GetsAPIResponse Default API Response for get data multiple result
type GetsAPIResponse[T any] struct {
	APIResponse
	Data       []T         `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

// GetOneAPIResponse Default API Response for get date by id
type GetOneAPIResponse[T any] struct {
	APIResponse
	Data T `json:"data"`
}
