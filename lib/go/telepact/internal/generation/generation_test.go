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

package generation_test

import (
	"testing"

	"github.com/brenbar/telepact/lib/go/telepact/internal/generation"
)

func TestNewGenerateContext(t *testing.T) {
	ctx := generation.NewGenerateContext(true, false, true, "fn.test")
	
	if ctx == nil {
		t.Fatal("Expected non-nil GenerateContext")
	}

	if !ctx.IncludeOptionalFields {
		t.Error("Expected IncludeOptionalFields to be true")
	}

	if ctx.RandomizeOptionalFieldGeneration {
		t.Error("Expected RandomizeOptionalFieldGeneration to be false")
	}

	if !ctx.AlwaysIncludeRequiredFields {
		t.Error("Expected AlwaysIncludeRequiredFields to be true")
	}

	if ctx.FunctionName != "fn.test" {
		t.Errorf("Expected FunctionName to be 'fn.test', got '%s'", ctx.FunctionName)
	}
}
