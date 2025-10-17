package mock

// MockStubType represents a stub for mocking function responses.
type MockStubType struct {
	FunctionName    string
	RequestPattern  interface{}
	ResponsePattern interface{}
}

// NewMockStub creates a new MockStubType.
func NewMockStub(functionName string, requestPattern, responsePattern interface{}) *MockStubType {
	return &MockStubType{
		FunctionName:    functionName,
		RequestPattern:  requestPattern,
		ResponsePattern: responsePattern,
	}
}

// Matches checks if a request matches this stub's pattern.
func (s *MockStubType) Matches(functionName string, requestBody interface{}) bool {
	if s.FunctionName != functionName {
		return false
	}
	return IsSubMap(s.RequestPattern, requestBody)
}

// GetResponse returns the stub's response.
// In full implementation, this would support templating and dynamic responses.
func (s *MockStubType) GetResponse() interface{} {
	return s.ResponsePattern
}
