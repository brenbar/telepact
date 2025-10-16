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

package validation

// ValidateUnionTags validates union tags
func ValidateUnionTags(tags []string, selectedTags map[string]interface{}, actual map[string]interface{}, validateTag func(tag string, payload map[string]interface{}, selectedTags map[string]interface{}) []*ValidationFailure, ctx *ValidateContext) []*ValidationFailure {
	if len(actual) != 1 {
		return []*ValidationFailure{
			NewValidationFailure(
				[]interface{}{},
				"ObjectSizeUnexpected",
				map[string]interface{}{
					"actual":   len(actual),
					"expected": 1,
				},
			),
		}
	}
	
	// Get the single tag and payload
	var unionTarget string
	var unionPayload interface{}
	for k, v := range actual {
		unionTarget = k
		unionPayload = v
		break
	}
	
	// Check if tag exists in reference
	found := false
	for _, tag := range tags {
		if tag == unionTarget {
			found = true
			break
		}
	}
	
	if !found {
		return []*ValidationFailure{
			NewValidationFailure(
				[]interface{}{unionTarget},
				"ObjectKeyDisallowed",
				map[string]interface{}{},
			),
		}
	}
	
	// Validate payload is a map
	payloadMap, ok := unionPayload.(map[string]interface{})
	if !ok {
		return GetTypeUnexpectedValidationFailure([]interface{}{unionTarget}, unionPayload, "Object")
	}
	
	// Validate the union struct
	nestedValidationFailures := validateTag(unionTarget, payloadMap, selectedTags)
	
	// Add union target to path
	var validationFailuresWithPath []*ValidationFailure
	for _, failure := range nestedValidationFailures {
		thisPath := append([]interface{}{unionTarget}, failure.Path...)
		validationFailuresWithPath = append(validationFailuresWithPath, NewValidationFailure(thisPath, failure.Reason, failure.Data))
	}
	
	return validationFailuresWithPath
}
