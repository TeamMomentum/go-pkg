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

var newAccountBudgetProposalClientHook clientHook

// AccountBudgetProposalCallOptions contains the retry settings for each method of AccountBudgetProposalClient.
type AccountBudgetProposalCallOptions struct {
	MutateAccountBudgetProposal []gax.CallOption
}

func defaultAccountBudgetProposalGRPCClientOptions() []option.ClientOption {
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

func defaultAccountBudgetProposalCallOptions() *AccountBudgetProposalCallOptions {
	return &AccountBudgetProposalCallOptions{
		MutateAccountBudgetProposal: []gax.CallOption{
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

// internalAccountBudgetProposalClient is an interface that defines the methods available from Google Ads API.
type internalAccountBudgetProposalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAccountBudgetProposal(context.Context, *servicespb.MutateAccountBudgetProposalRequest, ...gax.CallOption) (*servicespb.MutateAccountBudgetProposalResponse, error)
}

// AccountBudgetProposalClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing account-level budgets through proposals.
//
// A proposal is a request to create a new budget or make changes to an
// existing one.
//
// Mutates:
// The CREATE operation creates a new proposal.
// UPDATE operations aren’t supported.
// The REMOVE operation cancels a pending proposal.
type AccountBudgetProposalClient struct {
	// The internal transport-dependent client.
	internalClient internalAccountBudgetProposalClient

	// The call options for this service.
	CallOptions *AccountBudgetProposalCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AccountBudgetProposalClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AccountBudgetProposalClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AccountBudgetProposalClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAccountBudgetProposal creates, updates, or removes account budget proposals.  Operation statuses
// are returned.
//
// List of thrown errors:
// AccountBudgetProposalError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// DateError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
// StringLengthError (at )
func (c *AccountBudgetProposalClient) MutateAccountBudgetProposal(ctx context.Context, req *servicespb.MutateAccountBudgetProposalRequest, opts ...gax.CallOption) (*servicespb.MutateAccountBudgetProposalResponse, error) {
	return c.internalClient.MutateAccountBudgetProposal(ctx, req, opts...)
}

// accountBudgetProposalGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type accountBudgetProposalGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AccountBudgetProposalClient
	CallOptions **AccountBudgetProposalCallOptions

	// The gRPC API client.
	accountBudgetProposalClient servicespb.AccountBudgetProposalServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewAccountBudgetProposalClient creates a new account budget proposal service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing account-level budgets through proposals.
//
// A proposal is a request to create a new budget or make changes to an
// existing one.
//
// Mutates:
// The CREATE operation creates a new proposal.
// UPDATE operations aren’t supported.
// The REMOVE operation cancels a pending proposal.
func NewAccountBudgetProposalClient(ctx context.Context, opts ...option.ClientOption) (*AccountBudgetProposalClient, error) {
	clientOpts := defaultAccountBudgetProposalGRPCClientOptions()
	if newAccountBudgetProposalClientHook != nil {
		hookOpts, err := newAccountBudgetProposalClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AccountBudgetProposalClient{CallOptions: defaultAccountBudgetProposalCallOptions()}

	c := &accountBudgetProposalGRPCClient{
		connPool:                    connPool,
		accountBudgetProposalClient: servicespb.NewAccountBudgetProposalServiceClient(connPool),
		CallOptions:                 &client.CallOptions,
		logger:                      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *accountBudgetProposalGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *accountBudgetProposalGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *accountBudgetProposalGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *accountBudgetProposalGRPCClient) MutateAccountBudgetProposal(ctx context.Context, req *servicespb.MutateAccountBudgetProposalRequest, opts ...gax.CallOption) (*servicespb.MutateAccountBudgetProposalResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateAccountBudgetProposal[0:len((*c.CallOptions).MutateAccountBudgetProposal):len((*c.CallOptions).MutateAccountBudgetProposal)], opts...)
	var resp *servicespb.MutateAccountBudgetProposalResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.accountBudgetProposalClient.MutateAccountBudgetProposal, req, settings.GRPC, c.logger, "MutateAccountBudgetProposal")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
