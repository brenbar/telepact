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

// GenerateRandomArrayType generates a random array value (used internally by types package)
func GenerateRandomArrayType(blueprintValue interface{}, useBlueprintValue bool, generateElement func(interface{}, bool) interface{}, ctx *GenerateContext) []interface{} {
	if useBlueprintValue {
		if startingArray, ok := blueprintValue.([]interface{}); ok {
			array := make([]interface{}, 0, len(startingArray))
			for _, startingArrayValue := range startingArray {
				value := generateElement(startingArrayValue, true)
				array = append(array, value)
			}
			return array
		}
	}
	
	// Generate random array
	length := ctx.RandomGenerator.NextCollectionLength()
	array := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		value := generateElement(nil, false)
		array = append(array, value)
	}
	
	return array
}
