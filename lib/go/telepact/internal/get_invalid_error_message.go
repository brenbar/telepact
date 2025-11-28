// Copyright The Telepact Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"github.com/brenbar/telepact/lib/go/telepact"
	"github.com/brenbar/telepact/lib/go/telepact/internal/types"
	"github.com/brenbar/telepact/lib/go/telepact/internal/validation"
)

// GetInvalidErrorMessage creates an error message from validation failures.
func GetInvalidErrorMessage(
	errorType string,
	validationFailures []*validation.ValidationFailure,
	resultUnionType *types.TUnion,
	responseHeaders map[string]interface{},
) *telepact.Message {
	validationFailureCases := validation.MapValidationFailuresToInvalidFieldCases(validationFailures)
	
	newErrorResult := map[string]interface{}{
		errorType: map[string]interface{}{
			"cases": validationFailureCases,
		},
	}
	
	// Validate the error result internally
	ValidateResult(resultUnionType, newErrorResult)
	
	return telepact.NewMessage(responseHeaders, newErrorResult)
}
