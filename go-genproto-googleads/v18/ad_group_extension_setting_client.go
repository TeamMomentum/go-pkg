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

var newAdGroupExtensionSettingClientHook clientHook

// AdGroupExtensionSettingCallOptions contains the retry settings for each method of AdGroupExtensionSettingClient.
type AdGroupExtensionSettingCallOptions struct {
	MutateAdGroupExtensionSettings []gax.CallOption
}

func defaultAdGroupExtensionSettingGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupExtensionSettingCallOptions() *AdGroupExtensionSettingCallOptions {
	return &AdGroupExtensionSettingCallOptions{
		MutateAdGroupExtensionSettings: []gax.CallOption{
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

// internalAdGroupExtensionSettingClient is an interface that defines the methods available from Google Ads API.
type internalAdGroupExtensionSettingClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAdGroupExtensionSettings(context.Context, *servicespb.MutateAdGroupExtensionSettingsRequest, ...gax.CallOption) (*servicespb.MutateAdGroupExtensionSettingsResponse, error)
}

// AdGroupExtensionSettingClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ad group extension settings.
type AdGroupExtensionSettingClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupExtensionSettingClient

	// The call options for this service.
	CallOptions *AdGroupExtensionSettingCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupExtensionSettingClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupExtensionSettingClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AdGroupExtensionSettingClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAdGroupExtensionSettings creates, updates, or removes ad group extension settings. Operation
// statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// CriterionError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// ExtensionSettingError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// ListOperationError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperationAccessDeniedError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *AdGroupExtensionSettingClient) MutateAdGroupExtensionSettings(ctx context.Context, req *servicespb.MutateAdGroupExtensionSettingsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupExtensionSettingsResponse, error) {
	return c.internalClient.MutateAdGroupExtensionSettings(ctx, req, opts...)
}

// adGroupExtensionSettingGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupExtensionSettingGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AdGroupExtensionSettingClient
	CallOptions **AdGroupExtensionSettingCallOptions

	// The gRPC API client.
	adGroupExtensionSettingClient servicespb.AdGroupExtensionSettingServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewAdGroupExtensionSettingClient creates a new ad group extension setting service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ad group extension settings.
func NewAdGroupExtensionSettingClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupExtensionSettingClient, error) {
	clientOpts := defaultAdGroupExtensionSettingGRPCClientOptions()
	if newAdGroupExtensionSettingClientHook != nil {
		hookOpts, err := newAdGroupExtensionSettingClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AdGroupExtensionSettingClient{CallOptions: defaultAdGroupExtensionSettingCallOptions()}

	c := &adGroupExtensionSettingGRPCClient{
		connPool:                      connPool,
		adGroupExtensionSettingClient: servicespb.NewAdGroupExtensionSettingServiceClient(connPool),
		CallOptions:                   &client.CallOptions,
		logger:                        internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *adGroupExtensionSettingGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupExtensionSettingGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupExtensionSettingGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupExtensionSettingGRPCClient) MutateAdGroupExtensionSettings(ctx context.Context, req *servicespb.MutateAdGroupExtensionSettingsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupExtensionSettingsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateAdGroupExtensionSettings[0:len((*c.CallOptions).MutateAdGroupExtensionSettings):len((*c.CallOptions).MutateAdGroupExtensionSettings)], opts...)
	var resp *servicespb.MutateAdGroupExtensionSettingsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.adGroupExtensionSettingClient.MutateAdGroupExtensionSettings, req, settings.GRPC, c.logger, "MutateAdGroupExtensionSettings")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
