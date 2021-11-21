package client

import "errors"

var (
	ErrMakingRequest                = errors.New("error making http request")
	ErrExecutingRequest             = errors.New("error executing http request")
	ErrUnexpectedResponseStatusCode = errors.New("unexpected response status code")
	ErrParsingResponse              = errors.New("error parsing response")
)
