package integration

import (
	"testing"

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
	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/stretchr/testify/require"

	"github.com/cosmosregistry/example"
	modulev1 "github.com/cosmosregistry/example/api/module/v1"
	"github.com/cosmosregistry/example/keeper"
)

// ExampleModule is a configurator.ModuleOption that adds the example module to the app config.
func ExampleModule() configurator.ModuleOption {
	return func(config *configurator.Config) {
		config.ModuleConfigs[example.ModuleName] = &appv1alpha1.ModuleConfig{
			Name:   example.ModuleName,
			Config: appconfig.WrapAny(&modulev1.Module{}),
		}
	}
}

func TestIntegration(t *testing.T) {
	t.Parallel()

	var keeper keeper.Keeper
	AppConfig := configurator.NewAppConfig(
		configurator.AuthModule(),
		configurator.BankModule(),
		configurator.StakingModule(),
		configurator.TxModule(),
		configurator.ConsensusModule(),
		configurator.GenutilModule(),
		configurator.MintModule(),
		ExampleModule(),
	)

	_, err := simtestutil.Setup(AppConfig, &keeper)
	require.NoError(t, err)
}
