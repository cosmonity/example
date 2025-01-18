package module

import (
	"context"
	"encoding/json"
	"fmt"

	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	"cosmossdk.io/core/registry"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmonity/example"
	"github.com/cosmonity/example/keeper"
)

var (
	_ appmodulev2.HasGenesis            = AppModule{}
	_ appmodulev2.AppModule             = AppModule{}
	_ appmodulev2.HasRegisterInterfaces = AppModule{}
	_ appmodulev2.HasConsensusVersion   = AppModule{}
	_ appmodulev2.HasMigrations         = AppModule{}
)

// ConsensusVersion defines the current module consensus version.
const ConsensusVersion = 1

type AppModule struct {
	cdc    codec.Codec
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: keeper,
	}
}

// RegisterLegacyAminoCodec registers the example module's types on the LegacyAmino codec.
func (AppModule) RegisterLegacyAminoCodec(registry.AminoRegistrar) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the example module.
func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := example.RegisterQueryHandlerClient(context.Background(), mux, example.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// RegisterInterfaces registers interfaces and implementations of the example module.
func (AppModule) RegisterInterfaces(registrar registry.InterfaceRegistrar) {
	example.RegisterInterfaces(registrar)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries.
func (am AppModule) RegisterServices(registrar grpc.ServiceRegistrar) {
	example.RegisterMsgServer(registrar, keeper.NewMsgServerImpl(am.keeper))
	example.RegisterQueryServer(registrar, keeper.NewQueryServerImpl(am.keeper))
}

// RegisterMigrations registers in place module state migration migrations
func (am AppModule) RegisterMigrations(mr appmodulev2.MigrationRegistrar) error {
	// m := keeper.NewMigrator(am.keeper)
	// if err := mr.Register(example.ModuleName, 1, m.Migrate1to2); err != nil {
	// 	return fmt.Errorf("failed to migrate x/%s from version 1 to 2: %v", example.ModuleName, err)
	// }

	return nil
}

// DefaultGenesis returns default genesis state as raw bytes for the module.
func (am AppModule) DefaultGenesis() json.RawMessage {
	return am.cdc.MustMarshalJSON(example.NewGenesisState())
}

// ValidateGenesis performs genesis state validation for the circuit module.
func (am AppModule) ValidateGenesis(bz json.RawMessage) error {
	var data example.GenesisState
	if err := am.cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", example.ModuleName, err)
	}

	return data.Validate()
}

// InitGenesis performs genesis initialization for the example module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx context.Context, data json.RawMessage) error {
	var genesisState example.GenesisState
	if err := am.cdc.UnmarshalJSON(data, &genesisState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", example.ModuleName, err)
	}

	if err := am.keeper.InitGenesis(ctx, &genesisState); err != nil {
		return fmt.Errorf("failed to initialize %s genesis state: %v", example.ModuleName, err)
	}

	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the circuit
// module.
func (am AppModule) ExportGenesis(ctx context.Context) (json.RawMessage, error) {
	gs, err := am.keeper.ExportGenesis(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to export %s genesis state: %v", example.ModuleName, err)
	}

	return am.cdc.MarshalJSON(gs)
}
