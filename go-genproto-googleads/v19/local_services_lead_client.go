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

var newLocalServicesLeadClientHook clientHook

// LocalServicesLeadCallOptions contains the retry settings for each method of LocalServicesLeadClient.
type LocalServicesLeadCallOptions struct {
	AppendLeadConversation []gax.CallOption
}

func defaultLocalServicesLeadGRPCClientOptions() []option.ClientOption {
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

func defaultLocalServicesLeadCallOptions() *LocalServicesLeadCallOptions {
	return &LocalServicesLeadCallOptions{
		AppendLeadConversation: []gax.CallOption{
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

// internalLocalServicesLeadClient is an interface that defines the methods available from Google Ads API.
type internalLocalServicesLeadClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AppendLeadConversation(context.Context, *servicespb.AppendLeadConversationRequest, ...gax.CallOption) (*servicespb.AppendLeadConversationResponse, error)
}

// LocalServicesLeadClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service allows management of LocalServicesLead resources.
type LocalServicesLeadClient struct {
	// The internal transport-dependent client.
	internalClient internalLocalServicesLeadClient

	// The call options for this service.
	CallOptions *LocalServicesLeadCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *LocalServicesLeadClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *LocalServicesLeadClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *LocalServicesLeadClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AppendLeadConversation rPC to append Local Services Lead Conversation resources to Local Services
// Lead resources.
func (c *LocalServicesLeadClient) AppendLeadConversation(ctx context.Context, req *servicespb.AppendLeadConversationRequest, opts ...gax.CallOption) (*servicespb.AppendLeadConversationResponse, error) {
	return c.internalClient.AppendLeadConversation(ctx, req, opts...)
}

// localServicesLeadGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type localServicesLeadGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing LocalServicesLeadClient
	CallOptions **LocalServicesLeadCallOptions

	// The gRPC API client.
	localServicesLeadClient servicespb.LocalServicesLeadServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewLocalServicesLeadClient creates a new local services lead service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service allows management of LocalServicesLead resources.
func NewLocalServicesLeadClient(ctx context.Context, opts ...option.ClientOption) (*LocalServicesLeadClient, error) {
	clientOpts := defaultLocalServicesLeadGRPCClientOptions()
	if newLocalServicesLeadClientHook != nil {
		hookOpts, err := newLocalServicesLeadClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := LocalServicesLeadClient{CallOptions: defaultLocalServicesLeadCallOptions()}

	c := &localServicesLeadGRPCClient{
		connPool:                connPool,
		localServicesLeadClient: servicespb.NewLocalServicesLeadServiceClient(connPool),
		CallOptions:             &client.CallOptions,
		logger:                  internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *localServicesLeadGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *localServicesLeadGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *localServicesLeadGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *localServicesLeadGRPCClient) AppendLeadConversation(ctx context.Context, req *servicespb.AppendLeadConversationRequest, opts ...gax.CallOption) (*servicespb.AppendLeadConversationResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AppendLeadConversation[0:len((*c.CallOptions).AppendLeadConversation):len((*c.CallOptions).AppendLeadConversation)], opts...)
	var resp *servicespb.AppendLeadConversationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.localServicesLeadClient.AppendLeadConversation, req, settings.GRPC, c.logger, "AppendLeadConversation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
