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
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v17/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newKeywordPlanIdeaClientHook clientHook

// KeywordPlanIdeaCallOptions contains the retry settings for each method of KeywordPlanIdeaClient.
type KeywordPlanIdeaCallOptions struct {
	GenerateKeywordIdeas             []gax.CallOption
	GenerateKeywordHistoricalMetrics []gax.CallOption
	GenerateAdGroupThemes            []gax.CallOption
	GenerateKeywordForecastMetrics   []gax.CallOption
}

func defaultKeywordPlanIdeaGRPCClientOptions() []option.ClientOption {
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

func defaultKeywordPlanIdeaCallOptions() *KeywordPlanIdeaCallOptions {
	return &KeywordPlanIdeaCallOptions{
		GenerateKeywordIdeas: []gax.CallOption{
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
		GenerateKeywordHistoricalMetrics: []gax.CallOption{
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
		GenerateAdGroupThemes: []gax.CallOption{
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
		GenerateKeywordForecastMetrics: []gax.CallOption{
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

// internalKeywordPlanIdeaClient is an interface that defines the methods available from Google Ads API.
type internalKeywordPlanIdeaClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateKeywordIdeas(context.Context, *servicespb.GenerateKeywordIdeasRequest, ...gax.CallOption) *GenerateKeywordIdeaResultIterator
	GenerateKeywordHistoricalMetrics(context.Context, *servicespb.GenerateKeywordHistoricalMetricsRequest, ...gax.CallOption) (*servicespb.GenerateKeywordHistoricalMetricsResponse, error)
	GenerateAdGroupThemes(context.Context, *servicespb.GenerateAdGroupThemesRequest, ...gax.CallOption) (*servicespb.GenerateAdGroupThemesResponse, error)
	GenerateKeywordForecastMetrics(context.Context, *servicespb.GenerateKeywordForecastMetricsRequest, ...gax.CallOption) (*servicespb.GenerateKeywordForecastMetricsResponse, error)
}

// KeywordPlanIdeaClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to generate keyword ideas.
type KeywordPlanIdeaClient struct {
	// The internal transport-dependent client.
	internalClient internalKeywordPlanIdeaClient

	// The call options for this service.
	CallOptions *KeywordPlanIdeaCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeywordPlanIdeaClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeywordPlanIdeaClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *KeywordPlanIdeaClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateKeywordIdeas returns a list of keyword ideas.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanIdeaError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanIdeaClient) GenerateKeywordIdeas(ctx context.Context, req *servicespb.GenerateKeywordIdeasRequest, opts ...gax.CallOption) *GenerateKeywordIdeaResultIterator {
	return c.internalClient.GenerateKeywordIdeas(ctx, req, opts...)
}

// GenerateKeywordHistoricalMetrics returns a list of keyword historical metrics.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanIdeaClient) GenerateKeywordHistoricalMetrics(ctx context.Context, req *servicespb.GenerateKeywordHistoricalMetricsRequest, opts ...gax.CallOption) (*servicespb.GenerateKeywordHistoricalMetricsResponse, error) {
	return c.internalClient.GenerateKeywordHistoricalMetrics(ctx, req, opts...)
}

// GenerateAdGroupThemes returns a list of suggested AdGroups and suggested modifications
// (text, match type) for the given keywords.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanIdeaClient) GenerateAdGroupThemes(ctx context.Context, req *servicespb.GenerateAdGroupThemesRequest, opts ...gax.CallOption) (*servicespb.GenerateAdGroupThemesResponse, error) {
	return c.internalClient.GenerateAdGroupThemes(ctx, req, opts...)
}

// GenerateKeywordForecastMetrics returns metrics (such as impressions, clicks, total cost) of a keyword
// forecast for the given campaign.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanIdeaClient) GenerateKeywordForecastMetrics(ctx context.Context, req *servicespb.GenerateKeywordForecastMetricsRequest, opts ...gax.CallOption) (*servicespb.GenerateKeywordForecastMetricsResponse, error) {
	return c.internalClient.GenerateKeywordForecastMetrics(ctx, req, opts...)
}

// keywordPlanIdeaGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keywordPlanIdeaGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing KeywordPlanIdeaClient
	CallOptions **KeywordPlanIdeaCallOptions

	// The gRPC API client.
	keywordPlanIdeaClient servicespb.KeywordPlanIdeaServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewKeywordPlanIdeaClient creates a new keyword plan idea service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to generate keyword ideas.
func NewKeywordPlanIdeaClient(ctx context.Context, opts ...option.ClientOption) (*KeywordPlanIdeaClient, error) {
	clientOpts := defaultKeywordPlanIdeaGRPCClientOptions()
	if newKeywordPlanIdeaClientHook != nil {
		hookOpts, err := newKeywordPlanIdeaClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := KeywordPlanIdeaClient{CallOptions: defaultKeywordPlanIdeaCallOptions()}

	c := &keywordPlanIdeaGRPCClient{
		connPool:              connPool,
		keywordPlanIdeaClient: servicespb.NewKeywordPlanIdeaServiceClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *keywordPlanIdeaGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keywordPlanIdeaGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keywordPlanIdeaGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *keywordPlanIdeaGRPCClient) GenerateKeywordIdeas(ctx context.Context, req *servicespb.GenerateKeywordIdeasRequest, opts ...gax.CallOption) *GenerateKeywordIdeaResultIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateKeywordIdeas[0:len((*c.CallOptions).GenerateKeywordIdeas):len((*c.CallOptions).GenerateKeywordIdeas)], opts...)
	it := &GenerateKeywordIdeaResultIterator{}
	req = proto.Clone(req).(*servicespb.GenerateKeywordIdeasRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*servicespb.GenerateKeywordIdeaResult, string, error) {
		resp := &servicespb.GenerateKeywordIdeaResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.keywordPlanIdeaClient.GenerateKeywordIdeas(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResults(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *keywordPlanIdeaGRPCClient) GenerateKeywordHistoricalMetrics(ctx context.Context, req *servicespb.GenerateKeywordHistoricalMetricsRequest, opts ...gax.CallOption) (*servicespb.GenerateKeywordHistoricalMetricsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateKeywordHistoricalMetrics[0:len((*c.CallOptions).GenerateKeywordHistoricalMetrics):len((*c.CallOptions).GenerateKeywordHistoricalMetrics)], opts...)
	var resp *servicespb.GenerateKeywordHistoricalMetricsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanIdeaClient.GenerateKeywordHistoricalMetrics(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordPlanIdeaGRPCClient) GenerateAdGroupThemes(ctx context.Context, req *servicespb.GenerateAdGroupThemesRequest, opts ...gax.CallOption) (*servicespb.GenerateAdGroupThemesResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateAdGroupThemes[0:len((*c.CallOptions).GenerateAdGroupThemes):len((*c.CallOptions).GenerateAdGroupThemes)], opts...)
	var resp *servicespb.GenerateAdGroupThemesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanIdeaClient.GenerateAdGroupThemes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordPlanIdeaGRPCClient) GenerateKeywordForecastMetrics(ctx context.Context, req *servicespb.GenerateKeywordForecastMetricsRequest, opts ...gax.CallOption) (*servicespb.GenerateKeywordForecastMetricsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateKeywordForecastMetrics[0:len((*c.CallOptions).GenerateKeywordForecastMetrics):len((*c.CallOptions).GenerateKeywordForecastMetrics)], opts...)
	var resp *servicespb.GenerateKeywordForecastMetricsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanIdeaClient.GenerateKeywordForecastMetrics(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
