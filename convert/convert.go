package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// BoolValue -
func BoolValue(value interface{}) (bool, error) {
	if value == nil {
		return false, fmt.Errorf("value is nil")
	}

	res, ok := value.(bool)
	if ok {
		return res, nil
	}

	return false, fmt.Errorf("unexpected type of bool value, got %v", reflect.TypeOf(value).Kind())
}

// BoolDefault -
func BoolDefault(value interface{}, defaultValue bool) bool {
	res, err := BoolValue(value)
	if err != nil {
		return defaultValue
	}

	return res
}

// Bool -
func Bool(value interface{}) bool {
	return BoolDefault(value, false)
}

// BoolForce -
func BoolForce(value interface{}) bool {
	if value == nil {
		return false
	}

	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Bool:
		return value.(bool)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int(value) != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Uint(value) != 0
	case reflect.Float32, reflect.Float64:
		return Float(value) != 0
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map, reflect.String:
		return reflect.ValueOf(value).Len() != 0
	default:
		return true
	}
}

// IntValue -
func IntValue(value interface{}) (int64, error) {
	if value == nil {
		return 0, fmt.Errorf("value is nil")
	}

	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case reflect.Int:
		return int64(value.(int)), nil
	case reflect.Int8:
		return int64(value.(int8)), nil
	case reflect.Int16:
		return int64(value.(int16)), nil
	case reflect.Int32:
		return int64(value.(int32)), nil
	case reflect.Int64:
		return value.(int64), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(Uint(value)), nil
	case reflect.Float32, reflect.Float64:
		return int64(Float(value)), nil
	case reflect.String:
		res, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return res, nil
	default:
		return 0, fmt.Errorf("unexpected type of int value, got %v", kind)
	}
}

// IntDefault -
func IntDefault(value interface{}, defaultValue int64) int64 {
	res, err := IntValue(value)
	if err != nil {
		return defaultValue
	}
	return res
}

// Int -
func Int(value interface{}) int64 {
	return IntDefault(value, 0)
}

// UintValue -
func UintValue(value interface{}) (uint64, error) {
	if value == nil {
		return 0, fmt.Errorf("value is nil")
	}

	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case reflect.Uint:
		return uint64(value.(uint)), nil
	case reflect.Uint8:
		return uint64(value.(uint8)), nil
	case reflect.Uint16:
		return uint64(value.(uint16)), nil
	case reflect.Uint32:
		return uint64(value.(uint32)), nil
	case reflect.Uint64:
		return value.(uint64), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(Int(value)), nil
	case reflect.Float32, reflect.Float64:
		return uint64(Float(value)), nil
	case reflect.String:
		res, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return res, nil
	default:
		return 0, fmt.Errorf("unexpected type of uint value, got %v", kind)
	}
}

// UintDefault -
func UintDefault(value interface{}, defaultValue uint64) uint64 {
	res, err := UintValue(value)
	if err != nil {
		return defaultValue
	}
	return res
}

// Uint -
func Uint(value interface{}) uint64 {
	return UintDefault(value, 0)
}

// FloatValue -
func FloatValue(value interface{}) (float64, error) {
	if value == nil {
		return 0, fmt.Errorf("value is nil")
	}

	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case reflect.Float32:
		return float64(value.(float32)), nil
	case reflect.Float64:
		return value.(float64), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(Int(value)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(Uint(value)), nil
	case reflect.String:
		res, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0, err
		}
		return res, nil
	default:
		return 0, fmt.Errorf("unexpected type of float value, got %v", kind)
	}
}

// FloatDefault -
func FloatDefault(value interface{}, defaultValue float64) float64 {
	res, err := FloatValue(value)
	if err != nil {
		return defaultValue
	}
	return res
}

// Float -
func Float(value interface{}) float64 {
	return FloatDefault(value, 0)
}

// StringValue -
func StringValue(value interface{}) (string, error) {
	if value == nil {
		return "", fmt.Errorf("value is nil")
	}

	res, ok := value.(string)
	if ok {
		return res, nil
	}

	return "", fmt.Errorf("unexpected type of string value, got %v", reflect.TypeOf(value).Kind())
}

// StringDefault -
func StringDefault(value interface{}, defaultValue string) string {
	res, err := StringValue(value)
	if err != nil {
		return defaultValue
	}
	return res
}

// String -
func String(value interface{}) string {
	return StringDefault(value, "")
}

// StringForce -
func StringForce(value interface{}) string {
	if value == nil {
		return ""
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return value.(string)
	default:
		return fmt.Sprintf("%v", value)
	}
}

// UnsafeBytesToString -
func UnsafeBytesToString(value []byte) string {
	return *(*string)(unsafe.Pointer(&value))
}

// UnsafeStringToBytes -
func UnsafeStringToBytes(value string) []byte {
	return *(*[]byte)(unsafe.Pointer(&value))
}
