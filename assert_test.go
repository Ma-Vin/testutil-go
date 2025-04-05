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
// AssertEquals
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
