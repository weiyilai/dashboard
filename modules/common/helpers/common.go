// Copyright 2017 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"os"
	"strings"
)

// GetEnv - Lookup the environment variable provided and set to default value if variable isn't found
func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); len(value) > 0 {
		return value
	}

	return fallback
}

// GetResourceFromPath extracts the resource from the URL path /api/v1/<action>.
// Ignores potential subresources.
func GetResourceFromPath(path string) *string {
	if !strings.HasPrefix(path, "/api/v1") {
		return nil
	}

	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return nil
	}

	return &parts[3]
}
