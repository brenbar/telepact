package mock

import (
	"fmt"
	"strings"
)

// Verify checks that a specific invocation occurred.
// Returns an error if the invocation is not found.
func Verify(handle *MockHandle, functionName string, requestPattern interface{}) error {
	for _, inv := range handle.Invocations {
		if inv.FunctionName == functionName {
			requestBody := inv.Request.GetBodyPayload()
			if IsSubMap(requestPattern, requestBody) {
				return nil
			}
		}
	}

	// Build error message with invocation details
	var invocations []string
	for _, inv := range handle.Invocations {
		invocations = append(invocations, fmt.Sprintf("  - %s", inv.FunctionName))
	}

	if len(invocations) == 0 {
		return fmt.Errorf("verification failed: no invocations recorded")
	}

	return fmt.Errorf("verification failed: expected invocation of %s not found\nRecorded invocations:\n%s",
		functionName, strings.Join(invocations, "\n"))
}

// VerifyNoMoreInteractions checks that no unexpected invocations occurred.
// Returns an error if there are unverified invocations.
func VerifyNoMoreInteractions(handle *MockHandle) error {
	if len(handle.Invocations) > 0 {
		var invocations []string
		for _, inv := range handle.Invocations {
			invocations = append(invocations, fmt.Sprintf("  - %s", inv.FunctionName))
		}
		return fmt.Errorf("unexpected invocations found:\n%s", strings.Join(invocations, "\n"))
	}
	return nil
}

// VerifyInvocationCount checks that a function was invoked exactly n times.
func VerifyInvocationCount(handle *MockHandle, functionName string, expectedCount int) error {
	actualCount := 0
	for _, inv := range handle.Invocations {
		if inv.FunctionName == functionName {
			actualCount++
		}
	}

	if actualCount != expectedCount {
		return fmt.Errorf("verification failed: expected %d invocations of %s, got %d",
			expectedCount, functionName, actualCount)
	}
	return nil
}
