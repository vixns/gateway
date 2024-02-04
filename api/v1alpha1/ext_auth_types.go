// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

import (
	gwapiv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// +kubebuilder:validation:XValidation:message="one of grpc or http must be specified",rule="(has(self.grpc) || has(self.http))"
// +kubebuilder:validation:XValidation:message="only one of grpc or http can be specified",rule="(has(self.grpc) && !has(self.http)) || (!has(self.grpc) && has(self.http))"
//
// ExtAuth defines the configuration for External Authorization.
type ExtAuth struct {
	// GRPC defines the gRPC External Authorization service.
	// Either GRPCService or HTTPService must be specified,
	// and only one of them can be provided.
	GRPC *GRPCExtAuthService `json:"grpc,omitempty"`

	// HTTP defines the HTTP External Authorization service.
	// Either GRPCService or HTTPService must be specified,
	// and only one of them can be provided.
	HTTP *HTTPExtAuthService `json:"http,omitempty"`

	// HeadersToExtAuth defines the client request headers that will be included
	// in the request to the external authorization service.
	// Note: If not specified, the default behavior for gRPC and HTTP external
	// authorization services is different due to backward compatibility reasons.
	// All headers will be included in the check request to a gRPC authorization server.
	// Only the following headers will be included in the check request to an HTTP
	// authorization server: Host, Method, Path, Content-Length, and Authorization.
	// And these headers will always be included to the check request to an HTTP
	// authorization server by default, no matter whether they are specified
	// in HeadersToExtAuth or not.
	// +optional
	HeadersToExtAuth []string `json:"headersToExtAuth,omitempty"`
}

// GRPCExtAuthService defines the gRPC External Authorization service
// The authorization request message is defined in
// https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/auth/v3/external_auth.proto
type GRPCExtAuthService struct {
	// BackendObjectReference references a Kubernetes object that represents the
	// backend server to which the authorization request will be sent.
	// Only service Kind is supported for now.
	gwapiv1.BackendObjectReference `json:",inline"`
}

// HTTPExtAuthService defines the HTTP External Authorization service
type HTTPExtAuthService struct {
	// BackendObjectReference references a Kubernetes object that represents the
	// backend server to which the authorization request will be sent.
	// Only service Kind is supported for now.
	gwapiv1.BackendObjectReference `json:",inline"`

	// Path is the path of the HTTP External Authorization service.
	// If path is specified, the authorization request will be sent to that path,
	// or else the authorization request will be sent to the root path.
	Path *string `json:"path,omitempty"`

	// HeadersToBackend are the authorization response headers that will be added
	// to the original client request before sending it to the backend server.
	// Note that coexisting headers will be overridden.
	// If not specified, no authorization response headers will be added to the
	// original client request.
	// +optional
	HeadersToBackend []string `json:"headersToBackend,omitempty"`
}
