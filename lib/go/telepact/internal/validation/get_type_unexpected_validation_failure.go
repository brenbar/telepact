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

import (
	"github.com/brenbar/telepact/lib/go/telepact/internal/util"
)

// GetTypeUnexpectedValidationFailure creates a validation failure for type mismatches
func GetTypeUnexpectedValidationFailure(path []interface{}, value interface{}, expectedType string) []*ValidationFailure {
	actualType := util.GetType(value)
	data := map[string]interface{}{
		"actual": map[string]interface{}{
			actualType: map[string]interface{}{},
		},
		"expected": map[string]interface{}{
			expectedType: map[string]interface{}{},
		},
	}
	return []*ValidationFailure{
		NewValidationFailure(path, "TypeUnexpected", data),
	}
}
