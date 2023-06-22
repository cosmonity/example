package integration

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/stretchr/testify/require"

	"github.com/cosmosregistry/example"
)

// func ExampleModule() configurator.ModuleOption {
// 	return func(config *configurator.Config) {
// 		config.ModuleConfigs["example"] = &appv1alpha1.ModuleConfig{
// 			Name:   example.ModuleName,
// 			Config: appconfig.WrapAny(&examplev1.Module{}),
// 		}
// 	}
// }

func TestIntegration(t *testing.T) {
	t.Parallel()

	var keeper example.Keeper
	AppConfig := configurator.NewAppConfig(
		configurator.AuthModule(),
		configurator.BankModule(),
		configurator.StakingModule(),
		configurator.TxModule(),
		configurator.ConsensusModule(),
		configurator.ParamsModule(),
		configurator.GenutilModule(),
		configurator.MintModule(),
		// ExampleModule,
	)

	_, err := simtestutil.Setup(AppConfig, &keeper)
	require.NoError(t, err)
}
