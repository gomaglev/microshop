// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/proto/order/item/types.proto

package item

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
var _types_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on OrderItem with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OrderItem) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Price

	// no validation rules for Quantity

	// no validation rules for ProductId

	// no validation rules for OrderId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderItemValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderItemValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OrderItemValidationError is the validation error returned by
// OrderItem.Validate if the designated constraints aren't met.
type OrderItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderItemValidationError) ErrorName() string { return "OrderItemValidationError" }

// Error satisfies the builtin error interface
func (e OrderItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderItemValidationError{}

// Validate checks the field values on OrderItems with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OrderItems) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderItemsValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderItemsValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OrderItemsValidationError is the validation error returned by
// OrderItems.Validate if the designated constraints aren't met.
type OrderItemsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderItemsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderItemsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderItemsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderItemsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderItemsValidationError) ErrorName() string { return "OrderItemsValidationError" }

// Error satisfies the builtin error interface
func (e OrderItemsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderItems.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderItemsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderItemsValidationError{}
