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
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v18/resources"
	servicespb "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v18/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newGoogleAdsFieldClientHook clientHook

// GoogleAdsFieldCallOptions contains the retry settings for each method of GoogleAdsFieldClient.
type GoogleAdsFieldCallOptions struct {
	GetGoogleAdsField     []gax.CallOption
	SearchGoogleAdsFields []gax.CallOption
}

func defaultGoogleAdsFieldGRPCClientOptions() []option.ClientOption {
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

func defaultGoogleAdsFieldCallOptions() *GoogleAdsFieldCallOptions {
	return &GoogleAdsFieldCallOptions{
		GetGoogleAdsField: []gax.CallOption{
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
		SearchGoogleAdsFields: []gax.CallOption{
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

// internalGoogleAdsFieldClient is an interface that defines the methods available from Google Ads API.
type internalGoogleAdsFieldClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetGoogleAdsField(context.Context, *servicespb.GetGoogleAdsFieldRequest, ...gax.CallOption) (*resourcespb.GoogleAdsField, error)
	SearchGoogleAdsFields(context.Context, *servicespb.SearchGoogleAdsFieldsRequest, ...gax.CallOption) *GoogleAdsFieldIterator
}

// GoogleAdsFieldClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to fetch Google Ads API fields.
type GoogleAdsFieldClient struct {
	// The internal transport-dependent client.
	internalClient internalGoogleAdsFieldClient

	// The call options for this service.
	CallOptions *GoogleAdsFieldCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GoogleAdsFieldClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GoogleAdsFieldClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *GoogleAdsFieldClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetGoogleAdsField returns just the requested field.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *GoogleAdsFieldClient) GetGoogleAdsField(ctx context.Context, req *servicespb.GetGoogleAdsFieldRequest, opts ...gax.CallOption) (*resourcespb.GoogleAdsField, error) {
	return c.internalClient.GetGoogleAdsField(ctx, req, opts...)
}

// SearchGoogleAdsFields returns all fields that match the search query.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QueryError (at )
// QuotaError (at )
// RequestError (at )
func (c *GoogleAdsFieldClient) SearchGoogleAdsFields(ctx context.Context, req *servicespb.SearchGoogleAdsFieldsRequest, opts ...gax.CallOption) *GoogleAdsFieldIterator {
	return c.internalClient.SearchGoogleAdsFields(ctx, req, opts...)
}

// googleAdsFieldGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type googleAdsFieldGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing GoogleAdsFieldClient
	CallOptions **GoogleAdsFieldCallOptions

	// The gRPC API client.
	googleAdsFieldClient servicespb.GoogleAdsFieldServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewGoogleAdsFieldClient creates a new google ads field service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to fetch Google Ads API fields.
func NewGoogleAdsFieldClient(ctx context.Context, opts ...option.ClientOption) (*GoogleAdsFieldClient, error) {
	clientOpts := defaultGoogleAdsFieldGRPCClientOptions()
	if newGoogleAdsFieldClientHook != nil {
		hookOpts, err := newGoogleAdsFieldClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := GoogleAdsFieldClient{CallOptions: defaultGoogleAdsFieldCallOptions()}

	c := &googleAdsFieldGRPCClient{
		connPool:             connPool,
		googleAdsFieldClient: servicespb.NewGoogleAdsFieldServiceClient(connPool),
		CallOptions:          &client.CallOptions,
		logger:               internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *googleAdsFieldGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *googleAdsFieldGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *googleAdsFieldGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *googleAdsFieldGRPCClient) GetGoogleAdsField(ctx context.Context, req *servicespb.GetGoogleAdsFieldRequest, opts ...gax.CallOption) (*resourcespb.GoogleAdsField, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetGoogleAdsField[0:len((*c.CallOptions).GetGoogleAdsField):len((*c.CallOptions).GetGoogleAdsField)], opts...)
	var resp *resourcespb.GoogleAdsField
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.googleAdsFieldClient.GetGoogleAdsField, req, settings.GRPC, c.logger, "GetGoogleAdsField")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *googleAdsFieldGRPCClient) SearchGoogleAdsFields(ctx context.Context, req *servicespb.SearchGoogleAdsFieldsRequest, opts ...gax.CallOption) *GoogleAdsFieldIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).SearchGoogleAdsFields[0:len((*c.CallOptions).SearchGoogleAdsFields):len((*c.CallOptions).SearchGoogleAdsFields)], opts...)
	it := &GoogleAdsFieldIterator{}
	req = proto.Clone(req).(*servicespb.SearchGoogleAdsFieldsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcespb.GoogleAdsField, string, error) {
		resp := &servicespb.SearchGoogleAdsFieldsResponse{}
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
			resp, err = executeRPC(ctx, c.googleAdsFieldClient.SearchGoogleAdsFields, req, settings.GRPC, c.logger, "SearchGoogleAdsFields")
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
