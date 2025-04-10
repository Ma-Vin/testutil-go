[![Release](https://github.com/Ma-Vin/testutil-go/actions/workflows/go-release.yml/badge.svg?branch=release%2Fv1.0.0)](https://github.com/Ma-Vin/testutil-go/actions/workflows/go-release.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/ma-vin/testutil-go.svg)](https://pkg.go.dev/github.com/ma-vin/testutil-go)

# TestUtil

This repository provides some test utility functions for GoLang modules.

## Sonar analysis

* [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=alert_status&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)
* [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=reliability_rating&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)  [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=bugs&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)
* [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=security_rating&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)  [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=vulnerabilities&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)
* [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=sqale_rating&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)  [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=sqale_index&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)  [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=code_smells&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)
* [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=coverage&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)
* [![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=duplicated_lines_density&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)  [![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=ma-vin%3Atestutil-go&metric=ncloc&branch=release%2Fv1.0.0)](https://sonarcloud.io/summary/new_code?id=ma-vin%3Atestutil-go&branch=release%2Fv1.0.0)

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
