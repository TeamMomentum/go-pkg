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

var newUserDataClientHook clientHook

// UserDataCallOptions contains the retry settings for each method of UserDataClient.
type UserDataCallOptions struct {
	UploadUserData []gax.CallOption
}

func defaultUserDataGRPCClientOptions() []option.ClientOption {
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

func defaultUserDataCallOptions() *UserDataCallOptions {
	return &UserDataCallOptions{
		UploadUserData: []gax.CallOption{
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

// internalUserDataClient is an interface that defines the methods available from Google Ads API.
type internalUserDataClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	UploadUserData(context.Context, *servicespb.UploadUserDataRequest, ...gax.CallOption) (*servicespb.UploadUserDataResponse, error)
}

// UserDataClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage user data uploads.
// Any uploads made to a Customer Match list through this service will be
// eligible for matching as per the customer matching process. See
// https://support.google.com/google-ads/answer/7474263 (at https://support.google.com/google-ads/answer/7474263). However, the uploads
// made through this service will not be visible under the ‘Segment members’
// section for the Customer Match List in the Google Ads UI.
type UserDataClient struct {
	// The internal transport-dependent client.
	internalClient internalUserDataClient

	// The call options for this service.
	CallOptions *UserDataCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *UserDataClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *UserDataClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *UserDataClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// UploadUserData uploads the given user data.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// OfflineUserDataJobError (at )
// QuotaError (at )
// RequestError (at )
// UserDataError (at )
func (c *UserDataClient) UploadUserData(ctx context.Context, req *servicespb.UploadUserDataRequest, opts ...gax.CallOption) (*servicespb.UploadUserDataResponse, error) {
	return c.internalClient.UploadUserData(ctx, req, opts...)
}

// userDataGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type userDataGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing UserDataClient
	CallOptions **UserDataCallOptions

	// The gRPC API client.
	userDataClient servicespb.UserDataServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewUserDataClient creates a new user data service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage user data uploads.
// Any uploads made to a Customer Match list through this service will be
// eligible for matching as per the customer matching process. See
// https://support.google.com/google-ads/answer/7474263 (at https://support.google.com/google-ads/answer/7474263). However, the uploads
// made through this service will not be visible under the ‘Segment members’
// section for the Customer Match List in the Google Ads UI.
func NewUserDataClient(ctx context.Context, opts ...option.ClientOption) (*UserDataClient, error) {
	clientOpts := defaultUserDataGRPCClientOptions()
	if newUserDataClientHook != nil {
		hookOpts, err := newUserDataClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := UserDataClient{CallOptions: defaultUserDataCallOptions()}

	c := &userDataGRPCClient{
		connPool:       connPool,
		userDataClient: servicespb.NewUserDataServiceClient(connPool),
		CallOptions:    &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *userDataGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *userDataGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *userDataGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *userDataGRPCClient) UploadUserData(ctx context.Context, req *servicespb.UploadUserDataRequest, opts ...gax.CallOption) (*servicespb.UploadUserDataResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UploadUserData[0:len((*c.CallOptions).UploadUserData):len((*c.CallOptions).UploadUserData)], opts...)
	var resp *servicespb.UploadUserDataResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userDataClient.UploadUserData(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
