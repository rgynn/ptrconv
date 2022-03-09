package ptrconv

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// IntPtr simply converts int to *int
func IntPtr(input int) *int {
	return &input
}

// IntPtrString returns the string value of *int or nil
func IntPtrString(input *int) string {
	if input == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", *input)
}

// IntPtrFromString simply converts a string to to *int
func IntPtrFromString(input string) (*int, error) {
	result, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// IntPtrNilOrEmpty returns true if nil or empty int
func IntPtrNilOrEmpty(input *int) bool {
	return input == nil || input != nil && *input == 0
}

// Int8Ptr simply converts int8 to *int8
func Int8Ptr(input int8) *int8 {
	return &input
}

// Int8PtrString returns the string value of *int8 or nil
func Int8PtrString(input *int8) string {
	if input == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", *input)
}

// Int8PtrFromString simply converts a string to to *int8
func Int8PtrFromString(input string) (*int8, error) {
	result, err := strconv.ParseInt(input, 10, 8)
	if err != nil {
		return nil, err
	}
	result8 := int8(result)
	return &result8, nil
}

// Int8PtrNilOrEmpty returns true if nil or empty int8
func Int8PtrNilOrEmpty(input *int8) bool {
	return input == nil || input != nil && *input == 0
}

// Int16Ptr simply converts int16 to *int16
func Int16Ptr(input int64) *int64 {
	return &input
}

// Int16PtrString returns the string value of *int16 or nil
func Int16PtrString(input *int16) string {
	if input == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", *input)
}

// Int16PtrFromString simply converts a string to to *int16
func Int16PtrFromString(input string) (*int16, error) {
	result, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		return nil, err
	}
	result16 := int16(result)
	return &result16, nil
}

// Int16PtrNilOrEmpty returns true if nil or empty int16
func Int16PtrNilOrEmpty(input *int16) bool {
	return input == nil || input != nil && *input == 0
}

// Int32Ptr simply converts int32 to *int32
func Int32Ptr(input int32) *int32 {
	return &input
}

// Int32PtrString returns the string value of *int32 or nil
func Int32PtrString(input *int32) string {
	if input == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", *input)
}

// Int32PtrFromString simply converts a string to to *int32
func Int32PtrFromString(input string) (*int32, error) {
	result, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return nil, err
	}
	result32 := int32(result)
	return &result32, nil
}

// Int32PtrNilOrEmpty returns true if nil or empty int32
func Int32PtrNilOrEmpty(input *int32) bool {
	return input == nil || input != nil && *input == 0
}

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

// DurationPtr simply converts time.Duration to *time.Duration
func DurationPtr(input time.Duration) *time.Duration {
	return &input
}

// DurationPtrFromString simply converts a string to to *time.Duration
func DurationPtrFromString(input string) (*time.Duration, error) {
	result, err := time.ParseDuration(input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DurationPtrString returns the string value of *time.Duration or nil
func DurationPtrString(input *time.Duration) string {
	if input == nil {
		return "nil"
	}
	return input.String()
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
