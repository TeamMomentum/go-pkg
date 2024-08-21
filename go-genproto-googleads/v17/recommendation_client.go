// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

// copy from https://github.com/dictav/go-genproto-googleads
// and changed by TeamMomentum

package googleads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v17/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newRecommendationClientHook clientHook

// RecommendationCallOptions contains the retry settings for each method of RecommendationClient.
type RecommendationCallOptions struct {
	ApplyRecommendation     []gax.CallOption
	DismissRecommendation   []gax.CallOption
	GenerateRecommendations []gax.CallOption
}

func defaultRecommendationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("googleads.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultRecommendationCallOptions() *RecommendationCallOptions {
	return &RecommendationCallOptions{
		ApplyRecommendation: []gax.CallOption{
			gax.WithTimeout(14400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DismissRecommendation: []gax.CallOption{
			gax.WithTimeout(14400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GenerateRecommendations: []gax.CallOption{
			gax.WithTimeout(14400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalRecommendationClient is an interface that defines the methods available from Google Ads API.
type internalRecommendationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ApplyRecommendation(context.Context, *servicespb.ApplyRecommendationRequest, ...gax.CallOption) (*servicespb.ApplyRecommendationResponse, error)
	DismissRecommendation(context.Context, *servicespb.DismissRecommendationRequest, ...gax.CallOption) (*servicespb.DismissRecommendationResponse, error)
	GenerateRecommendations(context.Context, *servicespb.GenerateRecommendationsRequest, ...gax.CallOption) (*servicespb.GenerateRecommendationsResponse, error)
}

// RecommendationClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage recommendations.
type RecommendationClient struct {
	// The internal transport-dependent client.
	internalClient internalRecommendationClient

	// The call options for this service.
	CallOptions *RecommendationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RecommendationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RecommendationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RecommendationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ApplyRecommendation applies given recommendations with corresponding apply parameters.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RecommendationError (at )
// RequestError (at )
// UrlFieldError (at )
func (c *RecommendationClient) ApplyRecommendation(ctx context.Context, req *servicespb.ApplyRecommendationRequest, opts ...gax.CallOption) (*servicespb.ApplyRecommendationResponse, error) {
	return c.internalClient.ApplyRecommendation(ctx, req, opts...)
}

// DismissRecommendation dismisses given recommendations.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RecommendationError (at )
// RequestError (at )
func (c *RecommendationClient) DismissRecommendation(ctx context.Context, req *servicespb.DismissRecommendationRequest, opts ...gax.CallOption) (*servicespb.DismissRecommendationResponse, error) {
	return c.internalClient.DismissRecommendation(ctx, req, opts...)
}

// GenerateRecommendations generates Recommendations based off the requested recommendation_types.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RecommendationError (at )
// RequestError (at )
func (c *RecommendationClient) GenerateRecommendations(ctx context.Context, req *servicespb.GenerateRecommendationsRequest, opts ...gax.CallOption) (*servicespb.GenerateRecommendationsResponse, error) {
	return c.internalClient.GenerateRecommendations(ctx, req, opts...)
}

// recommendationGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type recommendationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RecommendationClient
	CallOptions **RecommendationCallOptions

	// The gRPC API client.
	recommendationClient servicespb.RecommendationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRecommendationClient creates a new recommendation service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage recommendations.
func NewRecommendationClient(ctx context.Context, opts ...option.ClientOption) (*RecommendationClient, error) {
	clientOpts := defaultRecommendationGRPCClientOptions()
	if newRecommendationClientHook != nil {
		hookOpts, err := newRecommendationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RecommendationClient{CallOptions: defaultRecommendationCallOptions()}

	c := &recommendationGRPCClient{
		connPool:             connPool,
		recommendationClient: servicespb.NewRecommendationServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *recommendationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *recommendationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *recommendationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *recommendationGRPCClient) ApplyRecommendation(ctx context.Context, req *servicespb.ApplyRecommendationRequest, opts ...gax.CallOption) (*servicespb.ApplyRecommendationResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ApplyRecommendation[0:len((*c.CallOptions).ApplyRecommendation):len((*c.CallOptions).ApplyRecommendation)], opts...)
	var resp *servicespb.ApplyRecommendationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.recommendationClient.ApplyRecommendation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *recommendationGRPCClient) DismissRecommendation(ctx context.Context, req *servicespb.DismissRecommendationRequest, opts ...gax.CallOption) (*servicespb.DismissRecommendationResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DismissRecommendation[0:len((*c.CallOptions).DismissRecommendation):len((*c.CallOptions).DismissRecommendation)], opts...)
	var resp *servicespb.DismissRecommendationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.recommendationClient.DismissRecommendation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *recommendationGRPCClient) GenerateRecommendations(ctx context.Context, req *servicespb.GenerateRecommendationsRequest, opts ...gax.CallOption) (*servicespb.GenerateRecommendationsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateRecommendations[0:len((*c.CallOptions).GenerateRecommendations):len((*c.CallOptions).GenerateRecommendations)], opts...)
	var resp *servicespb.GenerateRecommendationsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.recommendationClient.GenerateRecommendations(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
