// Copyright 2023 Google LLC
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
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v14/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newThirdPartyAppAnalyticsLinkClientHook clientHook

// ThirdPartyAppAnalyticsLinkCallOptions contains the retry settings for each method of ThirdPartyAppAnalyticsLinkClient.
type ThirdPartyAppAnalyticsLinkCallOptions struct {
	RegenerateShareableLinkId []gax.CallOption
}

func defaultThirdPartyAppAnalyticsLinkGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultThirdPartyAppAnalyticsLinkCallOptions() *ThirdPartyAppAnalyticsLinkCallOptions {
	return &ThirdPartyAppAnalyticsLinkCallOptions{
		RegenerateShareableLinkId: []gax.CallOption{
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

// internalThirdPartyAppAnalyticsLinkClient is an interface that defines the methods available from Google Ads API.
type internalThirdPartyAppAnalyticsLinkClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	RegenerateShareableLinkId(context.Context, *servicespb.RegenerateShareableLinkIdRequest, ...gax.CallOption) (*servicespb.RegenerateShareableLinkIdResponse, error)
}

// ThirdPartyAppAnalyticsLinkClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service allows management of links between Google Ads and third party
// app analytics.
type ThirdPartyAppAnalyticsLinkClient struct {
	// The internal transport-dependent client.
	internalClient internalThirdPartyAppAnalyticsLinkClient

	// The call options for this service.
	CallOptions *ThirdPartyAppAnalyticsLinkCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ThirdPartyAppAnalyticsLinkClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ThirdPartyAppAnalyticsLinkClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ThirdPartyAppAnalyticsLinkClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// RegenerateShareableLinkId regenerate ThirdPartyAppAnalyticsLink.shareable_link_id that should be
// provided to the third party when setting up app analytics.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ThirdPartyAppAnalyticsLinkClient) RegenerateShareableLinkId(ctx context.Context, req *servicespb.RegenerateShareableLinkIdRequest, opts ...gax.CallOption) (*servicespb.RegenerateShareableLinkIdResponse, error) {
	return c.internalClient.RegenerateShareableLinkId(ctx, req, opts...)
}

// thirdPartyAppAnalyticsLinkGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type thirdPartyAppAnalyticsLinkGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ThirdPartyAppAnalyticsLinkClient
	CallOptions **ThirdPartyAppAnalyticsLinkCallOptions

	// The gRPC API client.
	thirdPartyAppAnalyticsLinkClient servicespb.ThirdPartyAppAnalyticsLinkServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewThirdPartyAppAnalyticsLinkClient creates a new third party app analytics link service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service allows management of links between Google Ads and third party
// app analytics.
func NewThirdPartyAppAnalyticsLinkClient(ctx context.Context, opts ...option.ClientOption) (*ThirdPartyAppAnalyticsLinkClient, error) {
	clientOpts := defaultThirdPartyAppAnalyticsLinkGRPCClientOptions()
	if newThirdPartyAppAnalyticsLinkClientHook != nil {
		hookOpts, err := newThirdPartyAppAnalyticsLinkClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ThirdPartyAppAnalyticsLinkClient{CallOptions: defaultThirdPartyAppAnalyticsLinkCallOptions()}

	c := &thirdPartyAppAnalyticsLinkGRPCClient{
		connPool:                         connPool,
		thirdPartyAppAnalyticsLinkClient: servicespb.NewThirdPartyAppAnalyticsLinkServiceClient(connPool),
		CallOptions:                      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *thirdPartyAppAnalyticsLinkGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *thirdPartyAppAnalyticsLinkGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *thirdPartyAppAnalyticsLinkGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *thirdPartyAppAnalyticsLinkGRPCClient) RegenerateShareableLinkId(ctx context.Context, req *servicespb.RegenerateShareableLinkIdRequest, opts ...gax.CallOption) (*servicespb.RegenerateShareableLinkIdResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RegenerateShareableLinkId[0:len((*c.CallOptions).RegenerateShareableLinkId):len((*c.CallOptions).RegenerateShareableLinkId)], opts...)
	var resp *servicespb.RegenerateShareableLinkIdResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.thirdPartyAppAnalyticsLinkClient.RegenerateShareableLinkId(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
