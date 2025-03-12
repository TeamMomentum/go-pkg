// Copyright 2025 Google LLC
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

package googleads

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v18/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newContentCreatorInsightsClientHook clientHook

// ContentCreatorInsightsCallOptions contains the retry settings for each method of ContentCreatorInsightsClient.
type ContentCreatorInsightsCallOptions struct {
	GenerateCreatorInsights []gax.CallOption
}

func defaultContentCreatorInsightsGRPCClientOptions() []option.ClientOption {
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

func defaultContentCreatorInsightsCallOptions() *ContentCreatorInsightsCallOptions {
	return &ContentCreatorInsightsCallOptions{
		GenerateCreatorInsights: []gax.CallOption{
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

// internalContentCreatorInsightsClient is an interface that defines the methods available from Google Ads API.
type internalContentCreatorInsightsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateCreatorInsights(context.Context, *servicespb.GenerateCreatorInsightsRequest, ...gax.CallOption) (*servicespb.GenerateCreatorInsightsResponse, error)
}

// ContentCreatorInsightsClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Content Creator Insights Service helps users find information about YouTube
// Creators and their content and how these creators and their audiences can be
// reached with Google Ads. Accessible to allowlisted customers only.
type ContentCreatorInsightsClient struct {
	// The internal transport-dependent client.
	internalClient internalContentCreatorInsightsClient

	// The call options for this service.
	CallOptions *ContentCreatorInsightsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ContentCreatorInsightsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ContentCreatorInsightsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ContentCreatorInsightsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateCreatorInsights returns insights for a collection of YouTube Creators and Channels.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *ContentCreatorInsightsClient) GenerateCreatorInsights(ctx context.Context, req *servicespb.GenerateCreatorInsightsRequest, opts ...gax.CallOption) (*servicespb.GenerateCreatorInsightsResponse, error) {
	return c.internalClient.GenerateCreatorInsights(ctx, req, opts...)
}

// contentCreatorInsightsGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type contentCreatorInsightsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ContentCreatorInsightsClient
	CallOptions **ContentCreatorInsightsCallOptions

	// The gRPC API client.
	contentCreatorInsightsClient servicespb.ContentCreatorInsightsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewContentCreatorInsightsClient creates a new content creator insights service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Content Creator Insights Service helps users find information about YouTube
// Creators and their content and how these creators and their audiences can be
// reached with Google Ads. Accessible to allowlisted customers only.
func NewContentCreatorInsightsClient(ctx context.Context, opts ...option.ClientOption) (*ContentCreatorInsightsClient, error) {
	clientOpts := defaultContentCreatorInsightsGRPCClientOptions()
	if newContentCreatorInsightsClientHook != nil {
		hookOpts, err := newContentCreatorInsightsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ContentCreatorInsightsClient{CallOptions: defaultContentCreatorInsightsCallOptions()}

	c := &contentCreatorInsightsGRPCClient{
		connPool:                     connPool,
		contentCreatorInsightsClient: servicespb.NewContentCreatorInsightsServiceClient(connPool),
		CallOptions:                  &client.CallOptions,
		logger:                       internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *contentCreatorInsightsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *contentCreatorInsightsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *contentCreatorInsightsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *contentCreatorInsightsGRPCClient) GenerateCreatorInsights(ctx context.Context, req *servicespb.GenerateCreatorInsightsRequest, opts ...gax.CallOption) (*servicespb.GenerateCreatorInsightsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateCreatorInsights[0:len((*c.CallOptions).GenerateCreatorInsights):len((*c.CallOptions).GenerateCreatorInsights)], opts...)
	var resp *servicespb.GenerateCreatorInsightsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.contentCreatorInsightsClient.GenerateCreatorInsights, req, settings.GRPC, c.logger, "GenerateCreatorInsights")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
