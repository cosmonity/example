package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	queryv1beta1 "cosmossdk.io/api/cosmos/base/query/v1beta1"
	"github.com/cosmos/cosmos-sdk/types/query"
	examplev1 "github.com/cosmosregistry/example/api/v1"
)

var _ examplev1.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) examplev1.QueryServer {
	return queryServer{k: k}
}

type queryServer struct {
	examplev1.UnimplementedQueryServer

	k Keeper
}

// Counter defines the handler for the Query/Counter RPC method.
func (qs queryServer) Counter(ctx context.Context, req *examplev1.QueryCounterRequest) (*examplev1.QueryCounterResponse, error) {
	if _, err := qs.k.addressCodec.StringToBytes(req.Address); err != nil {
		return nil, fmt.Errorf("invalid sender address: %w", err)
	}

	counter, err := qs.k.Counter.Get(ctx, req.Address)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &examplev1.QueryCounterResponse{Counter: 0}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &examplev1.QueryCounterResponse{Counter: counter}, nil
}

// Counters defines the handler for the Query/Counters RPC method.
func (qs queryServer) Counters(ctx context.Context, req *examplev1.QueryCountersRequest) (*examplev1.QueryCountersResponse, error) {
	reqPagination := &query.PageRequest{
		Key:        req.Pagination.Key,
		Offset:     req.Pagination.Offset,
		Limit:      req.Pagination.Limit,
		CountTotal: req.Pagination.CountTotal,
		Reverse:    req.Pagination.Reverse,
	}

	counters, pageRes, err := query.CollectionPaginate(
		ctx,
		qs.k.Counter,
		reqPagination,
		func(key string, value uint64) (*examplev1.Counter, error) {
			return &examplev1.Counter{
				Address: key,
				Count:   value,
			}, nil
		})
	if err != nil {
		return nil, err
	}

	respPagination := &queryv1beta1.PageResponse{
		NextKey: pageRes.NextKey,
		Total:   pageRes.Total,
	} // TODO have helper in the SDK for this.

	return &examplev1.QueryCountersResponse{Counters: counters, Pagination: respPagination}, nil
}

// Params defines the handler for the Query/Params RPC method.
func (qs queryServer) Params(ctx context.Context, req *examplev1.QueryParamsRequest) (*examplev1.QueryParamsResponse, error) {
	params, err := qs.k.Params.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return &examplev1.QueryParamsResponse{Params: &examplev1.Params{}}, nil
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &examplev1.QueryParamsResponse{Params: &params}, nil
}
