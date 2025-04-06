package testutil

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestDetermineTestCaseFilePathScratch(t *testing.T) {
	testFilePath := DetermineTestCaseFilePath("a", "log", true, false)
	AssertHasSuffix(path.Join("testutil", "resources_test_a_scratch.log"), testFilePath, t, "testFilePath")
}

func TestDetermineTestCaseFilePath(t *testing.T) {
	testFilePath := DetermineTestCaseFilePath("b", "log", false, false)
	AssertHasSuffix(path.Join("testutil", "resources_test_b.log"), testFilePath, t, "testFilePath")
}

func TestDetermineTestCaseFilePathScratchClean(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	existingTestCaseFilePath := path.Join(path.Dir(filename), "resources_test_c_123_scratch.log")
	_, err := os.Create(existingTestCaseFilePath)
	AssertNil(err, t, "OpenFile err")

	testCasePaths := DetermineExistingTestCaseFilePaths("c")
	AssertEquals(1, len(testCasePaths), t, "len(testCasePaths)")
	AssertEquals(existingTestCaseFilePath, testCasePaths[0], t, "testCasePaths[0]")

	testFilePath := DetermineTestCaseFilePath("c", "log", true, true)
	_, err = os.Stat(existingTestCaseFilePath)
	AssertNotNil(err, t, "Stat err")
	AssertHasSuffix(path.Join("testutil", "resources_test_c_scratch.log"), testFilePath, t, "testFilePath")
}

func TestDetermineTestCaseFilePathClean(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	existingTestCaseFilePath := path.Join(path.Dir(filename), "resources_test_d_123.log")
	_, err := os.Create(existingTestCaseFilePath)
	AssertNil(err, t, "OpenFile err")

	testCasePaths := DetermineExistingTestCaseFilePaths("d")
	AssertEquals(1, len(testCasePaths), t, "len(testCasePaths)")
	AssertEquals(existingTestCaseFilePath, testCasePaths[0], t, "testCasePaths[0]")

	testFilePath := DetermineTestCaseFilePath("d", "log", false, true)
	_, err = os.Stat(existingTestCaseFilePath)
	AssertNotNil(err, t, "Stat err")
	AssertHasSuffix(path.Join("testutil", "resources_test_d.log"), testFilePath, t, "testFilePath")
}

func TestDetermineTestCaseFilePathAtScratch(t *testing.T) {
	testFilePath := DetermineTestCaseFilePathAt("e", "log", true, false, "testSub")
	AssertHasSuffix(path.Join("testutil", "testSub", "resources_test_e_scratch.log"), testFilePath, t, "testFilePath")
}

func TestDetermineTestCaseFilePathAt(t *testing.T) {
	testFilePath := DetermineTestCaseFilePathAt("f", "log", false, false, "testSub")
	AssertHasSuffix(path.Join("testutil", "testSub", "resources_test_f.log"), testFilePath, t, "testFilePath")
}

func TestDetermineTestCaseFilePathAtScratchClean(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	
	subDirPath := path.Join(path.Dir(filename), "testSub")
	AssertNil(os.MkdirAll(subDirPath, 0775), t, "os.MkdirAll(subDirPath, 0775)")

	existingTestCaseFilePath := path.Join(subDirPath, "resources_test_g_123_scratch.log")
	_, err := os.Create(existingTestCaseFilePath)
	AssertNil(err, t, "OpenFile err")

	testCasePaths := DetermineExistingTestCaseFilePathsAt("g", "testSub")
	AssertEquals(1, len(testCasePaths), t, "len(testCasePaths)")
	AssertEquals(existingTestCaseFilePath, testCasePaths[0], t, "testCasePaths[0]")

	testFilePath := DetermineTestCaseFilePathAt("g", "log", true, true, "testSub")
	_, err = os.Stat(existingTestCaseFilePath)
	AssertNotNil(err, t, "Stat err")
	AssertHasSuffix(path.Join("testutil", "testSub", "resources_test_g_scratch.log"), testFilePath, t, "testFilePath")
}

func TestDetermineTestCaseFilePathAtClean(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)

	subDirPath := path.Join(path.Dir(filename), "testSub")
	AssertNil(os.MkdirAll(subDirPath, 0775), t, "os.MkdirAll(subDirPath, 0775)")

	existingTestCaseFilePath := path.Join(subDirPath, "resources_test_h_123.log")
	_, err := os.Create(existingTestCaseFilePath)
	AssertNil(err, t, "OpenFile err")

	testCasePaths := DetermineExistingTestCaseFilePathsAt("h", "testSub")
	AssertEquals(1, len(testCasePaths), t, "len(testCasePaths)")
	AssertEquals(existingTestCaseFilePath, testCasePaths[0], t, "testCasePaths[0]")

	testFilePath := DetermineTestCaseFilePathAt("h", "log", false, true, "testSub")
	_, err = os.Stat(existingTestCaseFilePath)
	AssertNotNil(err, t, "Stat err")
	AssertHasSuffix(path.Join("testutil", "testSub", "resources_test_h.log"), testFilePath, t, "testFilePath")
}
