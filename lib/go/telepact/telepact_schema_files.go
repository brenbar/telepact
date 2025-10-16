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

package telepact

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// TelepactSchemaFiles manages telepact schema files from a directory
type TelepactSchemaFiles struct {
	FilenameToJSON map[string]string
}

// NewTelepactSchemaFiles creates a new TelepactSchemaFiles from a directory
func NewTelepactSchemaFiles(directory string) (*TelepactSchemaFiles, error) {
	filenameToJSON := make(map[string]string)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".telepact.json") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			filenameToJSON[info.Name()] = string(content)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &TelepactSchemaFiles{
		FilenameToJSON: filenameToJSON,
	}, nil
}
