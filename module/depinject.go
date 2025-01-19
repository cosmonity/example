package module

import (
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/codec"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"

	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"

	"go.cosmonity.xyz/example"
	"go.cosmonity.xyz/example/keeper"
)

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func init() {
	appconfig.Register(
		&example.Module{},
		appconfig.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
	AddressCodec address.Codec

	Config *example.Module
}

type ModuleOutputs struct {
	depinject.Out

	Module appmodule.AppModule
	Keeper keeper.Keeper
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance as authority if not provided
	var (
		authorityStr string
		err          error
	)

	authority := sdkaddress.Module("gov")
	if in.Config.Authority != "" {
		if configAuthority, err := in.AddressCodec.StringToBytes(in.Config.Authority); err == nil {
			authority = configAuthority
		} else {
			authority = sdkaddress.Module(in.Config.Authority)
		}
	}

	authorityStr, err = in.AddressCodec.BytesToString(authority)
	if err != nil {
		panic(err)
	}

	k := keeper.NewKeeper(in.Cdc, in.AddressCodec, in.StoreService, authorityStr)
	m := NewAppModule(in.Cdc, k)

	return ModuleOutputs{Module: m, Keeper: k}
}
