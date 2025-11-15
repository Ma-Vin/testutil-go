package testutil

import (
	"testing"
	"unsafe"
)

type dummyInterface interface {
}

type dummyStruct struct {
}

func getDummyInterface() dummyInterface {
	return dummyStruct{}
}

// ----------------
// AssertNil
// ----------------
func TestAssertNilPositive(t *testing.T) {
	checkAssertNilPositive(nil, "nil", t)
}

func TestAssertNilChanPositive(t *testing.T) {
	var c chan int
	checkAssertNilPositive(c, "chan", t)
}

func TestAssertNilFuncPositive(t *testing.T) {
	var f func()
	checkAssertNilPositive(f, "func", t)
}

func TestAssertNilMapPositive(t *testing.T) {
	var m map[string]int
	checkAssertNilPositive(m, "map", t)
}

func TestAssertNilSlicePositive(t *testing.T) {
	var s []int
	checkAssertNilPositive(s, "slice", t)
}

func TestAssertNilInterfacePositive(t *testing.T) {
	var i dummyInterface
	checkAssertNilPositive(i, "interface", t)
}

func TestAssertNilPointerPositive(t *testing.T) {
	var p *int
	checkAssertNilPositive(p, "Pointer", t)
}

func TestAssertNilUnsafePointerPositive(t *testing.T) {
	var p unsafe.Pointer
	checkAssertNilPositive(p, "unsafePointer", t)
}

func checkAssertNilPositive(v any, testUnitName string, t *testing.T) {
	t.Helper()
	tc := testing.T{}
	AssertNil(v, &tc, testUnitName)
	if tc.Failed() {
		t.Errorf("AssertNil is failed at %s", testUnitName)
	}
}

func TestAssertNilChanNegative(t *testing.T) {
	c := make(chan int)
	checkAssertNilNegative(c, "chan", t)
}

func TestAssertNilFuncNegative(t *testing.T) {
	f := func() {}
	checkAssertNilNegative(f, "func", t)
}

func TestAssertNilMapNegative(t *testing.T) {
	m := make(map[string]int)
	checkAssertNilNegative(m, "map", t)
}

func TestAssertNilSliceNegative(t *testing.T) {
	s := make([]int, 1)
	checkAssertNilNegative(s, "slice", t)
}

func TestAssertNilInterfaceNegative(t *testing.T) {
	i := getDummyInterface()
	checkAssertNilNegative(i, "interface", t)
}

func TestAssertNilPointerNegative(t *testing.T) {
	i := 1
	checkAssertNilNegative(&i, "Pointer", t)
}

func TestAssertNilUnsafePointerNegative(t *testing.T) {
	i := 1
	checkAssertNilNegative(unsafe.Pointer(&i), "unsafePointer", t)
}

func checkAssertNilNegative(v any, testUnitName string, t *testing.T) {
	t.Helper()
	tc := testing.T{}
	AssertNil(v, &tc, testUnitName)
	if !tc.Failed() {
		t.Errorf("AssertNil is not failed at %s", testUnitName)
	}
}

// ----------------
// AssertNotNil
// ----------------
func TestAssertNotNilChanPositive(t *testing.T) {
	c := make(chan int)
	checkAssertNotNilNPositive(c, "chan", t)
}

func TestAssertNotNilFuncPositive(t *testing.T) {
	f := func() {}
	checkAssertNotNilNPositive(f, "func", t)
}

func TestAssertNotNilMapPositive(t *testing.T) {
	m := make(map[string]int)
	checkAssertNotNilNPositive(m, "map", t)
}

func TestAssertNotNilSlicePositive(t *testing.T) {
	s := make([]int, 1)
	checkAssertNotNilNPositive(s, "slice", t)
}

func TestAssertNotNilInterfacePositive(t *testing.T) {
	i := getDummyInterface()
	checkAssertNotNilNPositive(i, "interface", t)
}

func TestAssertNotNilPointerPositive(t *testing.T) {
	i := 1
	checkAssertNotNilNPositive(&i, "Pointer", t)
}

func TestAssertNotNilUnsafePointerPositive(t *testing.T) {
	i := 1
	checkAssertNotNilNPositive(unsafe.Pointer(&i), "unsafePointer", t)
}

func checkAssertNotNilNPositive(v any, testUnitName string, t *testing.T) {
	t.Helper()
	tc := testing.T{}
	AssertNotNil(v, &tc, testUnitName)
	if tc.Failed() {
		t.Errorf("AssertNotNil is failed at %s", testUnitName)
	}
}

