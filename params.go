package example

import examplev1 "github.com/cosmosregistry/example/api/v1"

// DefaultParams returns default module parameters.
func DefaultParams() *examplev1.Params {
	return &examplev1.Params{
		// Set default values here.
	}
}

// ValidateParams does the sanity check on the params.
func ValidateParams(p *examplev1.Params) error {
	// Sanity check goes here.
	return nil
}
