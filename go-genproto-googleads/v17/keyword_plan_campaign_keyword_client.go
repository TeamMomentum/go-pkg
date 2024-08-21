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

var newKeywordPlanCampaignKeywordClientHook clientHook

// KeywordPlanCampaignKeywordCallOptions contains the retry settings for each method of KeywordPlanCampaignKeywordClient.
type KeywordPlanCampaignKeywordCallOptions struct {
	MutateKeywordPlanCampaignKeywords []gax.CallOption
}

func defaultKeywordPlanCampaignKeywordGRPCClientOptions() []option.ClientOption {
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

func defaultKeywordPlanCampaignKeywordCallOptions() *KeywordPlanCampaignKeywordCallOptions {
	return &KeywordPlanCampaignKeywordCallOptions{
		MutateKeywordPlanCampaignKeywords: []gax.CallOption{
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

// internalKeywordPlanCampaignKeywordClient is an interface that defines the methods available from Google Ads API.
type internalKeywordPlanCampaignKeywordClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateKeywordPlanCampaignKeywords(context.Context, *servicespb.MutateKeywordPlanCampaignKeywordsRequest, ...gax.CallOption) (*servicespb.MutateKeywordPlanCampaignKeywordsResponse, error)
}

// KeywordPlanCampaignKeywordClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage Keyword Plan campaign keywords. KeywordPlanCampaign is
// required to add the campaign keywords. Only negative keywords are supported.
// A maximum of 1000 negative keywords are allowed per plan. This includes both
// campaign negative keywords and ad group negative keywords.
type KeywordPlanCampaignKeywordClient struct {
	// The internal transport-dependent client.
	internalClient internalKeywordPlanCampaignKeywordClient

	// The call options for this service.
	CallOptions *KeywordPlanCampaignKeywordCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeywordPlanCampaignKeywordClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeywordPlanCampaignKeywordClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *KeywordPlanCampaignKeywordClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateKeywordPlanCampaignKeywords creates, updates, or removes Keyword Plan campaign keywords. Operation
// statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanAdGroupKeywordError (at )
// KeywordPlanCampaignKeywordError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
func (c *KeywordPlanCampaignKeywordClient) MutateKeywordPlanCampaignKeywords(ctx context.Context, req *servicespb.MutateKeywordPlanCampaignKeywordsRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlanCampaignKeywordsResponse, error) {
	return c.internalClient.MutateKeywordPlanCampaignKeywords(ctx, req, opts...)
}

// keywordPlanCampaignKeywordGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keywordPlanCampaignKeywordGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing KeywordPlanCampaignKeywordClient
	CallOptions **KeywordPlanCampaignKeywordCallOptions

	// The gRPC API client.
	keywordPlanCampaignKeywordClient servicespb.KeywordPlanCampaignKeywordServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewKeywordPlanCampaignKeywordClient creates a new keyword plan campaign keyword service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage Keyword Plan campaign keywords. KeywordPlanCampaign is
// required to add the campaign keywords. Only negative keywords are supported.
// A maximum of 1000 negative keywords are allowed per plan. This includes both
// campaign negative keywords and ad group negative keywords.
func NewKeywordPlanCampaignKeywordClient(ctx context.Context, opts ...option.ClientOption) (*KeywordPlanCampaignKeywordClient, error) {
	clientOpts := defaultKeywordPlanCampaignKeywordGRPCClientOptions()
	if newKeywordPlanCampaignKeywordClientHook != nil {
		hookOpts, err := newKeywordPlanCampaignKeywordClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := KeywordPlanCampaignKeywordClient{CallOptions: defaultKeywordPlanCampaignKeywordCallOptions()}

	c := &keywordPlanCampaignKeywordGRPCClient{
		connPool:                         connPool,
		keywordPlanCampaignKeywordClient: servicespb.NewKeywordPlanCampaignKeywordServiceClient(connPool),
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
func (c *keywordPlanCampaignKeywordGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keywordPlanCampaignKeywordGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keywordPlanCampaignKeywordGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *keywordPlanCampaignKeywordGRPCClient) MutateKeywordPlanCampaignKeywords(ctx context.Context, req *servicespb.MutateKeywordPlanCampaignKeywordsRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlanCampaignKeywordsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateKeywordPlanCampaignKeywords[0:len((*c.CallOptions).MutateKeywordPlanCampaignKeywords):len((*c.CallOptions).MutateKeywordPlanCampaignKeywords)], opts...)
	var resp *servicespb.MutateKeywordPlanCampaignKeywordsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanCampaignKeywordClient.MutateKeywordPlanCampaignKeywords(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
