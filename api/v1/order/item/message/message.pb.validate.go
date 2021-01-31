// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/order/item/message/message.proto

package message

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _message_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetMessageRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetMessageRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OrderId

	// no validation rules for ItemId

	// no validation rules for Id

	return nil
}

// GetMessageRequestValidationError is the validation error returned by
// GetMessageRequest.Validate if the designated constraints aren't met.
type GetMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMessageRequestValidationError) ErrorName() string {
	return "GetMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMessageRequestValidationError{}

// Validate checks the field values on GetMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetMessageResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetMessageResponseValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetMessageResponseValidationError is the validation error returned by
// GetMessageResponse.Validate if the designated constraints aren't met.
type GetMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMessageResponseValidationError) ErrorName() string {
	return "GetMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMessageResponseValidationError{}

// Validate checks the field values on ListMessagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListMessagesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OrderId

	// no validation rules for ItemId

	if m.GetPagination() == nil {
		return ListMessagesRequestValidationError{
			field:  "Pagination",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListMessagesRequestValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListMessagesRequestValidationError is the validation error returned by
// ListMessagesRequest.Validate if the designated constraints aren't met.
type ListMessagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMessagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMessagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMessagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMessagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMessagesRequestValidationError) ErrorName() string {
	return "ListMessagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListMessagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMessagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMessagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMessagesRequestValidationError{}

// Validate checks the field values on ListMessagesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListMessagesResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMessages()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListMessagesResponseValidationError{
				field:  "Messages",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListMessagesResponseValidationError is the validation error returned by
// ListMessagesResponse.Validate if the designated constraints aren't met.
type ListMessagesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMessagesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMessagesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMessagesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMessagesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMessagesResponseValidationError) ErrorName() string {
	return "ListMessagesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListMessagesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMessagesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMessagesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMessagesResponseValidationError{}

// Validate checks the field values on CreatMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatMessageRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OrderId

	// no validation rules for ItemId

	if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatMessageRequestValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreatMessageRequestValidationError is the validation error returned by
// CreatMessageRequest.Validate if the designated constraints aren't met.
type CreatMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatMessageRequestValidationError) ErrorName() string {
	return "CreatMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatMessageRequestValidationError{}

// Validate checks the field values on CreatMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatMessageResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreatMessageResponseValidationError is the validation error returned by
// CreatMessageResponse.Validate if the designated constraints aren't met.
type CreatMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatMessageResponseValidationError) ErrorName() string {
	return "CreatMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatMessageResponseValidationError{}

// Validate checks the field values on UpdateMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateMessageRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OrderId

	// no validation rules for ItemId

	// no validation rules for Id

	if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateMessageRequestValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateMessageRequestValidationError is the validation error returned by
// UpdateMessageRequest.Validate if the designated constraints aren't met.
type UpdateMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMessageRequestValidationError) ErrorName() string {
	return "UpdateMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMessageRequestValidationError{}

// Validate checks the field values on UpdateMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateMessageResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Updated

	return nil
}

// UpdateMessageResponseValidationError is the validation error returned by
// UpdateMessageResponse.Validate if the designated constraints aren't met.
type UpdateMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMessageResponseValidationError) ErrorName() string {
	return "UpdateMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMessageResponseValidationError{}

// Validate checks the field values on DeleteMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteMessageRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OrderId

	// no validation rules for ItemId

	// no validation rules for Id

	return nil
}

// DeleteMessageRequestValidationError is the validation error returned by
// DeleteMessageRequest.Validate if the designated constraints aren't met.
type DeleteMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMessageRequestValidationError) ErrorName() string {
	return "DeleteMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMessageRequestValidationError{}

// Validate checks the field values on DeleteMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteMessageResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Deleted

	return nil
}

// DeleteMessageResponseValidationError is the validation error returned by
// DeleteMessageResponse.Validate if the designated constraints aren't met.
type DeleteMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMessageResponseValidationError) ErrorName() string {
	return "DeleteMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMessageResponseValidationError{}