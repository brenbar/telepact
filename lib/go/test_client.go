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
	"math/rand"
	"time"
)

// RandomGenerator generates random test data
type RandomGenerator struct {
	rng                 *rand.Rand
	CollectionLengthMin int
	CollectionLengthMax int
}

// NewRandomGenerator creates a new RandomGenerator
func NewRandomGenerator(collectionLengthMin, collectionLengthMax int) *RandomGenerator {
	return &RandomGenerator{
		rng:                 rand.New(rand.NewSource(time.Now().UnixNano())),
		CollectionLengthMin: collectionLengthMin,
		CollectionLengthMax: collectionLengthMax,
	}
}

// SetSeed sets the random seed
func (r *RandomGenerator) SetSeed(seed int64) {
	r.rng.Seed(seed)
}

// GenerateString generates a random string
func (r *RandomGenerator) GenerateString() string {
	length := r.rng.Intn(20) + 1
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = byte(r.rng.Intn(26) + 97) // lowercase a-z
	}
	return string(bytes)
}

// GenerateInteger generates a random integer
func (r *RandomGenerator) GenerateInteger() int64 {
	return r.rng.Int63n(1000)
}

// GenerateNumber generates a random number
func (r *RandomGenerator) GenerateNumber() float64 {
	return r.rng.Float64() * 1000
}

// GenerateBoolean generates a random boolean
func (r *RandomGenerator) GenerateBoolean() bool {
	return r.rng.Intn(2) == 1
}

// GenerateCollectionLength generates a random collection length
func (r *RandomGenerator) GenerateCollectionLength() int {
	if r.CollectionLengthMax <= r.CollectionLengthMin {
		return r.CollectionLengthMin
	}
	return r.rng.Intn(r.CollectionLengthMax-r.CollectionLengthMin+1) + r.CollectionLengthMin
}

// TestClient provides testing utilities for Telepact clients
type TestClient struct {
	Client *Client
	Random *RandomGenerator
	Schema *TelepactSchema
}

// NewTestClient creates a new TestClient
func NewTestClient(client *Client) *TestClient {
	return &TestClient{
		Client: client,
		Random: NewRandomGenerator(0, 3),
		Schema: nil,
	}
}

// SetSeed sets the random seed for the test client
func (tc *TestClient) SetSeed(seed int64) {
	tc.Random.SetSeed(seed)
}

// Request makes a request using the underlying client
func (tc *TestClient) Request(ctx context.Context, message *Message) (*Message, error) {
	return tc.Client.Request(ctx, message)
}
