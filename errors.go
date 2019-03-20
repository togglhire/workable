package workable

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrNotImplemented is used for returning errors
var (
	ErrClientIDMissing     = errors.New("client id is missing")
	ErrShouldNotBeNil      = errors.New("should not be nil")
	ErrClientIsNil         = errors.New("client is nill, use NewClient function for initializing")
	ErrRefreshTokenMissing = errors.New("refresh token is missing")
)

type Error struct {
	Error struct {
		Name  string `json:"name"`
		State string `json:"state"`
	} `json:"error"`
	RedirectURI        string `json:"redirect_uri"`
	ResponseOnFragment string `json:"response_on_fragment"`
	Reason             string `json:"reason"`
	Description        string `json:"description"`
}

type ClientError struct {
	StatusCode   int
	ErrorMessage Error
}

func (e ClientError) Error() string {
	return fmt.Sprintf("Error: %#+v\n", e.ErrorMessage)
}

type ServerError struct {
	StatusCode   int
	ErrorMessage Error
}

func (e ServerError) Error() string {
	return fmt.Sprintf("Error: %#+v\n", e.ErrorMessage)
}

func isOK(r *http.Response) (bool, error) {
	if r == nil {
		return false, ErrShouldNotBeNil
	}
	return r.StatusCode >= 200 && r.StatusCode <= 299, nil
}

func isError(r *http.Response) (bool, error) {
	if r == nil {
		return false, ErrShouldNotBeNil
	}
	ok, err := isOK(r)
	return !ok, err
}

func isClientError(r *http.Response) (bool, error) {
	if r == nil {
		return false, ErrShouldNotBeNil
	}
	return r.StatusCode >= 400 && r.StatusCode <= 499, nil
}

func isServerError(r *http.Response) (bool, error) {
	if r == nil {
		return false, ErrShouldNotBeNil
	}
	return r.StatusCode >= 500 && r.StatusCode <= 599, nil
}

// IsClientError checks whether the error was a Client error or not. If it was a client error, the first return param is the value.
func IsClientError(err interface{}) (ClientError, bool) {
	s, ok := err.(ClientError)
	return s, ok
}

// IsServerError checks whether the error was a Server error or not. If it was a server error, the first return param is the value.
func IsServerError(err interface{}) (ServerError, bool) {
	s, ok := err.(ServerError)
	return s, ok
}
