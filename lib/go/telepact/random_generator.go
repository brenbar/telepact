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
	"encoding/binary"
)

var words = []string{
	"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta",
	"iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pi",
	"rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega",
}

// RandomGenerator generates deterministic random values for testing
type RandomGenerator struct {
	seed                 int32
	collectionLengthMin  int
	collectionLengthMax  int
	count                int
}

// NewRandomGenerator creates a new RandomGenerator
func NewRandomGenerator(collectionLengthMin, collectionLengthMax int) *RandomGenerator {
	rg := &RandomGenerator{
		collectionLengthMin: collectionLengthMin,
		collectionLengthMax: collectionLengthMax,
		count:               0,
	}
	rg.SetSeed(0)
	return rg
}

// SetSeed sets the random seed
func (rg *RandomGenerator) SetSeed(seed int32) {
	if seed == 0 {
		rg.seed = 1
	} else {
		rg.seed = seed
	}
}

// NextInt generates the next random integer
func (rg *RandomGenerator) NextInt() int {
	x := rg.seed
	x = x ^ (x << 16)
	x = x ^ (x >> 11)
	x = x ^ (x << 5)
	
	if x == 0 {
		rg.seed = 1
	} else {
		rg.seed = x
	}
	
	rg.count++
	result := rg.seed & 0x7fffffff
	return int(result)
}

// NextIntWithCeiling generates the next random integer with a ceiling
func (rg *RandomGenerator) NextIntWithCeiling(ceiling int) int {
	if ceiling == 0 {
		return 0
	}
	return rg.NextInt() % ceiling
}

// NextBoolean generates the next random boolean
func (rg *RandomGenerator) NextBoolean() bool {
	return rg.NextIntWithCeiling(31) > 15
}

// NextBytes generates the next random bytes
func (rg *RandomGenerator) NextBytes() []byte {
	val := uint32(rg.NextInt())
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, val)
	return b
}

// NextString generates the next random string
func (rg *RandomGenerator) NextString() string {
	index := rg.NextIntWithCeiling(len(words))
	return words[index]
}

// NextDouble generates the next random double
func (rg *RandomGenerator) NextDouble() float64 {
	return float64(rg.NextInt()&0x7fffffff) / float64(0x7fffffff)
}

// NextCollectionLength generates the next random collection length
func (rg *RandomGenerator) NextCollectionLength() int {
	return rg.NextIntWithCeiling(rg.collectionLengthMax-rg.collectionLengthMin) + rg.collectionLengthMin
}
