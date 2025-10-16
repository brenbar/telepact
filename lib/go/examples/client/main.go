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

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	telepact "github.com/brenbar/telepact/lib/go"
)

func main() {
	// Create adapter function
	adapter := func(ctx context.Context, message *telepact.Message, serializer *telepact.Serializer) (*telepact.Message, error) {
		// Serialize request
		requestBytes, err := serializer.Serialize(message)
		if err != nil {
			return nil, err
		}

		// Send HTTP request
		req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8000/api/telepact", bytes.NewReader(requestBytes))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Read response
		responseBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// Deserialize response
		return serializer.Deserialize(responseBytes)
	}

	// Create client
	options := telepact.NewClientOptions()
	client := telepact.NewClient(adapter, options)

	// Make greet request
	greetRequest := telepact.NewMessage(
		map[string]interface{}{},
		map[string]interface{}{
			"fn.greet": map[string]interface{}{
				"subject": "World",
			},
		},
	)

	greetResponse, err := client.Request(context.Background(), greetRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Greet response: %+v\n", greetResponse.Body)

	// Make add request
	addRequest := telepact.NewMessage(
		map[string]interface{}{},
		map[string]interface{}{
			"fn.add": map[string]interface{}{
				"x": 5.0,
				"y": 3.0,
			},
		},
	)

	addResponse, err := client.Request(context.Background(), addRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Add response: %+v\n", addResponse.Body)
}
