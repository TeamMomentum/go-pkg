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

var newCampaignFeedClientHook clientHook

// CampaignFeedCallOptions contains the retry settings for each method of CampaignFeedClient.
type CampaignFeedCallOptions struct {
	MutateCampaignFeeds []gax.CallOption
}

func defaultCampaignFeedGRPCClientOptions() []option.ClientOption {
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

func defaultCampaignFeedCallOptions() *CampaignFeedCallOptions {
	return &CampaignFeedCallOptions{
		MutateCampaignFeeds: []gax.CallOption{
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

// internalCampaignFeedClient is an interface that defines the methods available from Google Ads API.
type internalCampaignFeedClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCampaignFeeds(context.Context, *servicespb.MutateCampaignFeedsRequest, ...gax.CallOption) (*servicespb.MutateCampaignFeedsResponse, error)
}

// CampaignFeedClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage campaign feeds.
type CampaignFeedClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignFeedClient

	// The call options for this service.
	CallOptions *CampaignFeedCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignFeedClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignFeedClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CampaignFeedClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCampaignFeeds creates, updates, or removes campaign feeds. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignFeedError (at )
// CollectionSizeError (at )
// DatabaseError (at )
// DistinctError (at )
// FieldError (at )
// FunctionError (at )
// FunctionParsingError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NotEmptyError (at )
// NullError (at )
// OperationAccessDeniedError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *CampaignFeedClient) MutateCampaignFeeds(ctx context.Context, req *servicespb.MutateCampaignFeedsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignFeedsResponse, error) {
	return c.internalClient.MutateCampaignFeeds(ctx, req, opts...)
}

// campaignFeedGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignFeedGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CampaignFeedClient
	CallOptions **CampaignFeedCallOptions

	// The gRPC API client.
	campaignFeedClient servicespb.CampaignFeedServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewCampaignFeedClient creates a new campaign feed service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage campaign feeds.
func NewCampaignFeedClient(ctx context.Context, opts ...option.ClientOption) (*CampaignFeedClient, error) {
	clientOpts := defaultCampaignFeedGRPCClientOptions()
	if newCampaignFeedClientHook != nil {
		hookOpts, err := newCampaignFeedClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CampaignFeedClient{CallOptions: defaultCampaignFeedCallOptions()}

	c := &campaignFeedGRPCClient{
		connPool:           connPool,
		campaignFeedClient: servicespb.NewCampaignFeedServiceClient(connPool),
		CallOptions:        &client.CallOptions,
		logger:             internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *campaignFeedGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignFeedGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignFeedGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignFeedGRPCClient) MutateCampaignFeeds(ctx context.Context, req *servicespb.MutateCampaignFeedsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignFeedsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCampaignFeeds[0:len((*c.CallOptions).MutateCampaignFeeds):len((*c.CallOptions).MutateCampaignFeeds)], opts...)
	var resp *servicespb.MutateCampaignFeedsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.campaignFeedClient.MutateCampaignFeeds, req, settings.GRPC, c.logger, "MutateCampaignFeeds")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
