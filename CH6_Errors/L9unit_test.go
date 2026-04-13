package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateStatus(t *testing.T) {
	type testCase struct {
		status      string
		expectedErr error
	}

	runCases := []testCase{
		{"", errors.New("status cannot be empty")},
		{"This is a valid status update that is well within the character limit.", nil},
		{"This status update is way too long. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.", errors.New("status exceeds 140 characters")},
	}

	submitCases := append(runCases, []testCase{
		{"Another valid status.", nil},
		{"This status update, while derivative, contains exactly one hundred and forty-one characters, which is over the status update character limit.", errors.New("status exceeds 140 characters")},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		err := validateStatus(test.status)

		passed := false
		if test.expectedErr == nil {
			passed = (err == nil)
		} else {
			passed = (err != nil && err.Error() == test.expectedErr.Error())
		}

		expectedStr := "<nil>"
		if test.expectedErr != nil {
			expectedStr = test.expectedErr.Error()
		}
		actualStr := "<nil>"
		if err != nil {
			actualStr = err.Error()
		}

		if !passed {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %v
Expecting:  %v
Actual:     %v
Fail
`, test.status, expectedStr, actualStr)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %v
Expecting:  %v
Actual:     %v
Pass
`, test.status, expectedStr, actualStr)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
