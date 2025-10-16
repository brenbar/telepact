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

// GenerateRandomObjectType generates a random object (map) value (used internally by types package)
func GenerateRandomObjectType(blueprintValue interface{}, useBlueprintValue bool, generateValue func(interface{}, bool) interface{}, ctx *GenerateContext) map[string]interface{} {
	if useBlueprintValue {
		if startingObject, ok := blueprintValue.(map[string]interface{}); ok {
			object := make(map[string]interface{})
			for k, v := range startingObject {
				value := generateValue(v, true)
				object[k] = value
			}
			return object
		}
	}
	
	// Generate random object
	length := ctx.RandomGenerator.NextCollectionLength()
	object := make(map[string]interface{})
	for i := 0; i < length; i++ {
		key := ctx.RandomGenerator.NextString()
		value := generateValue(nil, false)
		object[key] = value
	}
	
	return object
}
