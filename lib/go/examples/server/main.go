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
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	telepact "github.com/brenbar/telepact/lib/go"
)

func main() {
	// Create a simple schema (in production, load from files)
	schema, err := telepact.FromJSON("[]")
	if err != nil {
		log.Fatal(err)
	}

	// Create handler
	handler := func(ctx context.Context, requestMessage *telepact.Message) (*telepact.Message, error) {
		functionName := requestMessage.GetBodyTarget()
		arguments := requestMessage.GetBodyPayload()

		log.Printf("Received request: %s with args: %+v", functionName, arguments)

		switch functionName {
		case "fn.greet":
			subject := arguments["subject"].(string)
			return telepact.NewMessage(
				map[string]interface{}{},
				map[string]interface{}{
					"Ok_": map[string]interface{}{
						"message": fmt.Sprintf("Hello %s!", subject),
					},
				},
			), nil
		case "fn.add":
			x := arguments["x"].(float64)
			y := arguments["y"].(float64)
			return telepact.NewMessage(
				map[string]interface{}{},
				map[string]interface{}{
					"Ok_": map[string]interface{}{
						"result": x + y,
					},
				},
			), nil
		default:
			return nil, telepact.NewTelepactError("Function not found", nil)
		}
	}

	// Create server options
	options := telepact.NewServerOptions()
	options.AuthRequired = false

	// Create server
	server, err := telepact.NewServer(schema, handler, options)
	if err != nil {
		log.Fatal(err)
	}

	// HTTP handler
	http.HandleFunc("/api/telepact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read request body
		requestBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Process request
		response, err := server.Process(r.Context(), requestBytes, nil)
		if err != nil {
			log.Printf("Error processing request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Determine content type
		contentType := "application/json"
		if _, hasBin := response.Headers["bin_"]; hasBin {
			contentType = "application/octet-stream"
		}

		// Write response
		w.Header().Set("Content-Type", contentType)
		w.Write(response.Bytes)
	})

	log.Println("Server starting on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
