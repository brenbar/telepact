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
	"encoding/binary"
)

// RandomGenerator generates random values for testing
type RandomGenerator struct {
	Seed                 int32
	CollectionLengthMin  int
	CollectionLengthMax  int
	Count                int
}

var words = []string{
	"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta",
	"iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pi",
	"rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega",
}

// NewRandomGenerator creates a new RandomGenerator
func NewRandomGenerator(collectionLengthMin, collectionLengthMax int) *RandomGenerator {
	rg := &RandomGenerator{
		CollectionLengthMin: collectionLengthMin,
		CollectionLengthMax: collectionLengthMax,
		Count:               0,
	}
	rg.SetSeed(0)
	return rg
}

// SetSeed sets the random seed
func (r *RandomGenerator) SetSeed(seed int32) {
	if seed == 0 {
		r.Seed = 1
	} else {
		r.Seed = seed
	}
}

// NextInt generates the next random integer
func (r *RandomGenerator) NextInt() int {
	x := r.Seed
	x = x ^ (x << 16)
	x = x ^ (x >> 11)
	x = x ^ (x << 5)
	if x == 0 {
		r.Seed = 1
	} else {
		r.Seed = x
	}
	r.Count++
	result := r.Seed & 0x7fffffff
	return int(result)
}

// NextIntWithCeiling generates a random integer with a ceiling
func (r *RandomGenerator) NextIntWithCeiling(ceiling int) int {
	if ceiling == 0 {
		return 0
	}
	return r.NextInt() % ceiling
}

// NextBoolean generates a random boolean
func (r *RandomGenerator) NextBoolean() bool {
	return r.NextIntWithCeiling(31) > 15
}

// NextBytes generates random bytes
func (r *RandomGenerator) NextBytes() []byte {
	val := r.NextInt()
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(val))
	return bytes
}

// NextString generates a random string
func (r *RandomGenerator) NextString() string {
	index := r.NextIntWithCeiling(len(words))
	return words[index]
}

// NextDouble generates a random double
func (r *RandomGenerator) NextDouble() float64 {
	return float64(r.NextInt()&0x7fffffff) / float64(0x7fffffff)
}

// NextCollectionLength generates a random collection length
func (r *RandomGenerator) NextCollectionLength() int {
	return r.NextIntWithCeiling(r.CollectionLengthMax-r.CollectionLengthMin) + r.CollectionLengthMin
}
