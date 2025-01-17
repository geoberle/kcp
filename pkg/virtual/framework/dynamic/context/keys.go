/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package context

import "context"

// apiDomainKeyContextKeyType is the type of the key for the request context value
// that will carry the API domain key.
type apiDomainKeyContextKeyType string

// apiDomainKeyContextKey is the key for the request context value
// that will carry the API domain key.
const apiDomainKeyContextKey apiDomainKeyContextKeyType = "VirtualWorkspaceAPIDomainKey"

// WithAPIDomainKey adds an API domain key to the context.
func WithAPIDomainKey(ctx context.Context, apiDomainKey string) context.Context {
	return context.WithValue(ctx, apiDomainKeyContextKey, apiDomainKey)
}

// APIDomainKeyFromContext retrieves the API domain key from the context, if any.
func APIDomainKeyFromContext(ctx context.Context) string {
	s, _ := ctx.Value(apiDomainKeyContextKey).(string)
	return s
}
