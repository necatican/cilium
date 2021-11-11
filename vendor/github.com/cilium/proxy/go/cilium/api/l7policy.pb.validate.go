// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cilium/api/l7policy.proto

package cilium

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on L7Policy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *L7Policy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on L7Policy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in L7PolicyMultiError, or nil
// if none found.
func (m *L7Policy) ValidateAll() error {
	return m.validate(true)
}

func (m *L7Policy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessLogPath

	// no validation rules for PolicyName

	// no validation rules for Denied_403Body

	if all {
		switch v := interface{}(m.GetIsIngress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, L7PolicyValidationError{
					field:  "IsIngress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, L7PolicyValidationError{
					field:  "IsIngress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIsIngress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return L7PolicyValidationError{
				field:  "IsIngress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return L7PolicyMultiError(errors)
	}
	return nil
}

// L7PolicyMultiError is an error wrapping multiple validation errors returned
// by L7Policy.ValidateAll() if the designated constraints aren't met.
type L7PolicyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m L7PolicyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m L7PolicyMultiError) AllErrors() []error { return m }

// L7PolicyValidationError is the validation error returned by
// L7Policy.Validate if the designated constraints aren't met.
type L7PolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e L7PolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e L7PolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e L7PolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e L7PolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e L7PolicyValidationError) ErrorName() string { return "L7PolicyValidationError" }

// Error satisfies the builtin error interface
func (e L7PolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sL7Policy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = L7PolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = L7PolicyValidationError{}
