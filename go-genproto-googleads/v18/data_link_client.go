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

var newDataLinkClientHook clientHook

// DataLinkCallOptions contains the retry settings for each method of DataLinkClient.
type DataLinkCallOptions struct {
	CreateDataLink []gax.CallOption
}

func defaultDataLinkGRPCClientOptions() []option.ClientOption {
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

func defaultDataLinkCallOptions() *DataLinkCallOptions {
	return &DataLinkCallOptions{
		CreateDataLink: []gax.CallOption{
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

// internalDataLinkClient is an interface that defines the methods available from Google Ads API.
type internalDataLinkClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateDataLink(context.Context, *servicespb.CreateDataLinkRequest, ...gax.CallOption) (*servicespb.CreateDataLinkResponse, error)
}

// DataLinkClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service allows management of data links between  a Google
// Ads customer and another data entity.
type DataLinkClient struct {
	// The internal transport-dependent client.
	internalClient internalDataLinkClient

	// The call options for this service.
	CallOptions *DataLinkCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DataLinkClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DataLinkClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *DataLinkClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateDataLink creates a data link. The requesting Google Ads account name and account ID
// will be shared with the third party (such as YouTube creators for video
// links) to whom you are creating the link with. Only customers on the
// allow-list can create data links.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// DataLinkError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *DataLinkClient) CreateDataLink(ctx context.Context, req *servicespb.CreateDataLinkRequest, opts ...gax.CallOption) (*servicespb.CreateDataLinkResponse, error) {
	return c.internalClient.CreateDataLink(ctx, req, opts...)
}

// dataLinkGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type dataLinkGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing DataLinkClient
	CallOptions **DataLinkCallOptions

	// The gRPC API client.
	dataLinkClient servicespb.DataLinkServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewDataLinkClient creates a new data link service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service allows management of data links between  a Google
// Ads customer and another data entity.
func NewDataLinkClient(ctx context.Context, opts ...option.ClientOption) (*DataLinkClient, error) {
	clientOpts := defaultDataLinkGRPCClientOptions()
	if newDataLinkClientHook != nil {
		hookOpts, err := newDataLinkClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := DataLinkClient{CallOptions: defaultDataLinkCallOptions()}

	c := &dataLinkGRPCClient{
		connPool:       connPool,
		dataLinkClient: servicespb.NewDataLinkServiceClient(connPool),
		CallOptions:    &client.CallOptions,
		logger:         internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *dataLinkGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *dataLinkGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *dataLinkGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *dataLinkGRPCClient) CreateDataLink(ctx context.Context, req *servicespb.CreateDataLinkRequest, opts ...gax.CallOption) (*servicespb.CreateDataLinkResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateDataLink[0:len((*c.CallOptions).CreateDataLink):len((*c.CallOptions).CreateDataLink)], opts...)
	var resp *servicespb.CreateDataLinkResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.dataLinkClient.CreateDataLink, req, settings.GRPC, c.logger, "CreateDataLink")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
