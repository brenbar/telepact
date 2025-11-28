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

package generation

// GenerateContext provides context for random value generation
type GenerateContext struct {
	IncludeOptionalFields        bool
	RandomizeOptionalFields      bool
	AlwaysIncludeRequiredFields  bool
	FnScope                      string
	RandomGenerator              RandomGenerator
}

// RandomGenerator interface for generating random values
type RandomGenerator interface {
	NextInt() int
	NextIntWithCeiling(ceiling int) int
	NextBoolean() bool
	NextBytes() []byte
	NextString() string
	NextDouble() float64
	NextCollectionLength() int
}

// NewGenerateContext creates a new GenerateContext
func NewGenerateContext(includeOptionalFields, randomizeOptionalFields, alwaysIncludeRequiredFields bool, fnScope string, randomGenerator RandomGenerator) *GenerateContext {
	return &GenerateContext{
		IncludeOptionalFields:        includeOptionalFields,
		RandomizeOptionalFields:      randomizeOptionalFields,
		AlwaysIncludeRequiredFields:  alwaysIncludeRequiredFields,
		FnScope:                      fnScope,
		RandomGenerator:              randomGenerator,
	}
}

// Copy creates a copy of the GenerateContext with optional overrides
func (ctx *GenerateContext) Copy(includeOptionalFields, randomizeOptionalFields, alwaysIncludeRequiredFields *bool, fnScope *string, randomGenerator RandomGenerator) *GenerateContext {
	newIncludeOptionalFields := ctx.IncludeOptionalFields
	if includeOptionalFields != nil {
		newIncludeOptionalFields = *includeOptionalFields
	}
	
	newRandomizeOptionalFields := ctx.RandomizeOptionalFields
	if randomizeOptionalFields != nil {
		newRandomizeOptionalFields = *randomizeOptionalFields
	}
	
	newAlwaysIncludeRequiredFields := ctx.AlwaysIncludeRequiredFields
	if alwaysIncludeRequiredFields != nil {
		newAlwaysIncludeRequiredFields = *alwaysIncludeRequiredFields
	}
	
	newFnScope := ctx.FnScope
	if fnScope != nil {
		newFnScope = *fnScope
	}
	
	newRandomGenerator := ctx.RandomGenerator
	if randomGenerator != nil {
		newRandomGenerator = randomGenerator
	}
	
	return NewGenerateContext(newIncludeOptionalFields, newRandomizeOptionalFields, newAlwaysIncludeRequiredFields, newFnScope, newRandomGenerator)
}
