//|
//|  Copyright The Telepact Authors
//|
//|  Licensed under the Apache License, Version 2.0 (the "License");
//|  you may not use this file except in compliance with the License.
//|  You may obtain a copy of the License at
//|
//|  https://www.apache.org/licenses/LICENSE-2.0
//|
//|  Unless required by applicable law or agreed to in writing, software
//|  distributed under the License is distributed on an "AS IS" BASIS,
//|  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//|  See the License for the specific language governing permissions and
//|  limitations under the License.
//|

package types

// TAny represents the 'any' type
type TAny struct{}

func (t *TAny) GetTypeParameterCount() int { return 0 }
func (t *TAny) GetName() string            { return "any" }

// TBoolean represents the 'boolean' type
type TBoolean struct{}

func (t *TBoolean) GetTypeParameterCount() int { return 0 }
func (t *TBoolean) GetName() string            { return "boolean" }

// TString represents the 'string' type
type TString struct{}

func (t *TString) GetTypeParameterCount() int { return 0 }
func (t *TString) GetName() string            { return "string" }

// TInteger represents the 'integer' type
type TInteger struct{}

func (t *TInteger) GetTypeParameterCount() int { return 0 }
func (t *TInteger) GetName() string            { return "integer" }

// TNumber represents the 'number' type
type TNumber struct{}

func (t *TNumber) GetTypeParameterCount() int { return 0 }
func (t *TNumber) GetName() string            { return "number" }

// TBytes represents the 'bytes' type
type TBytes struct{}

func (t *TBytes) GetTypeParameterCount() int { return 0 }
func (t *TBytes) GetName() string            { return "bytes" }

// TArray represents the 'array' type
type TArray struct{}

func (t *TArray) GetTypeParameterCount() int { return 1 }
func (t *TArray) GetName() string            { return "array" }

// TObject represents the 'object' type
type TObject struct{}

func (t *TObject) GetTypeParameterCount() int { return 2 }
func (t *TObject) GetName() string            { return "object" }

// TStruct represents a struct type
type TStruct struct {
	Name   string
	Fields []*TFieldDeclaration
}

func (t *TStruct) GetTypeParameterCount() int { return 0 }
func (t *TStruct) GetName() string            { return t.Name }

// TUnion represents a union type
type TUnion struct {
	Structs []*TStruct
}

func (t *TUnion) GetTypeParameterCount() int { return 0 }
func (t *TUnion) GetName() string            { return "union" }

// TSelect represents a select type
type TSelect struct {
	StructName string
	Fields     []string
}

func (t *TSelect) GetTypeParameterCount() int { return 0 }
func (t *TSelect) GetName() string            { return "select" }

// TError represents an error type
type TError struct{}

func (t *TError) GetTypeParameterCount() int { return 0 }
func (t *TError) GetName() string            { return "error" }

// THeaders represents headers type
type THeaders struct{}

func (t *THeaders) GetTypeParameterCount() int { return 0 }
func (t *THeaders) GetName() string            { return "headers" }

// TMockCall represents a mock call type
type TMockCall struct{}

func (t *TMockCall) GetTypeParameterCount() int { return 0 }
func (t *TMockCall) GetName() string            { return "mock_call" }

// TMockStub represents a mock stub type
type TMockStub struct{}

func (t *TMockStub) GetTypeParameterCount() int { return 0 }
func (t *TMockStub) GetName() string            { return "mock_stub" }
