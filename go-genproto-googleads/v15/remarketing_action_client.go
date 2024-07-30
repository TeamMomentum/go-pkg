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
	servicespb "github.com/dictav/go-genproto-googleads/pb/v15/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newRemarketingActionClientHook clientHook

// RemarketingActionCallOptions contains the retry settings for each method of RemarketingActionClient.
type RemarketingActionCallOptions struct {
	MutateRemarketingActions []gax.CallOption
}

func defaultRemarketingActionGRPCClientOptions() []option.ClientOption {
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

func defaultRemarketingActionCallOptions() *RemarketingActionCallOptions {
	return &RemarketingActionCallOptions{
		MutateRemarketingActions: []gax.CallOption{
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

// internalRemarketingActionClient is an interface that defines the methods available from Google Ads API.
type internalRemarketingActionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateRemarketingActions(context.Context, *servicespb.MutateRemarketingActionsRequest, ...gax.CallOption) (*servicespb.MutateRemarketingActionsResponse, error)
}

// RemarketingActionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage remarketing actions.
type RemarketingActionClient struct {
	// The internal transport-dependent client.
	internalClient internalRemarketingActionClient

	// The call options for this service.
	CallOptions *RemarketingActionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RemarketingActionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RemarketingActionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RemarketingActionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateRemarketingActions creates or updates remarketing actions. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ConversionActionError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *RemarketingActionClient) MutateRemarketingActions(ctx context.Context, req *servicespb.MutateRemarketingActionsRequest, opts ...gax.CallOption) (*servicespb.MutateRemarketingActionsResponse, error) {
	return c.internalClient.MutateRemarketingActions(ctx, req, opts...)
}

// remarketingActionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type remarketingActionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RemarketingActionClient
	CallOptions **RemarketingActionCallOptions

	// The gRPC API client.
	remarketingActionClient servicespb.RemarketingActionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRemarketingActionClient creates a new remarketing action service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage remarketing actions.
func NewRemarketingActionClient(ctx context.Context, opts ...option.ClientOption) (*RemarketingActionClient, error) {
	clientOpts := defaultRemarketingActionGRPCClientOptions()
	if newRemarketingActionClientHook != nil {
		hookOpts, err := newRemarketingActionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RemarketingActionClient{CallOptions: defaultRemarketingActionCallOptions()}

	c := &remarketingActionGRPCClient{
		connPool:                connPool,
		remarketingActionClient: servicespb.NewRemarketingActionServiceClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *remarketingActionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *remarketingActionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *remarketingActionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *remarketingActionGRPCClient) MutateRemarketingActions(ctx context.Context, req *servicespb.MutateRemarketingActionsRequest, opts ...gax.CallOption) (*servicespb.MutateRemarketingActionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateRemarketingActions[0:len((*c.CallOptions).MutateRemarketingActions):len((*c.CallOptions).MutateRemarketingActions)], opts...)
	var resp *servicespb.MutateRemarketingActionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.remarketingActionClient.MutateRemarketingActions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}