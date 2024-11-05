// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"
)

func (s *ErrRespStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type ApiKeyAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *ApiKeyAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *ApiKeyAuth) SetAPIKey(val string) {
	s.APIKey = val
}

type ErrResp struct {
	// A detailed error message.
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *ErrResp) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *ErrResp) SetError(val OptString) {
	s.Error = val
}

// ErrRespStatusCode wraps ErrResp with StatusCode.
type ErrRespStatusCode struct {
	StatusCode int
	Response   ErrResp
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrRespStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrRespStatusCode) GetResponse() ErrResp {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrRespStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrRespStatusCode) SetResponse(val ErrResp) {
	s.Response = val
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type ViewGetBadRequest struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *ViewGetBadRequest) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *ViewGetBadRequest) SetError(val OptString) {
	s.Error = val
}

func (*ViewGetBadRequest) viewGetRes() {}

// ViewGetInternalServerError is response for ViewGet operation.
type ViewGetInternalServerError struct{}

func (*ViewGetInternalServerError) viewGetRes() {}

type ViewGetNotFound struct {
	// A detailed error message.
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *ViewGetNotFound) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *ViewGetNotFound) SetError(val OptString) {
	s.Error = val
}

func (*ViewGetNotFound) viewGetRes() {}

type ViewGetOK struct {
	Result []ViewGetOKResultItem `json:"result"`
}

// GetResult returns the value of Result.
func (s *ViewGetOK) GetResult() []ViewGetOKResultItem {
	return s.Result
}

// SetResult sets the value of Result.
func (s *ViewGetOK) SetResult(val []ViewGetOKResultItem) {
	s.Result = val
}

func (*ViewGetOK) viewGetRes() {}

type ViewGetOKResultItem struct {
	Position ViewGetOKResultItemPosition `json:"position"`
	// Presigned URL for target position.
	URL string `json:"url"`
}

// GetPosition returns the value of Position.
func (s *ViewGetOKResultItem) GetPosition() ViewGetOKResultItemPosition {
	return s.Position
}

// GetURL returns the value of URL.
func (s *ViewGetOKResultItem) GetURL() string {
	return s.URL
}

// SetPosition sets the value of Position.
func (s *ViewGetOKResultItem) SetPosition(val ViewGetOKResultItemPosition) {
	s.Position = val
}

// SetURL sets the value of URL.
func (s *ViewGetOKResultItem) SetURL(val string) {
	s.URL = val
}

type ViewGetOKResultItemPosition struct {
	// Target of x position.
	X int `json:"x"`
	// Target of y position.
	Y int `json:"y"`
}

// GetX returns the value of X.
func (s *ViewGetOKResultItemPosition) GetX() int {
	return s.X
}

// GetY returns the value of Y.
func (s *ViewGetOKResultItemPosition) GetY() int {
	return s.Y
}

// SetX sets the value of X.
func (s *ViewGetOKResultItemPosition) SetX(val int) {
	s.X = val
}

// SetY sets the value of Y.
func (s *ViewGetOKResultItemPosition) SetY(val int) {
	s.Y = val
}
