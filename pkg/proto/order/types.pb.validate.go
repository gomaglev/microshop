// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/proto/order/types.proto

package order

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

// Validate checks the field values on Order with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Order) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Status

	if v, ok := interface{}(m.GetDeliveredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderValidationError{
				field:  "DeliveredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OrderValidationError is the validation error returned by Order.Validate if
// the designated constraints aren't met.
type OrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderValidationError) ErrorName() string { return "OrderValidationError" }

// Error satisfies the builtin error interface
func (e OrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderValidationError{}

// Validate checks the field values on Orders with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Orders) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrdersValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrdersValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OrdersValidationError is the validation error returned by Orders.Validate if
// the designated constraints aren't met.
type OrdersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrdersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrdersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrdersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrdersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrdersValidationError) ErrorName() string { return "OrdersValidationError" }

// Error satisfies the builtin error interface
func (e OrdersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrders.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrdersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrdersValidationError{}
