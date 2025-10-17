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

package types

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/validation"
	"github.com/brenbar/telepact/lib/go/telepact/internal/generation"
)

const mockStubName = "_ext.Stub_"

// TMockStub represents a mock stub type for testing
type TMockStub struct {
	Types map[string]TType
}

// NewTMockStub creates a new TMockStub
func NewTMockStub(types map[string]TType) *TMockStub {
	return &TMockStub{
		Types: types,
	}
}

// GetTypeParameterCount returns the number of type parameters
func (t *TMockStub) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a mock stub
func (t *TMockStub) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	return t.validateMockStub(value, ctx)
}

// validateMockStub validates a mock stub object (internal implementation to avoid import cycles)
func (t *TMockStub) validateMockStub(givenObj interface{}, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	givenMap, ok := givenObj.(map[string]interface{})
	if !ok {
		return validation.GetTypeUnexpectedValidationFailure([]interface{}{}, givenObj, "Object")
	}

	regexString := `^fn\..*$`

	keys := make([]string, 0, len(givenMap))
	for k := range givenMap {
		keys = append(keys, k)
	}

	matches := []string{}
	for _, k := range keys {
		if len(k) >= 3 && k[:3] == "fn." {
			matches = append(matches, k)
		}
	}

	if len(matches) != 1 {
		return []*validation.ValidationFailure{{
			Path:   []interface{}{},
			Reason: "ObjectKeyRegexMatchCountUnexpected",
			Data: map[string]interface{}{
				"regex":    regexString,
				"actual":   len(matches),
				"expected": 1,
				"keys":     keys,
			},
		}}
	}

	functionName := matches[0]
	functionDef, exists := t.Types[functionName]
	if !exists {
		return []*validation.ValidationFailure{{
			Path:   []interface{}{},
			Reason: "UnknownFunctionName",
			Data: map[string]interface{}{
				"functionName": functionName,
			},
		}}
	}

	functionDefResult, ok := functionDef.(*TUnion)
	if !ok {
		return []*validation.ValidationFailure{{
			Path:   []interface{}{},
			Reason: "InvalidFunctionType",
			Data: map[string]interface{}{
				"functionName": functionName,
			},
		}}
	}

	result := givenMap[functionName]

	// Validate using the result union
	resultFailures := functionDefResult.Validate(result, []interface{}{}, ctx)

	resultFailuresWithPath := make([]*validation.ValidationFailure, 0, len(resultFailures))
	for _, f := range resultFailures {
		newPath := append([]interface{}{functionName}, f.Path...)
		resultFailuresWithPath = append(resultFailuresWithPath, &validation.ValidationFailure{
			Path:   newPath,
			Reason: f.Reason,
			Data:   f.Data,
		})
	}

	return resultFailuresWithPath
}

// GenerateRandomValue generates a random mock stub value
func (t *TMockStub) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	if useBlueprintValue && blueprintValue != nil {
		return blueprintValue
	}
	return t.generateRandomMockStub(ctx)
}

// generateRandomMockStub generates a random mock stub (internal implementation to avoid import cycles)
func (t *TMockStub) generateRandomMockStub(ctx *generation.GenerateContext) interface{} {
	functionNames := make([]string, 0)
	for key := range t.Types {
		// Match pattern: fn.* but not ending with .-> or _
		if len(key) >= 3 && key[:3] == "fn." &&
			!(len(key) >= 3 && key[len(key)-3:] == ".->") &&
			!(len(key) >= 1 && key[len(key)-1:] == "_") {
			functionNames = append(functionNames, key)
		}
	}

	if len(functionNames) == 0 {
		return map[string]interface{}{}
	}

	selectedFnName := functionNames[ctx.RandomGenerator.NextIntWithCeiling(len(functionNames))]

	// Look for the result type (fn.Name.->)
	resultTypeName := selectedFnName + ".->"
	resultType, exists := t.Types[resultTypeName]
	if !exists {
		return map[string]interface{}{}
	}

	resultUnion, ok := resultType.(*TUnion)
	if !ok {
		return map[string]interface{}{}
	}

	return resultUnion.GenerateRandomValue(nil, false, []*TTypeDeclaration{}, ctx)
}

// GetName returns the type name
func (t *TMockStub) GetName(ctx *validation.ValidateContext) string {
	return mockStubName
}
