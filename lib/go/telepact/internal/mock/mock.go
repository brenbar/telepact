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

package mock

// MockStub represents a mock stub for testing
type MockStub struct {
	Request  interface{}
	Response interface{}
}

// MockInvocation represents a mock invocation record
type MockInvocation struct {
	Request  interface{}
	Response interface{}
}

// IsSubMap checks if expected is a sub-map of actual
func IsSubMap(expected, actual map[string]interface{}) bool {
	// TODO: Implement proper sub-map checking
	for key, expectedValue := range expected {
		actualValue, exists := actual[key]
		if !exists {
			return false
		}
		
		// Recursive check for nested maps
		if expectedMap, ok := expectedValue.(map[string]interface{}); ok {
			if actualMap, ok := actualValue.(map[string]interface{}); ok {
				if !IsSubMap(expectedMap, actualMap) {
					return false
				}
			} else {
				return false
			}
		} else {
			// Direct value comparison
			if expectedValue != actualValue {
				return false
			}
		}
	}
	
	return true
}