func TestAssertNotNilNegative(t *testing.T) {
	checkAssertNotNilNegative(nil, "nil", t)
}

func TestAssertNotNilChanNegative(t *testing.T) {
	var c chan int
	checkAssertNotNilNegative(c, "chan", t)
}

func TestAssertNotNilFuncNegative(t *testing.T) {
	var f func()
	checkAssertNotNilNegative(f, "func", t)
}

func TestAssertNotNilMapNegative(t *testing.T) {
	var m map[string]int
	checkAssertNotNilNegative(m, "map", t)
}

func TestAssertNotNilSliceNegative(t *testing.T) {
	var s []int
	checkAssertNotNilNegative(s, "slice", t)
}

func TestAssertNotNilInterfaceNegative(t *testing.T) {
	var i dummyInterface
	checkAssertNotNilNegative(i, "interface", t)
}

func TestAssertNotNilPointerNegative(t *testing.T) {
	var p *int
	checkAssertNotNilNegative(p, "Pointer", t)
}

func TestAssertNotNilUnsafePointerNegative(t *testing.T) {
	var p unsafe.Pointer
	checkAssertNotNilNegative(p, "unsafePointer", t)
}

func checkAssertNotNilNegative(v any, testUnitName string, t *testing.T) {
	t.Helper()
	tc := testing.T{}
	AssertNotNil(v, &tc, testUnitName)
	if !tc.Failed() {
		t.Errorf("AssertNotNil is not failed at %s", testUnitName)
	}
}

// ----------------
// AssertTrue
// ----------------
func TestAssertTruePositive(t *testing.T) {
	tc := testing.T{}
	AssertTrue(true, &tc, "AssertTrue")
	if tc.Failed() {
		t.Error("AssertTrue is failed")
	}
}

func TestAssertTrueNegative(t *testing.T) {
	tc := testing.T{}
	AssertTrue(false, &tc, "AssertTrue")
	if !tc.Failed() {
		t.Error("AssertTrue is not failed")
	}
}

// ----------------
// AssertFalse
// ----------------
func TestAssertFalsePositive(t *testing.T) {
	tc := testing.T{}
	AssertFalse(false, &tc, "AssertFalse")
	if tc.Failed() {
		t.Error("AssertFalse is failed")
	}
}

func TestAssertFalseNegative(t *testing.T) {
	tc := testing.T{}
	AssertFalse(true, &tc, "AssertFalse")
	if !tc.Failed() {
		t.Error("AssertFalse is not failed")
	}
}

// ----------------
// AssertEquals
// ----------------
func TestAssertEqualsPositive(t *testing.T) {
	tc := testing.T{}
	i := 1
	AssertEquals(i, i, &tc, "AssertEquals")
	if tc.Failed() {
		t.Error("AssertEquals is failed")
	}
}

func TestAssertEqualsDifferentValuesNegative(t *testing.T) {
	tc := testing.T{}
	AssertEquals(1, 2, &tc, "AssertEquals")
	if !tc.Failed() {
		t.Error("AssertFalse is not failed")
	}
}

func TestAssertEqualsDifferentTypesNegative(t *testing.T) {
	tc := testing.T{}
	i1 := int16(1)
	i2 := int32(1)
	AssertEquals(i1, i2, &tc, "AssertEquals")
	if !tc.Failed() {
		t.Error("AssertFalse is not failed")
	}
}

