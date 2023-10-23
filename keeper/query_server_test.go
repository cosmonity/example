package keeper_test

import (
	"testing"

	queryv1beta1 "cosmossdk.io/api/cosmos/base/query/v1beta1"
	"github.com/stretchr/testify/require"

	examplev1 "github.com/cosmosregistry/example/api/v1"
)

func TestQueryParams(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Params(f.ctx, &examplev1.QueryParamsRequest{})
	require.NoError(err)
	require.Equal(examplev1.Params{}, resp.Params)
}

func TestQueryCounter(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Counter(f.ctx, &examplev1.QueryCounterRequest{Address: f.addrs[0].String()})
	require.NoError(err)
	require.Equal(uint64(0), resp.Counter)

	_, err = f.msgServer.IncrementCounter(f.ctx, &examplev1.MsgIncrementCounter{Sender: f.addrs[0].String()})
	require.NoError(err)

	resp, err = f.queryServer.Counter(f.ctx, &examplev1.QueryCounterRequest{Address: f.addrs[0].String()})
	require.NoError(err)
	require.Equal(uint64(1), resp.Counter)
}

func TestQueryCounters(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Counters(f.ctx, &examplev1.QueryCountersRequest{})
	require.NoError(err)
	require.Equal(0, len(resp.Counters))

	_, err = f.msgServer.IncrementCounter(f.ctx, &examplev1.MsgIncrementCounter{Sender: f.addrs[0].String()})
	require.NoError(err)

	resp, err = f.queryServer.Counters(f.ctx, &examplev1.QueryCountersRequest{})
	require.NoError(err)
	require.Equal(1, len(resp.Counters))
	require.Equal(uint64(1), resp.Counters[0].Count)
	require.Equal(f.addrs[0].String(), resp.Counters[0].Address)
}

func TestQueryCountersPaginated(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Counters(f.ctx, &examplev1.QueryCountersRequest{Pagination: &queryv1beta1.PageRequest{Limit: 1}})
	require.NoError(err)
	require.Equal(0, len(resp.Counters))

	_, err = f.msgServer.IncrementCounter(f.ctx, &examplev1.MsgIncrementCounter{Sender: f.addrs[0].String()})
	require.NoError(err)
	_, err = f.msgServer.IncrementCounter(f.ctx, &examplev1.MsgIncrementCounter{Sender: f.addrs[1].String()})
	require.NoError(err)

	resp, err = f.queryServer.Counters(f.ctx, &examplev1.QueryCountersRequest{Pagination: &queryv1beta1.PageRequest{Limit: 1}})
	require.NoError(err)
	require.Equal(1, len(resp.Counters))
	require.Equal(uint64(1), resp.Counters[0].Count)
	require.Equal(f.addrs[1].String(), resp.Counters[0].Address)

	resp, err = f.queryServer.Counters(f.ctx, &examplev1.QueryCountersRequest{})
	require.NoError(err)
	require.Equal(2, len(resp.Counters))
}
