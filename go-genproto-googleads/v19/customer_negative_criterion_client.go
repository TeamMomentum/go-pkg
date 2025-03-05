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

var newCustomerNegativeCriterionClientHook clientHook

// CustomerNegativeCriterionCallOptions contains the retry settings for each method of CustomerNegativeCriterionClient.
type CustomerNegativeCriterionCallOptions struct {
	MutateCustomerNegativeCriteria []gax.CallOption
}

func defaultCustomerNegativeCriterionGRPCClientOptions() []option.ClientOption {
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

func defaultCustomerNegativeCriterionCallOptions() *CustomerNegativeCriterionCallOptions {
	return &CustomerNegativeCriterionCallOptions{
		MutateCustomerNegativeCriteria: []gax.CallOption{
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

// internalCustomerNegativeCriterionClient is an interface that defines the methods available from Google Ads API.
type internalCustomerNegativeCriterionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCustomerNegativeCriteria(context.Context, *servicespb.MutateCustomerNegativeCriteriaRequest, ...gax.CallOption) (*servicespb.MutateCustomerNegativeCriteriaResponse, error)
}

// CustomerNegativeCriterionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage customer negative criteria.
type CustomerNegativeCriterionClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomerNegativeCriterionClient

	// The call options for this service.
	CallOptions *CustomerNegativeCriterionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomerNegativeCriterionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomerNegativeCriterionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CustomerNegativeCriterionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCustomerNegativeCriteria creates or removes criteria. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CriterionError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomerNegativeCriterionClient) MutateCustomerNegativeCriteria(ctx context.Context, req *servicespb.MutateCustomerNegativeCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerNegativeCriteriaResponse, error) {
	return c.internalClient.MutateCustomerNegativeCriteria(ctx, req, opts...)
}

// customerNegativeCriterionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customerNegativeCriterionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CustomerNegativeCriterionClient
	CallOptions **CustomerNegativeCriterionCallOptions

	// The gRPC API client.
	customerNegativeCriterionClient servicespb.CustomerNegativeCriterionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewCustomerNegativeCriterionClient creates a new customer negative criterion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage customer negative criteria.
func NewCustomerNegativeCriterionClient(ctx context.Context, opts ...option.ClientOption) (*CustomerNegativeCriterionClient, error) {
	clientOpts := defaultCustomerNegativeCriterionGRPCClientOptions()
	if newCustomerNegativeCriterionClientHook != nil {
		hookOpts, err := newCustomerNegativeCriterionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CustomerNegativeCriterionClient{CallOptions: defaultCustomerNegativeCriterionCallOptions()}

	c := &customerNegativeCriterionGRPCClient{
		connPool:                        connPool,
		customerNegativeCriterionClient: servicespb.NewCustomerNegativeCriterionServiceClient(connPool),
		CallOptions:                     &client.CallOptions,
		logger:                          internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *customerNegativeCriterionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customerNegativeCriterionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customerNegativeCriterionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customerNegativeCriterionGRPCClient) MutateCustomerNegativeCriteria(ctx context.Context, req *servicespb.MutateCustomerNegativeCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerNegativeCriteriaResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCustomerNegativeCriteria[0:len((*c.CallOptions).MutateCustomerNegativeCriteria):len((*c.CallOptions).MutateCustomerNegativeCriteria)], opts...)
	var resp *servicespb.MutateCustomerNegativeCriteriaResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.customerNegativeCriterionClient.MutateCustomerNegativeCriteria, req, settings.GRPC, c.logger, "MutateCustomerNegativeCriteria")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
