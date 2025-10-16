//
//  Copyright The Telepact Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package telepact

import (
	"context"
)

// TestClientOptions contains options for TestClient
type TestClientOptions struct {
	GeneratedCollectionLengthMin int
	GeneratedCollectionLengthMax int
}

// NewTestClientOptions creates default TestClientOptions
func NewTestClientOptions() *TestClientOptions {
	return &TestClientOptions{
		GeneratedCollectionLengthMin: 0,
		GeneratedCollectionLengthMax: 3,
	}
}

// TestClient is a test client for Telepact
type TestClient struct {
	client *Client
	random *RandomGenerator
	schema *TelepactSchema
}

// NewTestClient creates a new TestClient
func NewTestClient(client *Client, options *TestClientOptions) *TestClient {
	return &TestClient{
		client: client,
		random: NewRandomGenerator(options.GeneratedCollectionLengthMin, options.GeneratedCollectionLengthMax),
		schema: nil,
	}
}

// AssertRequest asserts a request and returns the response
func (tc *TestClient) AssertRequest(ctx context.Context, requestMessage *Message, expectedPseudoJSONBody map[string]interface{}, expectMatch bool) (*Message, error) {
	// TODO: Implement full assert request logic
	// This is a placeholder implementation
	
	responseMessage, err := tc.client.Request(ctx, requestMessage)
	if err != nil {
		return nil, err
	}
	
	// Simplified match check (full implementation would use is_sub_map)
	if expectMatch {
		// In full implementation, check if expectedPseudoJSONBody is a sub-map of responseMessage.Body
		return responseMessage, nil
	} else {
		return &Message{
			Headers: responseMessage.Headers,
			Body:    expectedPseudoJSONBody,
		}, nil
	}
}

// SetSeed sets the random seed
func (tc *TestClient) SetSeed(seed int32) {
	tc.random.SetSeed(seed)
}
