//|
//|  Copyright The Telepact Authors
//|
//|  Licensed under the Apache License, Version 2.0 (the "License");
//|  you may not use this file except in compliance with the License.
//|  You may obtain a copy of the License at
//|
//|  https://www.apache.org/licenses/LICENSE-2.0
//|
//|  Unless required by applicable law or agreed to in writing, software
//|  distributed under the License is distributed on an "AS IS" BASIS,
//|  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//|  See the License for the specific language governing permissions and
//|  limitations under the License.
//|

package telepact

import (
	"context"
)

// TestClientOptions contains options for configuring a TestClient
type TestClientOptions struct {
	GeneratedCollectionLengthMin int
	GeneratedCollectionLengthMax int
}

// NewTestClientOptions creates default test client options
func NewTestClientOptions() *TestClientOptions {
	return &TestClientOptions{
		GeneratedCollectionLengthMin: 0,
		GeneratedCollectionLengthMax: 3,
	}
}

// TestClient is a client for testing telepact servers
type TestClient struct {
	Client *Client
	Random *RandomGenerator
	Schema *TelepactSchema
}

// NewTestClient creates a new test client
func NewTestClient(client *Client, options *TestClientOptions) *TestClient {
	return &TestClient{
		Client: client,
		Random: NewRandomGenerator(options.GeneratedCollectionLengthMin, options.GeneratedCollectionLengthMax),
		Schema: nil,
	}
}

// AssertRequest sends a request and asserts the response matches expectations
func (tc *TestClient) AssertRequest(ctx context.Context, requestMessage *Message, expectedPseudoJSONBody map[string]interface{}, expectMatch bool) (*Message, error) {
	// TODO: Implementation
	return nil, &TelepactError{Message: "Not implemented"}
}

// SetSeed sets the random seed for the test client
func (tc *TestClient) SetSeed(seed int32) {
	tc.Random.SetSeed(seed)
}
