package example

import examplev1 "github.com/cosmosregistry/example/api/v1"

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *examplev1.GenesisState {
	return &examplev1.GenesisState{
		Params: DefaultParams(),
	}
}

// ValidateGenesis performs basic genesis state validation returning an error upon any
func ValidateGenesis(gs *examplev1.GenesisState) error {
	uniq := make(map[string]bool)
	for _, counter := range gs.Counters {
		if _, ok := uniq[counter.Address]; ok {
			return ErrDuplicateAddress
		}

		uniq[counter.Address] = true
	}

	if err := ValidateParams(gs.Params); err != nil {
		return err
	}

	return nil
}
