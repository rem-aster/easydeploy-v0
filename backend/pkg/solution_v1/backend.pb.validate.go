// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: backend.proto

package solution_v1

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

// Validate checks the field values on SolutionInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SolutionInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SolutionInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SolutionInfoMultiError, or
// nil if none found.
func (m *SolutionInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *SolutionInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 50 {
		err := SolutionInfoValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDescription()); l < 1 || l > 500 {
		err := SolutionInfoValidationError{
			field:  "Description",
			reason: "value length must be between 1 and 500 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SolutionInfoMultiError(errors)
	}

	return nil
}

// SolutionInfoMultiError is an error wrapping multiple validation errors
// returned by SolutionInfo.ValidateAll() if the designated constraints aren't met.
type SolutionInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SolutionInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SolutionInfoMultiError) AllErrors() []error { return m }

// SolutionInfoValidationError is the validation error returned by
// SolutionInfo.Validate if the designated constraints aren't met.
type SolutionInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SolutionInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SolutionInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SolutionInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SolutionInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SolutionInfoValidationError) ErrorName() string { return "SolutionInfoValidationError" }

// Error satisfies the builtin error interface
func (e SolutionInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSolutionInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SolutionInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SolutionInfoValidationError{}

// Validate checks the field values on Solution with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Solution) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Solution with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SolutionMultiError, or nil
// if none found.
func (m *Solution) ValidateAll() error {
	return m.validate(true)
}

func (m *Solution) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SolutionValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SolutionValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SolutionValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SolutionValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SolutionValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SolutionValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SolutionValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SolutionValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SolutionValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SolutionMultiError(errors)
	}

	return nil
}

// SolutionMultiError is an error wrapping multiple validation errors returned
// by Solution.ValidateAll() if the designated constraints aren't met.
type SolutionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SolutionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SolutionMultiError) AllErrors() []error { return m }

// SolutionValidationError is the validation error returned by
// Solution.Validate if the designated constraints aren't met.
type SolutionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SolutionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SolutionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SolutionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SolutionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SolutionValidationError) ErrorName() string { return "SolutionValidationError" }

// Error satisfies the builtin error interface
func (e SolutionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSolution.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SolutionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SolutionValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListResponseMultiError, or
// nil if none found.
func (m *ListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSolutions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Solutions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Solutions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResponseValidationError{
					field:  fmt.Sprintf("Solutions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListResponseMultiError(errors)
	}

	return nil
}

// ListResponseMultiError is an error wrapping multiple validation errors
// returned by ListResponse.ValidateAll() if the designated constraints aren't met.
type ListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResponseMultiError) AllErrors() []error { return m }

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}

// Validate checks the field values on DeployRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeployRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeployRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeployRequestMultiError, or
// nil if none found.
func (m *DeployRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeployRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SolutionId

	// no validation rules for SshAddress

	// no validation rules for SshKey

	// no validation rules for ExtraVars

	if len(errors) > 0 {
		return DeployRequestMultiError(errors)
	}

	return nil
}

// DeployRequestMultiError is an error wrapping multiple validation errors
// returned by DeployRequest.ValidateAll() if the designated constraints
// aren't met.
type DeployRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeployRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeployRequestMultiError) AllErrors() []error { return m }

// DeployRequestValidationError is the validation error returned by
// DeployRequest.Validate if the designated constraints aren't met.
type DeployRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeployRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeployRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeployRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeployRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeployRequestValidationError) ErrorName() string { return "DeployRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeployRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeployRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeployRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeployRequestValidationError{}

// Validate checks the field values on DeployResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeployResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeployResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeployResponseMultiError,
// or nil if none found.
func (m *DeployResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeployResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeployResponseMultiError(errors)
	}

	return nil
}

// DeployResponseMultiError is an error wrapping multiple validation errors
// returned by DeployResponse.ValidateAll() if the designated constraints
// aren't met.
type DeployResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeployResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeployResponseMultiError) AllErrors() []error { return m }

// DeployResponseValidationError is the validation error returned by
// DeployResponse.Validate if the designated constraints aren't met.
type DeployResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeployResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeployResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeployResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeployResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeployResponseValidationError) ErrorName() string { return "DeployResponseValidationError" }

// Error satisfies the builtin error interface
func (e DeployResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeployResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeployResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeployResponseValidationError{}

// Validate checks the field values on GetDeployStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeployStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeployStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeployStatusRequestMultiError, or nil if none found.
func (m *GetDeployStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeployStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetDeployStatusRequestMultiError(errors)
	}

	return nil
}

// GetDeployStatusRequestMultiError is an error wrapping multiple validation
// errors returned by GetDeployStatusRequest.ValidateAll() if the designated
// constraints aren't met.
type GetDeployStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeployStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeployStatusRequestMultiError) AllErrors() []error { return m }

// GetDeployStatusRequestValidationError is the validation error returned by
// GetDeployStatusRequest.Validate if the designated constraints aren't met.
type GetDeployStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeployStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeployStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeployStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeployStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeployStatusRequestValidationError) ErrorName() string {
	return "GetDeployStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeployStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeployStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeployStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeployStatusRequestValidationError{}

// Validate checks the field values on GetDeployStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeployStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeployStatusResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeployStatusResponseMultiError, or nil if none found.
func (m *GetDeployStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeployStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Error

	if len(errors) > 0 {
		return GetDeployStatusResponseMultiError(errors)
	}

	return nil
}

// GetDeployStatusResponseMultiError is an error wrapping multiple validation
// errors returned by GetDeployStatusResponse.ValidateAll() if the designated
// constraints aren't met.
type GetDeployStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeployStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeployStatusResponseMultiError) AllErrors() []error { return m }

// GetDeployStatusResponseValidationError is the validation error returned by
// GetDeployStatusResponse.Validate if the designated constraints aren't met.
type GetDeployStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeployStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeployStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeployStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeployStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeployStatusResponseValidationError) ErrorName() string {
	return "GetDeployStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeployStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeployStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeployStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeployStatusResponseValidationError{}
