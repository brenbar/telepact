package mock

import "github.com/brenbar/telepact/lib/go/telepact"

// MockInvocation represents a recorded function invocation.
type MockInvocation struct {
	FunctionName string
	Request      *telepact.Message
	Response     *telepact.Message
}

// NewMockInvocation creates a new MockInvocation.
func NewMockInvocation(functionName string, request, response *telepact.Message) *MockInvocation {
	return &MockInvocation{
		FunctionName: functionName,
		Request:      request,
		Response:     response,
	}
}
