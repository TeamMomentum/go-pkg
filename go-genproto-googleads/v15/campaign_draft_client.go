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

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v15/services"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var newCampaignDraftClientHook clientHook

// CampaignDraftCallOptions contains the retry settings for each method of CampaignDraftClient.
type CampaignDraftCallOptions struct {
	MutateCampaignDrafts         []gax.CallOption
	PromoteCampaignDraft         []gax.CallOption
	ListCampaignDraftAsyncErrors []gax.CallOption
}

func defaultCampaignDraftGRPCClientOptions() []option.ClientOption {
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

func defaultCampaignDraftCallOptions() *CampaignDraftCallOptions {
	return &CampaignDraftCallOptions{
		MutateCampaignDrafts: []gax.CallOption{
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
		PromoteCampaignDraft: []gax.CallOption{
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
		ListCampaignDraftAsyncErrors: []gax.CallOption{
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

// internalCampaignDraftClient is an interface that defines the methods available from Google Ads API.
type internalCampaignDraftClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCampaignDrafts(context.Context, *servicespb.MutateCampaignDraftsRequest, ...gax.CallOption) (*servicespb.MutateCampaignDraftsResponse, error)
	PromoteCampaignDraft(context.Context, *servicespb.PromoteCampaignDraftRequest, ...gax.CallOption) (*PromoteCampaignDraftOperation, error)
	PromoteCampaignDraftOperation(name string) *PromoteCampaignDraftOperation
	ListCampaignDraftAsyncErrors(context.Context, *servicespb.ListCampaignDraftAsyncErrorsRequest, ...gax.CallOption) *StatusIterator
}

// CampaignDraftClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage campaign drafts.
type CampaignDraftClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignDraftClient

	// The call options for this service.
	CallOptions *CampaignDraftCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignDraftClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignDraftClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CampaignDraftClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCampaignDrafts creates, updates, or removes campaign drafts. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignDraftError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignDraftClient) MutateCampaignDrafts(ctx context.Context, req *servicespb.MutateCampaignDraftsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignDraftsResponse, error) {
	return c.internalClient.MutateCampaignDrafts(ctx, req, opts...)
}

// PromoteCampaignDraft promotes the changes in a draft back to the base campaign.
//
// This method returns a Long Running Operation (LRO) indicating if the
// Promote is done. Use [Operations.GetOperation] to poll the LRO until it
// is done. Only a done status is returned in the response. See the status
// in the Campaign Draft resource to determine if the promotion was
// successful. If the LRO failed, use
// CampaignDraftService.ListCampaignDraftAsyncErrors
// to view the list of error reasons.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignDraftError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignDraftClient) PromoteCampaignDraft(ctx context.Context, req *servicespb.PromoteCampaignDraftRequest, opts ...gax.CallOption) (*PromoteCampaignDraftOperation, error) {
	return c.internalClient.PromoteCampaignDraft(ctx, req, opts...)
}

// PromoteCampaignDraftOperation returns a new PromoteCampaignDraftOperation from a given name.
// The name must be that of a previously created PromoteCampaignDraftOperation, possibly from a different process.
func (c *CampaignDraftClient) PromoteCampaignDraftOperation(name string) *PromoteCampaignDraftOperation {
	return c.internalClient.PromoteCampaignDraftOperation(name)
}

// ListCampaignDraftAsyncErrors returns all errors that occurred during CampaignDraft promote. Throws an
// error if called before campaign draft is promoted.
// Supports standard list paging.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignDraftClient) ListCampaignDraftAsyncErrors(ctx context.Context, req *servicespb.ListCampaignDraftAsyncErrorsRequest, opts ...gax.CallOption) *StatusIterator {
	return c.internalClient.ListCampaignDraftAsyncErrors(ctx, req, opts...)
}

// campaignDraftGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignDraftGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CampaignDraftClient
	CallOptions **CampaignDraftCallOptions

	// The gRPC API client.
	campaignDraftClient servicespb.CampaignDraftServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCampaignDraftClient creates a new campaign draft service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage campaign drafts.
func NewCampaignDraftClient(ctx context.Context, opts ...option.ClientOption) (*CampaignDraftClient, error) {
	clientOpts := defaultCampaignDraftGRPCClientOptions()
	if newCampaignDraftClientHook != nil {
		hookOpts, err := newCampaignDraftClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CampaignDraftClient{CallOptions: defaultCampaignDraftCallOptions()}

	c := &campaignDraftGRPCClient{
		connPool:            connPool,
		campaignDraftClient: servicespb.NewCampaignDraftServiceClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *campaignDraftGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignDraftGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignDraftGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignDraftGRPCClient) MutateCampaignDrafts(ctx context.Context, req *servicespb.MutateCampaignDraftsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignDraftsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCampaignDrafts[0:len((*c.CallOptions).MutateCampaignDrafts):len((*c.CallOptions).MutateCampaignDrafts)], opts...)
	var resp *servicespb.MutateCampaignDraftsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignDraftClient.MutateCampaignDrafts(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *campaignDraftGRPCClient) PromoteCampaignDraft(ctx context.Context, req *servicespb.PromoteCampaignDraftRequest, opts ...gax.CallOption) (*PromoteCampaignDraftOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "campaign_draft", url.QueryEscape(req.GetCampaignDraft()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PromoteCampaignDraft[0:len((*c.CallOptions).PromoteCampaignDraft):len((*c.CallOptions).PromoteCampaignDraft)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignDraftClient.PromoteCampaignDraft(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &PromoteCampaignDraftOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *campaignDraftGRPCClient) ListCampaignDraftAsyncErrors(ctx context.Context, req *servicespb.ListCampaignDraftAsyncErrorsRequest, opts ...gax.CallOption) *StatusIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListCampaignDraftAsyncErrors[0:len((*c.CallOptions).ListCampaignDraftAsyncErrors):len((*c.CallOptions).ListCampaignDraftAsyncErrors)], opts...)
	it := &StatusIterator{}
	req = proto.Clone(req).(*servicespb.ListCampaignDraftAsyncErrorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*statuspb.Status, string, error) {
		resp := &servicespb.ListCampaignDraftAsyncErrorsResponse{}
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
			resp, err = c.campaignDraftClient.ListCampaignDraftAsyncErrors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetErrors(), resp.GetNextPageToken(), nil
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

// PromoteCampaignDraftOperation manages a long-running operation from PromoteCampaignDraft.
type PromoteCampaignDraftOperation struct {
	lro *longrunning.Operation
}

// PromoteCampaignDraftOperation returns a new PromoteCampaignDraftOperation from a given name.
// The name must be that of a previously created PromoteCampaignDraftOperation, possibly from a different process.
func (c *campaignDraftGRPCClient) PromoteCampaignDraftOperation(name string) *PromoteCampaignDraftOperation {
	return &PromoteCampaignDraftOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *PromoteCampaignDraftOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *PromoteCampaignDraftOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *PromoteCampaignDraftOperation) Metadata() (*emptypb.Empty, error) {
	var meta emptypb.Empty
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *PromoteCampaignDraftOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *PromoteCampaignDraftOperation) Name() string {
	return op.lro.Name()
}

// StatusIterator manages a stream of *statuspb.Status.
type StatusIterator struct {
	items    []*statuspb.Status
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*statuspb.Status, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *StatusIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *StatusIterator) Next() (*statuspb.Status, error) {
	var item *statuspb.Status
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *StatusIterator) bufLen() int {
	return len(it.items)
}

func (it *StatusIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
