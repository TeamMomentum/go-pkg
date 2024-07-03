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
	servicespb "github.com/dictav/go-genproto-googleads/pb/v13/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newAdGroupCriterionClientHook clientHook

// AdGroupCriterionCallOptions contains the retry settings for each method of AdGroupCriterionClient.
type AdGroupCriterionCallOptions struct {
	MutateAdGroupCriteria []gax.CallOption
}

func defaultAdGroupCriterionGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupCriterionCallOptions() *AdGroupCriterionCallOptions {
	return &AdGroupCriterionCallOptions{
		MutateAdGroupCriteria: []gax.CallOption{
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

// internalAdGroupCriterionClient is an interface that defines the methods available from Google Ads API.
type internalAdGroupCriterionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAdGroupCriteria(context.Context, *servicespb.MutateAdGroupCriteriaRequest, ...gax.CallOption) (*servicespb.MutateAdGroupCriteriaResponse, error)
}

// AdGroupCriterionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ad group criteria.
type AdGroupCriterionClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupCriterionClient

	// The call options for this service.
	CallOptions *AdGroupCriterionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupCriterionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupCriterionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AdGroupCriterionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAdGroupCriteria creates, updates, or removes criteria. Operation statuses are returned.
//
// List of thrown errors:
// AdGroupCriterionError (at )
// AdxError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// BiddingError (at )
// BiddingStrategyError (at )
// CollectionSizeError (at )
// ContextError (at )
// CriterionError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MultiplierError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperationAccessDeniedError (at )
// OperatorError (at )
// PolicyViolationError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *AdGroupCriterionClient) MutateAdGroupCriteria(ctx context.Context, req *servicespb.MutateAdGroupCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupCriteriaResponse, error) {
	return c.internalClient.MutateAdGroupCriteria(ctx, req, opts...)
}

// adGroupCriterionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupCriterionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AdGroupCriterionClient
	CallOptions **AdGroupCriterionCallOptions

	// The gRPC API client.
	adGroupCriterionClient servicespb.AdGroupCriterionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewAdGroupCriterionClient creates a new ad group criterion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ad group criteria.
func NewAdGroupCriterionClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupCriterionClient, error) {
	clientOpts := defaultAdGroupCriterionGRPCClientOptions()
	if newAdGroupCriterionClientHook != nil {
		hookOpts, err := newAdGroupCriterionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AdGroupCriterionClient{CallOptions: defaultAdGroupCriterionCallOptions()}

	c := &adGroupCriterionGRPCClient{
		connPool:               connPool,
		adGroupCriterionClient: servicespb.NewAdGroupCriterionServiceClient(connPool),
		CallOptions:            &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *adGroupCriterionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupCriterionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupCriterionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupCriterionGRPCClient) MutateAdGroupCriteria(ctx context.Context, req *servicespb.MutateAdGroupCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupCriteriaResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateAdGroupCriteria[0:len((*c.CallOptions).MutateAdGroupCriteria):len((*c.CallOptions).MutateAdGroupCriteria)], opts...)
	var resp *servicespb.MutateAdGroupCriteriaResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupCriterionClient.MutateAdGroupCriteria(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
