package telepact

import (
	"context"
	"encoding/json"
)

// TestClientOptions contains configuration options for TestClient
type TestClientOptions struct {
	GeneratedCollectionLengthMin int
	GeneratedCollectionLengthMax int
}

// NewTestClientOptions creates default options for TestClient
func NewTestClientOptions() *TestClientOptions {
	return &TestClientOptions{
		GeneratedCollectionLengthMin: 0,
		GeneratedCollectionLengthMax: 3,
	}
}

// TestClient wraps a Client for testing purposes
type TestClient struct {
	client *Client
	random *RandomGenerator
	schema *TelepactSchema
}

// NewTestClient creates a new TestClient instance
func NewTestClient(client *Client, options *TestClientOptions) *TestClient {
	if options == nil {
		options = NewTestClientOptions()
	}

	return &TestClient{
		client: client,
		random: NewRandomGenerator(
			options.GeneratedCollectionLengthMin,
			options.GeneratedCollectionLengthMax,
		),
		schema: nil,
	}
}

// AssertRequest sends a request and asserts the response matches expectations
func (tc *TestClient) AssertRequest(
	ctx context.Context,
	requestMessage *Message,
	expectedPseudoJsonBody map[string]interface{},
	expectMatch bool,
) (*Message, error) {
	// Lazy load schema if not yet loaded
	if tc.schema == nil {
		// Request API schema
		apiRequest := &Message{
			Headers: make(map[string]interface{}),
			Body:    map[string]interface{}{"fn.api_": map[string]interface{}{}},
		}
		
		response, err := tc.client.Request(ctx, apiRequest)
		if err != nil {
			return nil, err
		}

		// Extract API schema from response
		okBody, ok := response.Body["Ok_"].(map[string]interface{})
		if !ok {
			return nil, &TelepactError{Message: "Failed to extract Ok_ from API response"}
		}

		api, ok := okBody["api"]
		if !ok {
			return nil, &TelepactError{Message: "Failed to extract api from API response"}
		}

		// Parse schema
		jsonBytes, err := json.Marshal(api)
		if err != nil {
			return nil, err
		}

		schema, err := NewTelepactSchemaFromJSON(string(jsonBytes))
		if err != nil {
			return nil, err
		}

		tc.schema = schema
	}

	// Send the actual request
	responseMessage, err := tc.client.Request(ctx, requestMessage)
	if err != nil {
		return nil, err
	}

	// Check if response matches expected
	// TODO: Implement is_sub_map logic
	didMatch := isSubMap(expectedPseudoJsonBody, responseMessage.Body)

	if expectMatch {
		if !didMatch {
			return nil, &TelepactError{
				Message: "Expected response body was not a sub map",
			}
		}
		return responseMessage, nil
	} else {
		if didMatch {
			return nil, &TelepactError{
				Message: "Expected response body was a sub map",
			}
		}

		// Generate alternative response using blueprint
		// TODO: Full implementation with GenerateContext
		return responseMessage, nil
	}
}

// SetSeed sets the random seed for deterministic generation
func (tc *TestClient) SetSeed(seed int32) {
	tc.random.SetSeed(seed)
}

// isSubMap checks if expected is a sub-map of actual
// TODO: Move to internal/mock package and implement full logic
func isSubMap(expected, actual map[string]interface{}) bool {
	for key, expectedValue := range expected {
		actualValue, ok := actual[key]
		if !ok {
			return false
		}

		// TODO: Implement deep comparison logic
		// For now, just check key existence
		_ = expectedValue
		_ = actualValue
	}
	return true
}
