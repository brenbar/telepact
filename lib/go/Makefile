#|
#|  Copyright The Telepact Authors
#|
#|  Licensed under the Apache License, Version 2.0 (the "License");
#|  you may not use this file except in compliance with the License.
#|  You may obtain a copy of the License at
#|
#|  https://www.apache.org/licenses/LICENSE-2.0
#|
#|  Unless required by applicable law or agreed to in writing, software
#|  distributed under the License is distributed on an "AS IS" BASIS,
#|  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#|  See the License for the specific language governing permissions and
#|  limitations under the License.
#|

VERSION := $(shell cat ../../VERSION.txt)

.PHONY: build clean test fmt lint

build:
	go build ./...

clean:
	go clean -cache -testcache

test:
	go test -v ./...

fmt:
	go fmt ./...

lint:
	go vet ./...

tidy:
	go mod tidy

all: fmt lint build test
