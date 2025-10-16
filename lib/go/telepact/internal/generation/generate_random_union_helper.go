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

// GenerateRandomUnionType generates a random union value (used internally by types package)
func GenerateRandomUnionType(blueprintValue interface{}, useBlueprintValue bool, tags []string, generateTag func(tag string, blueprintValue interface{}, useBlueprintValue bool) interface{}, ctx *GenerateContext) map[string]interface{} {
	if !useBlueprintValue {
		// Select random tag
		randomIndex := ctx.RandomGenerator.NextIntWithCeiling(len(tags))
		unionTag := tags[randomIndex]
		
		return map[string]interface{}{
			unionTag: generateTag(unionTag, nil, false),
		}
	}
	
	// Use blueprint value
	startingUnion, ok := blueprintValue.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}
	
	// Get the single tag and payload
	var unionTag string
	var unionStartingStruct interface{}
	for k, v := range startingUnion {
		unionTag = k
		unionStartingStruct = v
		break
	}
	
	return map[string]interface{}{
		unionTag: generateTag(unionTag, unionStartingStruct, true),
	}
}
