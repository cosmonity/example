package keeper_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"cosmossdk.io/core/genesis"
	"github.com/cosmosregistry/example"
	"github.com/stretchr/testify/require"
)

func TestDefaultGenesis(t *testing.T) {
	fixture := initFixture(t)

	target := &genesis.RawJSONTarget{}
	err := fixture.k.Schema.DefaultGenesis(target.Target())
	require.NoError(t, err)

	result, err := target.JSON()
	require.NoError(t, err)
	buf := &bytes.Buffer{}
	err = json.Compact(buf, result)
	require.NoError(t, err)

	require.Equal(t, `{"counter":[],"params":[]}`, buf.String())
}

func TestExportGenesis(t *testing.T) {
	fixture := initFixture(t)

	_, err := fixture.msgServer.IncrementCounter(fixture.ctx, &example.MsgIncrementCounter{
		Sender: fixture.addrs[0].String(),
	})
	require.NoError(t, err)

	target := &genesis.RawJSONTarget{}
	err = fixture.k.Schema.ExportGenesis(fixture.ctx, target.Target())
	require.NoError(t, err)

	result, err := target.JSON()
	require.NoError(t, err)
	buf := &bytes.Buffer{}
	err = json.Compact(buf, result)
	require.NoError(t, err)

	require.Equal(t, `{"counter":[{"key":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","value":"1"}],"params":[]}`, buf.String())
}
