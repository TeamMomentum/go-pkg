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
	servicespb "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v19/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newTravelAssetSuggestionClientHook clientHook

// TravelAssetSuggestionCallOptions contains the retry settings for each method of TravelAssetSuggestionClient.
type TravelAssetSuggestionCallOptions struct {
	SuggestTravelAssets []gax.CallOption
}

func defaultTravelAssetSuggestionGRPCClientOptions() []option.ClientOption {
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

func defaultTravelAssetSuggestionCallOptions() *TravelAssetSuggestionCallOptions {
	return &TravelAssetSuggestionCallOptions{
		SuggestTravelAssets: []gax.CallOption{
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

// internalTravelAssetSuggestionClient is an interface that defines the methods available from Google Ads API.
type internalTravelAssetSuggestionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SuggestTravelAssets(context.Context, *servicespb.SuggestTravelAssetsRequest, ...gax.CallOption) (*servicespb.SuggestTravelAssetsResponse, error)
}

// TravelAssetSuggestionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to retrieve Travel asset suggestions.
type TravelAssetSuggestionClient struct {
	// The internal transport-dependent client.
	internalClient internalTravelAssetSuggestionClient

	// The call options for this service.
	CallOptions *TravelAssetSuggestionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TravelAssetSuggestionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TravelAssetSuggestionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TravelAssetSuggestionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SuggestTravelAssets returns Travel Asset suggestions. Asset
// suggestions are returned on a best-effort basis. There are no guarantees
// that all possible asset types will be returned for any given hotel
// property.
func (c *TravelAssetSuggestionClient) SuggestTravelAssets(ctx context.Context, req *servicespb.SuggestTravelAssetsRequest, opts ...gax.CallOption) (*servicespb.SuggestTravelAssetsResponse, error) {
	return c.internalClient.SuggestTravelAssets(ctx, req, opts...)
}

// travelAssetSuggestionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type travelAssetSuggestionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TravelAssetSuggestionClient
	CallOptions **TravelAssetSuggestionCallOptions

	// The gRPC API client.
	travelAssetSuggestionClient servicespb.TravelAssetSuggestionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewTravelAssetSuggestionClient creates a new travel asset suggestion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to retrieve Travel asset suggestions.
func NewTravelAssetSuggestionClient(ctx context.Context, opts ...option.ClientOption) (*TravelAssetSuggestionClient, error) {
	clientOpts := defaultTravelAssetSuggestionGRPCClientOptions()
	if newTravelAssetSuggestionClientHook != nil {
		hookOpts, err := newTravelAssetSuggestionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TravelAssetSuggestionClient{CallOptions: defaultTravelAssetSuggestionCallOptions()}

	c := &travelAssetSuggestionGRPCClient{
		connPool:                    connPool,
		travelAssetSuggestionClient: servicespb.NewTravelAssetSuggestionServiceClient(connPool),
		CallOptions:                 &client.CallOptions,
		logger:                      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *travelAssetSuggestionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *travelAssetSuggestionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *travelAssetSuggestionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *travelAssetSuggestionGRPCClient) SuggestTravelAssets(ctx context.Context, req *servicespb.SuggestTravelAssetsRequest, opts ...gax.CallOption) (*servicespb.SuggestTravelAssetsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SuggestTravelAssets[0:len((*c.CallOptions).SuggestTravelAssets):len((*c.CallOptions).SuggestTravelAssets)], opts...)
	var resp *servicespb.SuggestTravelAssetsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.travelAssetSuggestionClient.SuggestTravelAssets, req, settings.GRPC, c.logger, "SuggestTravelAssets")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
