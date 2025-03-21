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

var newBillingSetupClientHook clientHook

// BillingSetupCallOptions contains the retry settings for each method of BillingSetupClient.
type BillingSetupCallOptions struct {
	MutateBillingSetup []gax.CallOption
}

func defaultBillingSetupGRPCClientOptions() []option.ClientOption {
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

func defaultBillingSetupCallOptions() *BillingSetupCallOptions {
	return &BillingSetupCallOptions{
		MutateBillingSetup: []gax.CallOption{
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

// internalBillingSetupClient is an interface that defines the methods available from Google Ads API.
type internalBillingSetupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateBillingSetup(context.Context, *servicespb.MutateBillingSetupRequest, ...gax.CallOption) (*servicespb.MutateBillingSetupResponse, error)
}

// BillingSetupClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for designating the business entity responsible for accrued costs.
//
// A billing setup is associated with a payments account.  Billing-related
// activity for all billing setups associated with a particular payments account
// will appear on a single invoice generated monthly.
//
// Mutates:
// The REMOVE operation cancels a pending billing setup.
// The CREATE operation creates a new billing setup.
type BillingSetupClient struct {
	// The internal transport-dependent client.
	internalClient internalBillingSetupClient

	// The call options for this service.
	CallOptions *BillingSetupCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BillingSetupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BillingSetupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BillingSetupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateBillingSetup creates a billing setup, or cancels an existing billing setup.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// BillingSetupError (at )
// DateError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *BillingSetupClient) MutateBillingSetup(ctx context.Context, req *servicespb.MutateBillingSetupRequest, opts ...gax.CallOption) (*servicespb.MutateBillingSetupResponse, error) {
	return c.internalClient.MutateBillingSetup(ctx, req, opts...)
}

// billingSetupGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type billingSetupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BillingSetupClient
	CallOptions **BillingSetupCallOptions

	// The gRPC API client.
	billingSetupClient servicespb.BillingSetupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewBillingSetupClient creates a new billing setup service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for designating the business entity responsible for accrued costs.
//
// A billing setup is associated with a payments account.  Billing-related
// activity for all billing setups associated with a particular payments account
// will appear on a single invoice generated monthly.
//
// Mutates:
// The REMOVE operation cancels a pending billing setup.
// The CREATE operation creates a new billing setup.
func NewBillingSetupClient(ctx context.Context, opts ...option.ClientOption) (*BillingSetupClient, error) {
	clientOpts := defaultBillingSetupGRPCClientOptions()
	if newBillingSetupClientHook != nil {
		hookOpts, err := newBillingSetupClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BillingSetupClient{CallOptions: defaultBillingSetupCallOptions()}

	c := &billingSetupGRPCClient{
		connPool:           connPool,
		billingSetupClient: servicespb.NewBillingSetupServiceClient(connPool),
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
func (c *billingSetupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *billingSetupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *billingSetupGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *billingSetupGRPCClient) MutateBillingSetup(ctx context.Context, req *servicespb.MutateBillingSetupRequest, opts ...gax.CallOption) (*servicespb.MutateBillingSetupResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateBillingSetup[0:len((*c.CallOptions).MutateBillingSetup):len((*c.CallOptions).MutateBillingSetup)], opts...)
	var resp *servicespb.MutateBillingSetupResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.billingSetupClient.MutateBillingSetup, req, settings.GRPC, c.logger, "MutateBillingSetup")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
