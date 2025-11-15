package testutil

import (
	"reflect"
	"strings"
	"testing"
)

type IntNumber interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type FloatNumber interface {
	float32 | float64
}

const (
	equalFailedMessageTemplate    = "Element %s is not as expected.\nexpected: '%v'\nactual:   '%v'"
	notEqualFailedMessageTemplate = "Element %s equals the unexpected.\nunexpected: '%v'\nactual:     '%v'"
)

// Checks that 'actual' is not nil. If it is nil [testing.T.Errorf] will be called
func AssertNotNil(actual any, t *testing.T, objectName string) {
	t.Helper()
	if isNil(actual) {
		t.Errorf("Element %s should not be nil", objectName)
	}
}

// Checks that 'actual' is nil. If it is not nil [testing.T.Errorf] will be called
func AssertNil(actual any, t *testing.T, objectName string) {
	t.Helper()
	if !isNil(actual) {
		t.Errorf("Element %s should be nil, but has value '%v'", objectName, actual)
	}
}

// Checks that 'actual' is true. If it is false [testing.T.Errorf] will be called
func AssertTrue(actual bool, t *testing.T, objectName string) {
	t.Helper()
	AssertEquals(true, actual, t, objectName)
}

// Checks that 'actual' is false. If it is true [testing.T.Errorf] will be called
func AssertFalse(actual bool, t *testing.T, objectName string) {
	t.Helper()
	AssertEquals(false, actual, t, objectName)
}

// Checks that 'actual' is not nil and equal to 'expected'. If it is nil or not equal [testing.T.Errorf] will be called
func AssertEquals(expected any, actual any, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if expected != actual {
		t.Errorf(equalFailedMessageTemplate, objectName, expected, actual)
	}
}

// Checks that 'actual' is not nil and equal to 'expected' when both are transformed to an int64/uint64 before comparison. If it is nil or not equal [testing.T.Errorf] will be called
func AssertAlmostEqualsInt[X IntNumber, Y IntNumber](expected X, actual Y, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if isNotAlmostEqualInt(expected, actual) {
		t.Errorf(equalFailedMessageTemplate, objectName, expected, actual)
	}
}

// Checks that 'actual' is not nil and equal to 'expected' when both are transformed to an float64 before comparison with respect to a given delta. If it is nil or not equal [testing.T.Errorf] will be called
func AssertAlmostEqualsFloat[X FloatNumber, Y FloatNumber](expected X, actual Y, delta float64, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if isNotAlmostEqualFloat(expected, actual, delta) {
		t.Errorf(equalFailedMessageTemplate, objectName, expected, actual)
	}
}

// Checks that 'actual' is not nil and not equal to 'expected'. If it is nil or equal [testing.T.Errorf] will be called
func AssertNotEquals(notExpected any, actual any, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if notExpected == actual {
		t.Errorf(notEqualFailedMessageTemplate, objectName, notExpected, actual)
	}
}

// Checks that 'actual' is not nil and not equal to 'expected' when both are transformed to an int64/uint64 before comparison. If it is nil or equal [testing.T.Errorf] will be called
func AssertNotAlmostEqualsInt[X IntNumber, Y IntNumber](notExpected X, actual Y, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if !isNotAlmostEqualInt(notExpected, actual) {
		t.Errorf(notEqualFailedMessageTemplate, objectName, notExpected, actual)
	}
}

// Checks that 'actual' is not nil and not equal to 'expected' when both are transformed to an int64/uint64 before comparison with respect to a given delta. If it is nil or equal [testing.T.Errorf] will be called
func AssertNotAlmostEqualsFloat[X FloatNumber, Y FloatNumber](notExpected X, actual Y, delta float64, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if !isNotAlmostEqualFloat(notExpected, actual, delta) {
		t.Errorf(notEqualFailedMessageTemplate, objectName, notExpected, actual)
	}
}

// Checks that 'actual' is not nil and has 'expected' as prefix. If it is nil or does not have the correct prefix [testing.T.Errorf] will be called
func AssertHasPrefix(expected string, actual string, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("Element %s is not as expected. expected prefix: '%v' actual: '%v'", objectName, expected, actual)
	}
}

// Checks that 'actual' is not nil and has 'expected' as suffix. If it is nil or does not have the correct suffix [testing.T.Errorf] will be called
func AssertHasSuffix(expected string, actual string, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if !strings.HasSuffix(actual, expected) {
		t.Errorf("Element %s is not as expected. expected suffix: '%v' actual: '%v'", objectName, expected, actual)
	}
}

// Checks that 'actual' is not nil and has 'expectedToContain' is contained at 'actual'. If it is nil or does not contain the substring [testing.T.Errorf] will be called
func AssertContains(expectedToContain string, actual string, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if !strings.Contains(actual, expectedToContain) {
		t.Errorf("Element %s does not contain expected string. expected to be contained: '%v' actual: '%v'", objectName, expectedToContain, actual)
	}
}

// Checks that 'actual' is not nil and has 'expectedToContain' is not contained at 'actual'. If it is nil or does contain the substring [testing.T.Errorf] will be called
func AssertNotContains(expectedToContain string, actual string, t *testing.T, objectName string) {
	t.Helper()
	AssertNotNil(actual, t, objectName)
	if !strings.Contains(actual, expectedToContain) {
		t.Errorf("Element %s does contain expected string. expected not to be contained: '%v' actual: '%v'", objectName, expectedToContain, actual)
	}
}

func isNil(toCheck any) bool {
	v := reflect.ValueOf(toCheck)
	k := v.Kind()

	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return toCheck == nil
	}
}

func getKind(i any) reflect.Kind {
	v := reflect.ValueOf(i)
	return v.Kind()
}

func isNotAlmostEqualInt[X IntNumber, Y IntNumber](i1 X, i2 Y) bool {
	isUInt1 := isUInt(getKind(i1))
	isUInt2 := isUInt(getKind(i2))
	return (isUInt1 && !isUInt2 && i2 < 0) || (!isUInt1 && isUInt2 && i1 < 0) || (!isUInt1 && !isUInt2 && int64(i1) != int64(i2)) || uint64(i1) != uint64(i2)
}

func isNotAlmostEqualFloat[X FloatNumber, Y FloatNumber](f1 X, f2 Y, delta float64) bool {
	float1 := float64(f1)
	float2 := float64(f2)
	return (float1 <= float2 && float2-float1 > delta) || (float2 < float1 && float1-float2 > delta)
}

func isUInt(k reflect.Kind) bool {
	return k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64
}
