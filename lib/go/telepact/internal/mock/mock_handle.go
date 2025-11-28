package mock

import (
	"context"
	"fmt"

	"github.com/brenbar/telepact/lib/go/telepact"
	"github.com/brenbar/telepact/lib/go/telepact/internal/generation"
)

// MockHandle handles mock server requests.
// It checks stubs for matching patterns and generates responses.
type MockHandle struct {
	Schema            *telepact.MockTelepactSchema
	Stubs             []*MockStubType
	Invocations       []*MockInvocation
	RandomGenerator   *telepact.RandomGenerator
	GenerationOptions *GenerationOptions
}

// GenerationOptions configures random value generation for mocks.
type GenerationOptions struct {
	EnableMessageResponseGeneration   bool
	EnableOptionalFieldGeneration     bool
	RandomizeOptionalFieldGeneration  bool
	GeneratedCollectionLengthMin      int
	GeneratedCollectionLengthMax      int
}

// NewMockHandle creates a new MockHandle.
func NewMockHandle(schema *telepact.MockTelepactSchema, options *GenerationOptions) *MockHandle {
	if options == nil {
		options = &GenerationOptions{
			EnableMessageResponseGeneration:  true,
			EnableOptionalFieldGeneration:    true,
			RandomizeOptionalFieldGeneration: true,
			GeneratedCollectionLengthMin:     0,
			GeneratedCollectionLengthMax:     3,
		}
	}

	return &MockHandle{
		Schema:            schema,
		Stubs:             []*MockStubType{},
		Invocations:       []*MockInvocation{},
		RandomGenerator:   telepact.NewRandomGenerator(options.GeneratedCollectionLengthMin, options.GeneratedCollectionLengthMax),
		GenerationOptions: options,
	}
}

// Handle processes a mock server request.
func (h *MockHandle) Handle(ctx context.Context, request *telepact.Message) (*telepact.Message, error) {
	functionName := request.GetBodyTarget()

	requestBody := request.GetBodyPayload()

	// Check stubs for matching pattern
	for _, stub := range h.Stubs {
		if stub.Matches(functionName, requestBody) {
			responseBody := make(map[string]interface{})
			if stubResp, ok := stub.GetResponse().(map[string]interface{}); ok {
				responseBody = stubResp
			} else {
				responseBody[functionName] = stub.GetResponse()
			}
			response := telepact.NewMessage(
				map[string]interface{}{},
				responseBody,
			)
			h.recordInvocation(functionName, request, response)
			return response, nil
		}
	}

	// Generate random response if enabled
	if h.GenerationOptions.EnableMessageResponseGeneration {
		response, err := h.generateResponse(functionName)
		if err != nil {
			return nil, err
		}
		h.recordInvocation(functionName, request, response)
		return response, nil
	}

	// Return error if no stub matches and generation disabled
	return nil, fmt.Errorf("no stub found for function %s and generation disabled", functionName)
}

// AddStub adds a stub for mocking responses.
func (h *MockHandle) AddStub(stub *MockStubType) {
	h.Stubs = append(h.Stubs, stub)
}

// SetSeed sets the random seed for deterministic generation.
func (h *MockHandle) SetSeed(seed int64) {
	h.RandomGenerator = telepact.NewRandomGenerator(h.GenerationOptions.GeneratedCollectionLengthMin, h.GenerationOptions.GeneratedCollectionLengthMax)
	h.RandomGenerator.SetSeed(int32(seed))
}

// GetInvocations returns all recorded invocations.
func (h *MockHandle) GetInvocations() []*MockInvocation {
	return h.Invocations
}

// ClearInvocations clears all recorded invocations.
func (h *MockHandle) ClearInvocations() {
	h.Invocations = []*MockInvocation{}
}

// recordInvocation records a function invocation.
func (h *MockHandle) recordInvocation(functionName string, request, response *telepact.Message) {
	invocation := NewMockInvocation(functionName, request, response)
	h.Invocations = append(h.Invocations, invocation)
}

// generateResponse generates a random response for a function.
func (h *MockHandle) generateResponse(functionName string) (*telepact.Message, error) {
	// Get function result type from schema
	resultTypeName := functionName + ".->"
	resultType, exists := h.Schema.Parsed[resultTypeName]
	if !exists {
		return nil, fmt.Errorf("function result type not found: %s", resultTypeName)
	}

	// Create generation context
	ctx := &generation.GenerateContext{
		RandomGenerator:             h.RandomGenerator,
		AlwaysIncludeRequiredFields: true,
		IncludeOptionalFields:       h.GenerationOptions.EnableOptionalFieldGeneration,
		RandomizeOptionalFields:     h.GenerationOptions.RandomizeOptionalFieldGeneration,
		FnScope:                     "",
	}

	// Generate random result
	result := resultType.GenerateRandomValue(nil, false, nil, ctx)

	responseBody := make(map[string]interface{})
	if resultMap, ok := result.(map[string]interface{}); ok {
		responseBody = resultMap
	} else {
		responseBody[functionName] = result
	}

	return telepact.NewMessage(map[string]interface{}{}, responseBody), nil
}
