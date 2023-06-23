package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	// blank import for app wiring registration
	_ "github.com/cosmos/cosmos-sdk/x/auth"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	_ "github.com/cosmos/cosmos-sdk/x/bank"
	_ "github.com/cosmos/cosmos-sdk/x/consensus"
	_ "github.com/cosmos/cosmos-sdk/x/genutil"
	_ "github.com/cosmos/cosmos-sdk/x/mint"
	_ "github.com/cosmos/cosmos-sdk/x/staking"

	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	"github.com/cosmosregistry/example"
	examplemodulev1 "github.com/cosmosregistry/example/api/module/v1"
	"github.com/cosmosregistry/example/keeper"
	_ "github.com/cosmosregistry/example/module"
)

// ExampleModule is a configurator.ModuleOption that adds the example module to the app config.
func ExampleModule() configurator.ModuleOption {
	return func(config *configurator.Config) {
		config.ModuleConfigs[example.ModuleName] = &appv1alpha1.ModuleConfig{
			Name:   example.ModuleName,
			Config: appconfig.WrapAny(&examplemodulev1.Module{}),
		}
	}
}

func TestIntegration(t *testing.T) {
	t.Parallel()

	appConfig := configurator.NewAppConfig(
		configurator.AuthModule(),
		configurator.BankModule(),
		configurator.StakingModule(),
		configurator.TxModule(),
		configurator.ConsensusModule(),
		configurator.GenutilModule(),
		configurator.MintModule(),
		ExampleModule(),
	)
	logger := log.NewTestLogger(t)

	var keeper keeper.Keeper
	_, err := simtestutil.Setup(
		depinject.Supply(appConfig, logger),
		&keeper,
	)
	require.NoError(t, err)
}
