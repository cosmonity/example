package keeper

import (
	"context"

	"github.com/cosmosregistry/example"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *example.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	// Here you need to import all your module state.
	// Until core genesis is finished, which will do that for you.

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*example.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	// Here you need to export all your module state.
	// Until core genesis is finished, which will do that for you.

	return &example.GenesisState{
		Params: params,
	}, nil
}
