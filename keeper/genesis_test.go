package keeper_test

import (
	"testing"

	"cosmossdk.io/core/genesis"
	"github.com/cosmosregistry/example"
	"github.com/stretchr/testify/require"
)

func TestExportGenesis(t *testing.T) {
	fixture := initFixture(t)

	_, err := fixture.msgServer.IncrementCounter(fixture.ctx, &example.MsgIncrementCounter{
		Sender: fixture.addrs[0].String(),
	})
	require.NoError(t, err)

	target := &genesis.RawJSONTarget{}
	err = fixture.k.Schema.ExportGenesis(fixture.ctx, target.Target())
	require.NoError(t, err)
}
