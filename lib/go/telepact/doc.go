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

/*
Package telepact provides a Go implementation of the Telepact protocol.

Telepact is a multi-language API ecosystem built around a unified schema definition.
Define your API in .telepact.json files, and use this library to implement clients and servers.

# Basic Usage

Server:

	schema, err := telepact.NewTelepactSchemaFromDirectory("/path/to/schemas")
	if err != nil {
		panic(err)
	}

	handler := func(ctx context.Context, msg *telepact.Message) (*telepact.Message, error) {
		// Handle request
		return response, nil
	}

	options := &telepact.ServerOptions{
		AuthRequired: true,
		Serialization: telepact.NewDefaultSerialization(),
	}

	server, err := telepact.NewServer(schema, handler, options)
	if err != nil {
		panic(err)
	}

Client:

	adapter := func(ctx context.Context, m *telepact.Message, s *telepact.Serializer) (*telepact.Message, error) {
		// Send message via your transport
		return response, nil
	}

	options := &telepact.ClientOptions{
		UseBinary: false,
		AlwaysSendJSON: true,
	}

	client := telepact.NewClient(adapter, options)

# Transport Agnostic

The core library is transport-agnostic. It provides process (server) and adapter (client)
functions that deal with raw byte arrays, allowing you to wire them into any transport
layer (HTTP, WebSockets, gRPC, etc.).

# Status

This is a port of the Python Telepact library (lib/py) to Go. The port is in progress.
See PROGRESS.md for detailed status.
*/
package telepact