// ----------------
// AssertAlmostEqualsInt
// ----------------
func TestAssertAlmostEqualsIntPositive(t *testing.T) {
	checkAssertAlmostEqualsIntPositive(int(1), int(1), "int(1),int(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), int8(1), "int(1),int8(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), int16(1), "int(1),int16(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), int32(1), "int(1),int32(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), int64(1), "int(1),int64(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), uint(1), "int(1),uint(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), uint8(1), "int(1),uint8(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), uint16(1), "int(1),uint16(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), uint32(1), "int(1),uint32(1)", t)
	checkAssertAlmostEqualsIntPositive(int(1), uint64(1), "int(1),uint64(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), int(1), "uint(1),int(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), int8(1), "uint(1),int8(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), int16(1), "uint(1),int16(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), int32(1), "uint(1),int32(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), int64(1), "uint(1),int64(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), uint(1), "uint(1),uint(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), uint8(1), "uint(1),uint8(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), uint16(1), "uint(1),uint16(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), uint32(1), "uint(1),uint32(1)", t)
	checkAssertAlmostEqualsIntPositive(uint(1), uint64(1), "uint(1),uint64(1)", t)
	checkAssertAlmostEqualsIntPositive(int(-1), int(-1), "int(-1),int(-1)", t)
	checkAssertAlmostEqualsIntPositive(int(-1), int8(-1), "int(-1),int8(-1)", t)
	checkAssertAlmostEqualsIntPositive(int(-1), int16(-1), "int(-1),int16(-1)", t)
	checkAssertAlmostEqualsIntPositive(int(-1), int32(-1), "int(-1),int32(-1)", t)
	checkAssertAlmostEqualsIntPositive(int(-1), int64(-1), "int(-1),int64(-1)", t)
}

func checkAssertAlmostEqualsIntPositive[X IntNumber, Y IntNumber](expected X, actual Y, testUnitName string, t *testing.T) {
	tc := testing.T{}
	AssertAlmostEqualsInt(expected, actual, &tc, testUnitName)
	if tc.Failed() {
		t.Errorf("AssertAlmostEqualsInt is failed at %s", testUnitName)
	}
}

func TestAssertAlmostEqualsIntNegative(t *testing.T) {
	checkAssertAlmostEqualsIntNegative(int(1), int(2), "int(1),int(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), int8(2), "int(1),int8(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), int16(2), "int(1),int16(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), int32(2), "int(1),int32(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), int64(2), "int(1),int64(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), uint(2), "int(1),uint(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), uint8(2), "int(1),uint8(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), uint16(2), "int(1),uint16(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), uint32(2), "int(1),uint32(2)", t)
	checkAssertAlmostEqualsIntNegative(int(1), uint64(2), "int(1),uint64(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int(2), "uint(1),int(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int8(2), "uint(1),int8(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int16(2), "uint(1),int16(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int32(2), "uint(1),int32(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int64(2), "uint(1),int64(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), uint(2), "uint(1),uint(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), uint8(2), "uint(1),uint8(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), uint16(2), "uint(1),uint16(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), uint32(2), "uint(1),uint32(2)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), uint64(2), "uint(1),uint64(2)", t)

	checkAssertAlmostEqualsIntNegative(int(-1), uint(1), "int(-1),uint(1)", t)
	checkAssertAlmostEqualsIntNegative(int(-1), uint8(1), "int(-1),uint8(1)", t)
	checkAssertAlmostEqualsIntNegative(int(-1), uint16(1), "int(-1),uint16(1)", t)
	checkAssertAlmostEqualsIntNegative(int(-1), uint32(1), "int(-1),uint32(1)", t)
	checkAssertAlmostEqualsIntNegative(int(-1), uint64(1), "int(-1),uint64(1)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int(-1), "uint(1),int(-1)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int8(-1), "uint(1),int8(-1)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int16(-1), "uint(1),int16(-1)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int32(-1), "uint(1),int32(-1)", t)
	checkAssertAlmostEqualsIntNegative(uint(1), int64(-1), "uint(1),int64(-1)", t)
}

func checkAssertAlmostEqualsIntNegative[X IntNumber, Y IntNumber](expected X, actual Y, testUnitName string, t *testing.T) {
	tc := testing.T{}
	AssertAlmostEqualsInt(expected, actual, &tc, testUnitName)
	if !tc.Failed() {
		t.Errorf("AssertAlmostEqualsInt is not failed at %s", testUnitName)
	}
}

// ----------------
// AssertAlmostEqualsFloat
// ----------------
func TestAssertAlmostEqualsFloatPositive(t *testing.T) {
	checkAssertAlmostEqualsFloatPositive(float32(1.0), float32(1.0), 0.01, "float32(1.0),float32(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float64(1.0), float64(1.0), 0.01, "float64(1.0),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float32(1.0), float64(1.0), 0.01, "float32(1.0),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float64(1.0), float32(1.0), 0.01, "float64(1.0),float32(1.0),0.01", t)

	checkAssertAlmostEqualsFloatPositive(float32(1.001), float32(1.0), 0.01, "float32(1.001),float32(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float64(1.001), float64(1.0), 0.01, "float64(1.001),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float32(1.001), float64(1.0), 0.01, "float32(1.001),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float64(1.001), float32(1.0), 0.01, "float64(1.001),float32(1.0),0.01", t)

	checkAssertAlmostEqualsFloatPositive(float32(0.999), float32(1.0), 0.01, "float32(0.999),float32(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float64(0.999), float64(1.0), 0.01, "float64(0.999),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float32(0.999), float64(1.0), 0.01, "float32(0.999),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatPositive(float64(0.999), float32(1.0), 0.01, "float64(0.999),float32(1.0),0.01", t)
}

func checkAssertAlmostEqualsFloatPositive[X FloatNumber, Y FloatNumber](expected X, actual Y, delta float64, testUnitName string, t *testing.T) {
	tc := testing.T{}
	AssertAlmostEqualsFloat(expected, actual, delta, &tc, testUnitName)
	if tc.Failed() {
		t.Errorf("AssertAlmostEqualsFloat is failed at %s", testUnitName)
	}
}

func TestAssertAlmostEqualsFloatNegative(t *testing.T) {
	checkAssertAlmostEqualsFloatNegative(float32(1.1), float32(1.0), 0.01, "float32(1.1),float32(1.0),0.01", t)
	checkAssertAlmostEqualsFloatNegative(float64(1.1), float64(1.0), 0.01, "float64(1.1),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatNegative(float32(1.1), float64(1.0), 0.01, "float32(1.1),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatNegative(float64(1.1), float32(1.0), 0.01, "float64(1.1),float32(1.0),0.01", t)

	checkAssertAlmostEqualsFloatNegative(float32(0.9), float32(1.0), 0.01, "float32(0.9),float32(1.0),0.01", t)
	checkAssertAlmostEqualsFloatNegative(float64(0.9), float64(1.0), 0.01, "float64(0.9),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatNegative(float32(0.9), float64(1.0), 0.01, "float32(0.9),float64(1.0),0.01", t)
	checkAssertAlmostEqualsFloatNegative(float64(0.9), float32(1.0), 0.01, "float64(0.9),float32(1.0),0.01", t)
}

func checkAssertAlmostEqualsFloatNegative[X FloatNumber, Y FloatNumber](expected X, actual Y, delta float64, testUnitName string, t *testing.T) {
	tc := testing.T{}
	AssertAlmostEqualsFloat(expected, actual, delta, &tc, testUnitName)
	if !tc.Failed() {
		t.Errorf("AssertAlmostEqualsFloat is not failed at %s", testUnitName)
	}
}

// ----------------
// AssertNotAlmostEqualsFloat
// ----------------
func TestAssertNotAlmostEqualsFloatPositive(t *testing.T) {
	checkAssertNotAlmostEqualsFloatPositive(float32(1.1), float32(1.0), 0.01, "float32(1.1),float32(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatPositive(float64(1.1), float64(1.0), 0.01, "float64(1.1),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatPositive(float32(1.1), float64(1.0), 0.01, "float32(1.1),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatPositive(float64(1.1), float32(1.0), 0.01, "float64(1.1),float32(1.0),0.01", t)

	checkAssertNotAlmostEqualsFloatPositive(float32(0.9), float32(1.0), 0.01, "float32(0.9),float32(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatPositive(float64(0.9), float64(1.0), 0.01, "float64(0.9),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatPositive(float32(0.9), float64(1.0), 0.01, "float32(0.9),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatPositive(float64(0.9), float32(1.0), 0.01, "float64(0.9),float32(1.0),0.01", t)
}

func checkAssertNotAlmostEqualsFloatPositive[X FloatNumber, Y FloatNumber](expected X, actual Y, delta float64, testUnitName string, t *testing.T) {
	tc := testing.T{}
	AssertNotAlmostEqualsFloat(expected, actual, delta, &tc, testUnitName)
	if tc.Failed() {
		t.Errorf("AssertNotAlmostEqualsFloat is failed at %s", testUnitName)
	}
}

func TestAssertNotAlmostEqualsFloatNegative(t *testing.T) {
	checkAssertNotAlmostEqualsFloatNegative(float32(1.0), float32(1.0), 0.01, "float32(1.0),float32(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float64(1.0), float64(1.0), 0.01, "float64(1.0),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float32(1.0), float64(1.0), 0.01, "float32(1.0),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float64(1.0), float32(1.0), 0.01, "float64(1.0),float32(1.0),0.01", t)

	checkAssertNotAlmostEqualsFloatNegative(float32(1.001), float32(1.0), 0.01, "float32(1.001),float32(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float64(1.001), float64(1.0), 0.01, "float64(1.001),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float32(1.001), float64(1.0), 0.01, "float32(1.001),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float64(1.001), float32(1.0), 0.01, "float64(1.001),float32(1.0),0.01", t)

	checkAssertNotAlmostEqualsFloatNegative(float32(0.999), float32(1.0), 0.01, "float32(0.999),float32(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float64(0.999), float64(1.0), 0.01, "float64(0.999),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float32(0.999), float64(1.0), 0.01, "float32(0.999),float64(1.0),0.01", t)
	checkAssertNotAlmostEqualsFloatNegative(float64(0.999), float32(1.0), 0.01, "float64(0.999),float32(1.0),0.01", t)
}

func checkAssertNotAlmostEqualsFloatNegative[X FloatNumber, Y FloatNumber](expected X, actual Y, delta float64, testUnitName string, t *testing.T) {
	tc := testing.T{}
	AssertNotAlmostEqualsFloat(expected, actual, delta, &tc, testUnitName)
	if !tc.Failed() {
		t.Errorf("AssertNotAlmostEqualsFloat is not failed at %s", testUnitName)
	}
}

// ----------------
// AssertNotEquals
// ----------------
func TestAssertNotEqualsDifferentValuesPositive(t *testing.T) {
	tc := testing.T{}
	AssertNotEquals(1, 2, &tc, "AssertNotEquals")
	if tc.Failed() {
		t.Error("AssertNotEquals is failed")
	}
}

func TestAssertNotEqualsDifferentTypesPositive(t *testing.T) {
	tc := testing.T{}
	i1 := int16(1)
	i2 := int32(1)
	AssertNotEquals(i1, i2, &tc, "AssertNotEquals")
	if tc.Failed() {
		t.Error("AssertNotEquals is failed")
	}
}

func TestAssertNotEqualsNegative(t *testing.T) {
	tc := testing.T{}
	i := 1
	AssertNotEquals(i, i, &tc, "AssertEquals")
	if !tc.Failed() {
		t.Error("AssertNotEquals is not failed")
	}
}

// ----------------
// AssertAlmostEqualsInt
// ----------------
func TestAssertNotAlmostEqualsIntPositive(t *testing.T) {
	checkAssertNotAlmostEqualsIntPositive(int(1), int(2), "int(1),int(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), int8(2), "int(1),int8(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), int16(2), "int(1),int16(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), int32(2), "int(1),int32(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), int64(2), "int(1),int64(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), uint(2), "int(1),uint(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), uint8(2), "int(1),uint8(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), uint16(2), "int(1),uint16(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), uint32(2), "int(1),uint32(2)", t)
	checkAssertNotAlmostEqualsIntPositive(int(1), uint64(2), "int(1),uint64(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int(2), "uint(1),int(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int8(2), "uint(1),int8(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int16(2), "uint(1),int16(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int32(2), "uint(1),int32(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int64(2), "uint(1),int64(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), uint(2), "uint(1),uint(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), uint8(2), "uint(1),uint8(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), uint16(2), "uint(1),uint16(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), uint32(2), "uint(1),uint32(2)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), uint64(2), "uint(1),uint64(2)", t)

	checkAssertNotAlmostEqualsIntPositive(int(-1), uint(1), "int(-1),uint(1)", t)
	checkAssertNotAlmostEqualsIntPositive(int(-1), uint8(1), "int(-1),uint8(1)", t)
	checkAssertNotAlmostEqualsIntPositive(int(-1), uint16(1), "int(-1),uint16(1)", t)
	checkAssertNotAlmostEqualsIntPositive(int(-1), uint32(1), "int(-1),uint32(1)", t)
	checkAssertNotAlmostEqualsIntPositive(int(-1), uint64(1), "int(-1),uint64(1)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int(-1), "uint(1),int(-1)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int8(-1), "uint(1),int8(-1)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int16(-1), "uint(1),int16(-1)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int32(-1), "uint(1),int32(-1)", t)
	checkAssertNotAlmostEqualsIntPositive(uint(1), int64(-1), "uint(1),int64(-1)", t)
}

func checkAssertNotAlmostEqualsIntPositive[X IntNumber, Y IntNumber](expected X, actual Y, testUnitName string, t *testing.T) {
	tc := testing.T{}
	AssertNotAlmostEqualsInt(expected, actual, &tc, testUnitName)
	if tc.Failed() {
		t.Errorf("AssertNotAlmostEqualsInt is failed at %s", testUnitName)
	}
}

func TestAssertNotAlmostEqualsIntNegative(t *testing.T) {
	checkAssertNotAlmostEqualsIntNegative(int(1), int(1), "int(1),int(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), int8(1), "int(1),int8(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), int16(1), "int(1),int16(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), int32(1), "int(1),int32(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), int64(1), "int(1),int64(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), uint(1), "int(1),uint(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), uint8(1), "int(1),uint8(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), uint16(1), "int(1),uint16(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), uint32(1), "int(1),uint32(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(1), uint64(1), "int(1),uint64(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), int(1), "uint(1),int(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), int8(1), "uint(1),int8(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), int16(1), "uint(1),int16(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), int32(1), "uint(1),int32(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), int64(1), "uint(1),int64(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), uint(1), "uint(1),uint(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), uint8(1), "uint(1),uint8(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), uint16(1), "uint(1),uint16(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), uint32(1), "uint(1),uint32(1)", t)
	checkAssertNotAlmostEqualsIntNegative(uint(1), uint64(1), "uint(1),uint64(1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(-1), int(-1), "int(-1),int(-1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(-1), int8(-1), "int(-1),int8(-1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(-1), int16(-1), "int(-1),int16(-1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(-1), int32(-1), "int(-1),int32(-1)", t)
	checkAssertNotAlmostEqualsIntNegative(int(-1), int64(-1), "int(-1),int64(-1)", t)
}

func checkAssertNotAlmostEqualsIntNegative[X IntNumber, Y IntNumber](expected X, actual Y, testUnitName string, t *testing.T) {
	tc := testing.T{}
	AssertNotAlmostEqualsInt(expected, actual, &tc, testUnitName)
	if !tc.Failed() {
		t.Errorf("AssertNotAlmostEqualsInt is not failed at %s", testUnitName)
	}
}

// ----------------
// AssertHasPrefix
// ----------------
func TestHasPrefixPositive(t *testing.T) {
	tc := testing.T{}
	AssertHasPrefix("PREFIX", "PREFIXabc", &tc, "AssertHasPrefix")
	if tc.Failed() {
		t.Error("AssertHasPrefix is failed")
	}
}

func TestHasPrefixPNegative(t *testing.T) {
	tc := testing.T{}
	AssertHasPrefix("PREFIX", "PREFIabc", &tc, "AssertHasPrefix")
	if !tc.Failed() {
		t.Error("AssertHasPrefix is not failed")
	}
}

// ----------------
// AssertHasSuffix
// ----------------
func TestHasSuffixPositive(t *testing.T) {
	tc := testing.T{}
	AssertHasSuffix("SUFFIX", "abcSUFFIX", &tc, "AssertHasSuffix")
	if tc.Failed() {
		t.Error("AssertHasSuffix is failed")
	}
}

func TestHasSuffixNegative(t *testing.T) {
	tc := testing.T{}
	AssertHasSuffix("SUFFIX", "abcSUFFI", &tc, "AssertHasSuffix")
	if !tc.Failed() {
		t.Error("AssertHasSuffix is not failed")
	}
}

// ----------------
// AssertContains
// ----------------
func TestContainsPositive(t *testing.T) {
	tc := testing.T{}
	AssertContains("SUFFIX", "abcSUFFIX", &tc, "AssertContains")
	if tc.Failed() {
		t.Error("AssertContains is failed")
	}
}

func TestContainsNegative(t *testing.T) {
	tc := testing.T{}
	AssertContains("SUFFIX", "abcSUFFI", &tc, "AssertContains")
	if !tc.Failed() {
		t.Error("AssertContains is not failed")
	}
}

// ----------------
// AssertNotContains
// ----------------
func TestNotContainsPositive(t *testing.T) {
	tc := testing.T{}
	AssertNotContains("SUFFIX", "abcSUFFI", &tc, "AssertNotContains")
	if !tc.Failed() {
		t.Error("AssertNotContains is not failed")
	}
}

func TestNotContainsNegative(t *testing.T) {
	tc := testing.T{}
	AssertNotContains("SUFFIX", "abcSUFFIX", &tc, "AssertNotContains")
	if tc.Failed() {
		t.Error("AssertNotContains is failed")
	}
}
