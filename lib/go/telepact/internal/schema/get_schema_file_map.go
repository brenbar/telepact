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

package schema

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// GetSchemaFileMap reads all .telepact.json files from a directory
func GetSchemaFileMap(directory string) (map[string]string, error) {
	fileMap := make(map[string]string)
	
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".telepact.json") {
			filePath := filepath.Join(directory, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			fileMap[file.Name()] = string(content)
		}
	}
	
	return fileMap, nil
}
