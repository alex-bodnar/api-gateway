package errs

import (
	"fmt"
	"strings"
)

const (
	errUnauthorized     = "unauthorized"
	errEmpty            = "empty"
	errBadGateway       = "bad_gateway"
	errForbidden        = "forbidden"
	errNotFound         = "not_found"
	errMethodNotAllowed = "method_not_allowed"
	errEntityExists     = "entity_exists"
	errConflict         = "conflict"
	errValidation       = "validation_error"
	errInternal         = "internal_server_error"

	divider   = " - "
	separator = ","
)

type (
	// Empty - an error that indicates that nothing was found by the specified search parameters
	Empty struct {
		What string `json:"what"`
	}

	// BadGateway - server error response code indicates that the server,
	// while acting as a gateway or proxy, received an invalid response from the upstream server.
	BadGateway struct {
		Cause string `json:"cause"`
	}

	// Unauthorized - the error determines the situation when the user is not authorized for the specified operation
	Unauthorized struct {
		Cause string `json:"cause"`
	}

	// NotFound - describes situation for any item not found
	NotFound struct {
		What string `json:"what"`
	}

	MethodNotAllowed struct {
		Cause string `json:"cause"`
	}

	// FieldsValidation - describes situation for bad request with text
	FieldsValidation struct {
		Errors []string `json:"errors"`
	}

	// BadRequest - describes situation for bad request with text
	BadRequest struct {
		Cause string `json:"cause"`
	}

	// Forbidden - describes situation for forbidden access to operation
	Forbidden struct {
		Cause string `json:"cause"`
	}

	// AlreadyExists - describes situation for item already exists
	AlreadyExists struct {
		What string `json:"what"`
	}

	// Conflict - describes situation for conflict in logic with current application state
	Conflict struct {
		Cause string `json:"what"`
	}

	// Internal - describes situation for internal server error
	Internal struct {
		Cause string `json:"cause"`
	}
)

func (e FieldsValidation) Error() string {
	if len(e.Errors) == 0 {
		return errValidation
	}

	return fmt.Sprintf("%s%s%s",
		errValidation,
		divider,
		strings.Join(e.Errors, separator),
	)
}

func (e Empty) Error() string {
	if e.What != "" {
		return format(errEmpty, e.What)
	}

	return errEmpty
}

func (e BadGateway) Error() string {
	if e.Cause != "" {
		return format(errBadGateway, e.Cause)
	}

	return errBadGateway
}

func (e Forbidden) Error() string {
	if e.Cause != "" {
		return format(errForbidden, e.Cause)
	}

	return errForbidden
}

func (e Unauthorized) Error() string {
	if e.Cause != "" {
		return format(errUnauthorized, e.Cause)
	}

	return errUnauthorized
}

func (e BadRequest) Error() string {
	return e.Cause
}

func (e NotFound) Error() string {
	if e.What != "" {
		return format(errNotFound, e.What)
	}

	return errNotFound
}

func (e MethodNotAllowed) Error() string {
	if e.Cause != "" {
		return format(errMethodNotAllowed, e.Cause)
	}

	return errMethodNotAllowed
}

func (e AlreadyExists) Error() string {
	return format(errEntityExists, e.What)
}

func (e Conflict) Error() string {
	if e.Cause != "" {
		return format(errConflict, e.Cause)
	}

	return errConflict
}

func (e Internal) Error() string {
	if e.Cause != "" {
		return format(errInternal, e.Cause)
	}

	return errInternal
}

// format formats error
func format(messages ...string) string {
	return strings.Join(messages, divider)
}
