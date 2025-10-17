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

const mockCallName = "_ext.Call_"

// TMockCall represents a mock call type for testing
type TMockCall struct {
	Types map[string]TType
}

// NewTMockCall creates a new TMockCall
func NewTMockCall(types map[string]TType) *TMockCall {
	return &TMockCall{
		Types: types,
	}
}

// GetTypeParameterCount returns the number of type parameters
func (t *TMockCall) GetTypeParameterCount() int {
	return 0
}

// Validate validates a value as a mock call
func (t *TMockCall) Validate(value interface{}, typeParameters []*TTypeDeclaration, ctx *validation.ValidateContext) []*validation.ValidationFailure {
	return t.validateMockCall(value, ctx)
}

// validateMockCall validates a mock call object (internal implementation to avoid import cycles)
func (t *TMockCall) validateMockCall(givenObj interface{}, ctx *validation.ValidateContext) []*validation.ValidationFailure {
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

	functionDefCall, ok := functionDef.(*TUnion)
	if !ok {
		return []*validation.ValidationFailure{{
			Path:   []interface{}{},
			Reason: "InvalidFunctionType",
			Data: map[string]interface{}{
				"functionName": functionName,
			},
		}}
	}

	input := givenMap[functionName]

	// Validate using the function's tag struct (which has the same name as the function)
	functionDefCallTag, exists := functionDefCall.Tags[functionName]
	if !exists {
		return []*validation.ValidationFailure{{
			Path:   []interface{}{functionName},
			Reason: "FunctionTagMissing",
			Data: map[string]interface{}{
				"functionName": functionName,
			},
		}}
	}

	inputFailures := functionDefCallTag.Validate(input, nil, ctx)

	inputFailuresWithPath := make([]*validation.ValidationFailure, 0, len(inputFailures))
	for _, f := range inputFailures {
		newPath := append([]interface{}{functionName}, f.Path...)
		inputFailuresWithPath = append(inputFailuresWithPath, &validation.ValidationFailure{
			Path:   newPath,
			Reason: f.Reason,
			Data:   f.Data,
		})
	}

	// Filter out RequiredObjectKeyMissing failures (optional fields)
	result := make([]*validation.ValidationFailure, 0, len(inputFailuresWithPath))
	for _, f := range inputFailuresWithPath {
		if f.Reason != "RequiredObjectKeyMissing" {
			result = append(result, f)
		}
	}

	return result
}

// GenerateRandomValue generates a random mock call value
func (t *TMockCall) GenerateRandomValue(blueprintValue interface{}, useBlueprintValue bool, typeParameters []*TTypeDeclaration, ctx *generation.GenerateContext) interface{} {
	if useBlueprintValue && blueprintValue != nil {
		return blueprintValue
	}
	return t.generateRandomMockCall(ctx)
}

// generateRandomMockCall generates a random mock call (internal implementation to avoid import cycles)
func (t *TMockCall) generateRandomMockCall(ctx *generation.GenerateContext) interface{} {
	functionNames := make([]string, 0)
	for key := range t.Types {
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

	selectedFn, ok := t.Types[selectedFnName].(*TUnion)
	if !ok {
		return map[string]interface{}{}
	}

	alwaysIncludeReq := false
	newCtx := ctx.Copy(nil, nil, &alwaysIncludeReq, nil, ctx.RandomGenerator)

	return selectedFn.GenerateRandomValue(nil, false, nil, newCtx)
}

// GetName returns the type name
func (t *TMockCall) GetName(ctx *validation.ValidateContext) string {
	return mockCallName
}
