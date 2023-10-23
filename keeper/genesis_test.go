package keeper_test

import (
	"testing"

	"github.com/cosmosregistry/example"
	examplev1 "github.com/cosmosregistry/example/api/v1"
	"github.com/stretchr/testify/require"
)

func TestInitGenesis(t *testing.T) {
	fixture := initFixture(t)

	data := &examplev1.GenesisState{
		Counters: []*examplev1.Counter{
			{
				Address: fixture.addrs[0].String(),
				Count:   5,
			},
		},
		Params: example.DefaultParams(),
	}
	err := fixture.k.InitGenesis(fixture.ctx, data)
	require.NoError(t, err)

	params, err := fixture.k.Params.Get(fixture.ctx)
	require.NoError(t, err)

	require.Equal(t, example.DefaultParams(), params)

	count, err := fixture.k.Counter.Get(fixture.ctx, fixture.addrs[0].String())
	require.NoError(t, err)
	require.Equal(t, uint64(5), count)
}

func TestExportGenesis(t *testing.T) {
	fixture := initFixture(t)

	_, err := fixture.msgServer.IncrementCounter(fixture.ctx, &examplev1.MsgIncrementCounter{
		Sender: fixture.addrs[0].String(),
	})
	require.NoError(t, err)

	out, err := fixture.k.ExportGenesis(fixture.ctx)
	require.NoError(t, err)

	require.Equal(t, example.DefaultParams(), out.Params)
	require.Equal(t, uint64(1), out.Counters[0].Count)
}
