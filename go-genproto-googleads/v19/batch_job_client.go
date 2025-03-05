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

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v19/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newBatchJobClientHook clientHook

// BatchJobCallOptions contains the retry settings for each method of BatchJobClient.
type BatchJobCallOptions struct {
	MutateBatchJob        []gax.CallOption
	ListBatchJobResults   []gax.CallOption
	RunBatchJob           []gax.CallOption
	AddBatchJobOperations []gax.CallOption
}

func defaultBatchJobGRPCClientOptions() []option.ClientOption {
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

func defaultBatchJobCallOptions() *BatchJobCallOptions {
	return &BatchJobCallOptions{
		MutateBatchJob: []gax.CallOption{
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
		ListBatchJobResults: []gax.CallOption{
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
		RunBatchJob: []gax.CallOption{
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
		AddBatchJobOperations: []gax.CallOption{
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

// internalBatchJobClient is an interface that defines the methods available from Google Ads API.
type internalBatchJobClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateBatchJob(context.Context, *servicespb.MutateBatchJobRequest, ...gax.CallOption) (*servicespb.MutateBatchJobResponse, error)
	ListBatchJobResults(context.Context, *servicespb.ListBatchJobResultsRequest, ...gax.CallOption) *BatchJobResultIterator
	RunBatchJob(context.Context, *servicespb.RunBatchJobRequest, ...gax.CallOption) (*RunBatchJobOperation, error)
	RunBatchJobOperation(name string) *RunBatchJobOperation
	AddBatchJobOperations(context.Context, *servicespb.AddBatchJobOperationsRequest, ...gax.CallOption) (*servicespb.AddBatchJobOperationsResponse, error)
}

// BatchJobClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage batch jobs.
type BatchJobClient struct {
	// The internal transport-dependent client.
	internalClient internalBatchJobClient

	// The call options for this service.
	CallOptions *BatchJobCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BatchJobClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BatchJobClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BatchJobClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateBatchJob mutates a batch job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
func (c *BatchJobClient) MutateBatchJob(ctx context.Context, req *servicespb.MutateBatchJobRequest, opts ...gax.CallOption) (*servicespb.MutateBatchJobResponse, error) {
	return c.internalClient.MutateBatchJob(ctx, req, opts...)
}

// ListBatchJobResults returns the results of the batch job. The job must be done.
// Supports standard list paging.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// BatchJobError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *BatchJobClient) ListBatchJobResults(ctx context.Context, req *servicespb.ListBatchJobResultsRequest, opts ...gax.CallOption) *BatchJobResultIterator {
	return c.internalClient.ListBatchJobResults(ctx, req, opts...)
}

// RunBatchJob runs the batch job.
//
// The Operation.metadata field type is BatchJobMetadata. When finished, the
// long running operation will not contain errors or a response. Instead, use
// ListBatchJobResults to get the results of the job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// BatchJobError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *BatchJobClient) RunBatchJob(ctx context.Context, req *servicespb.RunBatchJobRequest, opts ...gax.CallOption) (*RunBatchJobOperation, error) {
	return c.internalClient.RunBatchJob(ctx, req, opts...)
}

// RunBatchJobOperation returns a new RunBatchJobOperation from a given name.
// The name must be that of a previously created RunBatchJobOperation, possibly from a different process.
func (c *BatchJobClient) RunBatchJobOperation(name string) *RunBatchJobOperation {
	return c.internalClient.RunBatchJobOperation(name)
}

// AddBatchJobOperations add operations to the batch job.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// BatchJobError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
func (c *BatchJobClient) AddBatchJobOperations(ctx context.Context, req *servicespb.AddBatchJobOperationsRequest, opts ...gax.CallOption) (*servicespb.AddBatchJobOperationsResponse, error) {
	return c.internalClient.AddBatchJobOperations(ctx, req, opts...)
}

// batchJobGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type batchJobGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BatchJobClient
	CallOptions **BatchJobCallOptions

	// The gRPC API client.
	batchJobClient servicespb.BatchJobServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewBatchJobClient creates a new batch job service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage batch jobs.
func NewBatchJobClient(ctx context.Context, opts ...option.ClientOption) (*BatchJobClient, error) {
	clientOpts := defaultBatchJobGRPCClientOptions()
	if newBatchJobClientHook != nil {
		hookOpts, err := newBatchJobClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BatchJobClient{CallOptions: defaultBatchJobCallOptions()}

	c := &batchJobGRPCClient{
		connPool:       connPool,
		batchJobClient: servicespb.NewBatchJobServiceClient(connPool),
		CallOptions:    &client.CallOptions,
		logger:         internaloption.GetLogger(opts),
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
func (c *batchJobGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *batchJobGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *batchJobGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *batchJobGRPCClient) MutateBatchJob(ctx context.Context, req *servicespb.MutateBatchJobRequest, opts ...gax.CallOption) (*servicespb.MutateBatchJobResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateBatchJob[0:len((*c.CallOptions).MutateBatchJob):len((*c.CallOptions).MutateBatchJob)], opts...)
	var resp *servicespb.MutateBatchJobResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.batchJobClient.MutateBatchJob, req, settings.GRPC, c.logger, "MutateBatchJob")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *batchJobGRPCClient) ListBatchJobResults(ctx context.Context, req *servicespb.ListBatchJobResultsRequest, opts ...gax.CallOption) *BatchJobResultIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListBatchJobResults[0:len((*c.CallOptions).ListBatchJobResults):len((*c.CallOptions).ListBatchJobResults)], opts...)
	it := &BatchJobResultIterator{}
	req = proto.Clone(req).(*servicespb.ListBatchJobResultsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*servicespb.BatchJobResult, string, error) {
		resp := &servicespb.ListBatchJobResultsResponse{}
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
			resp, err = executeRPC(ctx, c.batchJobClient.ListBatchJobResults, req, settings.GRPC, c.logger, "ListBatchJobResults")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResults(), resp.GetNextPageToken(), nil
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

func (c *batchJobGRPCClient) RunBatchJob(ctx context.Context, req *servicespb.RunBatchJobRequest, opts ...gax.CallOption) (*RunBatchJobOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RunBatchJob[0:len((*c.CallOptions).RunBatchJob):len((*c.CallOptions).RunBatchJob)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.batchJobClient.RunBatchJob, req, settings.GRPC, c.logger, "RunBatchJob")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RunBatchJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *batchJobGRPCClient) AddBatchJobOperations(ctx context.Context, req *servicespb.AddBatchJobOperationsRequest, opts ...gax.CallOption) (*servicespb.AddBatchJobOperationsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddBatchJobOperations[0:len((*c.CallOptions).AddBatchJobOperations):len((*c.CallOptions).AddBatchJobOperations)], opts...)
	var resp *servicespb.AddBatchJobOperationsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.batchJobClient.AddBatchJobOperations, req, settings.GRPC, c.logger, "AddBatchJobOperations")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RunBatchJobOperation returns a new RunBatchJobOperation from a given name.
// The name must be that of a previously created RunBatchJobOperation, possibly from a different process.
func (c *batchJobGRPCClient) RunBatchJobOperation(name string) *RunBatchJobOperation {
	return &RunBatchJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
