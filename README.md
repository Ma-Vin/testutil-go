[![Build and test](https://github.com/Ma-Vin/testutil/actions/workflows/go-build.yaml/badge.svg)](https://github.com/Ma-Vin/testutil/actions/workflows/go-build.yaml)

# TestUtil

This repository provides some test utility functions for GoLang modules.

## Asserts

There are some assert functions to check the correctness of an existing result. If a check is not successful the `Errorf` function of the go `testing.T` framework will be called. To distinguish different fails, a name of the testing object is required as parameter and will be added to the fail-message.

The following checks are available

* `AssertNotNil`
* `AssertNil`
* `AssertTrue`
* `AssertFalse`
* `AssertEquals`
* `AssertAlmostEqualsInt`: Compares `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32` and `uint64`. E.g. the following assert will no fail

``` Go
 AssertAlmostEqualsInt(int8(1), uint16(1), t, "int8 vs uint16")
```

* `AssertAlmostEqualsFloat`: Compares `float32` and `float64` with a given delta up to which the two values are considered equal.
* `AssertNotEquals`
* `AssertNotAlmostEqualsInt`: analogous to `AssertAlmostEqualsInt`
* `AssertNotAlmostEqualsFloat`: analogous to `AssertAlmostEqualsFloat`
* `AssertHasPrefix`
* `AssertHasSuffix`

## Resources

There are some functions to determine existing test resources or a name for new ones.

The following functions are available

* `DetermineTestCaseFilePath`: returns a full qualified path within the directory of the caller. The name begins with the filename of the go file appended by the `testCase`, optional `scratch` (to add `**/*scratch*` at `.gitignore`) and a file ending instead of the callers `go` ending (An underscore will be added between each element). If `clean` is passed with `true` all files found by `DetermineExistingTestCaseFilePaths` will be deleted.
* `DetermineTestCaseFilePathAt`: analogous to `DetermineTestCaseFilePath` but the directory starts at the directory of `go.mod` and given sub folders by `subPathElements` parameter
* `DetermineExistingTestCaseFilePaths`: Returns a slice of paths to file which have a prefix caller name appended by the testcase. E.g. Caller `/some/project/caller.go` would find files like `/some/project/caller_test123.anything` if the testcase is `test123`
* `DetermineExistingTestCaseFilePathsAt` analogous to `DetermineExistingTestCaseFilePaths` but the directory starts at the directory of `go.mod` and given sub folders by `subPathElements` parameter
