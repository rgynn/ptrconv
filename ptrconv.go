package ptrconv

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Int32Ptr simply converts int32 to *int32
func Int32Ptr(i int32) *int32 { return &i }

// BoolPtr simply converts bool to *bool
func BoolPtr(input bool) *bool {
	return &input
}

// BoolPtrString returns the string value of *bool or nil
func BoolPtrString(input *bool) string {
	if input == nil {
		return "nil"
	}
	return fmt.Sprintf("%t", *input)
}

// BoolPtrFromString simply converts a string to to *bool
func BoolPtrFromString(input string) (*bool, error) {
	result, err := strconv.ParseBool(input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Int64Ptr simply converts int64 to *int64
func Int64Ptr(input int64) *int64 {
	return &input
}

// Int64PtrString returns the string value of *int64 or nil
func Int64PtrString(input *int64) string {
	if input == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", *input)
}

// Int64PtrFromString simply converts a string to to *int64
func Int64PtrFromString(input string) (*int64, error) {
	result, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Int64PtrNilOrEmpty returns true if nil or empty int64
func Int64PtrNilOrEmpty(input *int64) bool {
	return input == nil || input != nil && *input == 0
}

// TimePtr simply converts time.Time to *time.Time
func TimePtr(input time.Time) *time.Time {
	return &input
}

// TimePtrFromString simply converts a string to to *time.Time with RFC3339 format
func TimePtrFromString(input string) (*time.Time, error) {
	result, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// TimePtrString returns the string value of *time.Time in time.RFC3339 format or nil
func TimePtrString(input *time.Time) string {
	if input == nil {
		return "nil"
	}
	return input.Format(time.RFC3339)
}

// TimePtrNilOrEmpty returns true if nil or empty time.Time
func TimePtrNilOrEmpty(input *time.Time) bool {
	return input == nil || input != nil && input.IsZero()
}

// StringPtr simply converts string to *string
func StringPtr(input string) *string {
	return &input
}

// StringPtrString returns the string value of *string or nil
func StringPtrString(input *string) string {
	if input == nil {
		return "nil"
	}
	return *input
}

// StringPtrNilOrEmpty returns true if nil or empty string
func StringPtrNilOrEmpty(input *string) bool {
	return input == nil || input != nil && *input == ""
}

// StringSlicePtrNilOrEmpty returns true if nil or empty
func StringSlicePtrNilOrEmpty(input *[]string) bool {
	return input == nil || input != nil && len(*input) == 0
}

// IntSliceString returns the string value of []int or nil
func IntSliceString(input []int) string {
	if input == nil {
		return "nil"
	}
	return strings.Replace(fmt.Sprint(input), " ", ", ", -1)
}
