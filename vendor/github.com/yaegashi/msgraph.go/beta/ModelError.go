// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ErrorDetail undocumented
type ErrorDetail struct {
	// Object is the base model of ErrorDetail
	Object
	// ErrorCode undocumented
	ErrorCode *string `json:"errorCode,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Details undocumented
	Details []InnerErrorDetail `json:"details,omitempty"`
}
