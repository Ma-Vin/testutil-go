package testutil

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

// Determines the file path for a test case file for the caller of this function. The directory will be equal to the directory of the callers one.
// The file ending, an indicator whether to append 'scratch' or not and an indicator whether to delete all test case file or not are to pass also.
func DetermineTestCaseFilePath(testCase string, fileEnding string, isScratch bool, clean bool) string {
	_, filename, _, _ := runtime.Caller(1)
	return determineTestCaseFilePathInternal(filename, testCase, fileEnding, isScratch, clean)
}

// Equal to DetermineTestCaseFilePath, but the directory of the test files is located based on the project root directory and given sub path elements
func DetermineTestCaseFilePathAt(testCase string, fileEnding string, isScratch bool, clean bool, subPathElements ...string) string {
	_, filename, _, _ := runtime.Caller(1)
	projectRootDir := getProjectRootDir(filename)

	testCaseDir := projectRootDir
	for _, s := range subPathElements {
		testCaseDir = path.Join(testCaseDir, s)
	}
	err := os.MkdirAll(testCaseDir, 0775)
	if err != nil {
		fmt.Printf("Failed to create directory '%s' use directory of caller instead: %v", testCaseDir, err)
		fmt.Println()
		return determineTestCaseFilePathInternal(filename, testCase, fileEnding, isScratch, clean)
	}

	filePath := path.Join(testCaseDir, path.Base(filename))
	return determineTestCaseFilePathInternal(filePath, testCase, fileEnding, isScratch, clean)
}

func determineTestCaseFilePathInternal(filePath string, testCase string, fileEnding string, isScratch bool, clean bool) string {
	if clean {
		for _, filePath := range determineExistingFilePaths(testCase, filePath) {
			os.Remove(filePath)
		}
	}

	additional := ""
	if isScratch {
		additional = "_scratch"
	}

	return fmt.Sprintf("%s_%s%s.%s", filePath[:len(filePath)-3], testCase, additional, fileEnding)
}

// Determines the existing test case files of the caller of this function
func DetermineExistingTestCaseFilePaths(testCase string) []string {
	_, filename, _, _ := runtime.Caller(1)
	return determineExistingFilePaths(testCase, filename)
}

// Equal to DetermineExistingTestCaseFilePaths, but the directory of the test files is located based on the project root directory and given sub path elements
func DetermineExistingTestCaseFilePathsAt(testCase string, subPathElements ...string) []string {
	_, filename, _, _ := runtime.Caller(1)
	projectRootDir := getProjectRootDir(filename)

	testCaseDir := projectRootDir
	for _, s := range subPathElements {
		testCaseDir = path.Join(testCaseDir, s)
	}
	return determineExistingFilePathsAt(testCase, testCaseDir, path.Base(filename))
}

// Determines all existing test case files for a given path to caller file
func determineExistingFilePaths(testCase string, callerFilePath string) []string {
	return determineExistingFilePathsAt(testCase, path.Dir(callerFilePath), path.Base(callerFilePath))
}

// Determines all existing test case files at a given directory and testfile name
func determineExistingFilePathsAt(testCase string, dirPath string, baseFile string) []string {
	dirEntries, err := os.ReadDir(dirPath)
	res := make([]string, 0)
	if err == nil {
		prefixToRemove := baseFile[:len(baseFile)-3] + "_" + testCase
		for _, dirEntry := range dirEntries {
			if strings.HasPrefix(dirEntry.Name(), prefixToRemove) {
				res = append(res, path.Join(dirPath, dirEntry.Name()))
			}
		}
	}

	return res
}

// Iterates backward a given path until a go.mod file is found
func getProjectRootDir(callerPath string) string {
	dirPath := path.Dir(callerPath)
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("Failed to determine workspace root directory. Stop at '%s': %v", callerPath, err)
		fmt.Println()
		return callerPath
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.Name() == "go.mod" {
			return dirPath
		}
	}
	return getProjectRootDir(dirPath)
}
