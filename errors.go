package example

import "cosmossdk.io/errors/v2"

var (
	// ErrDuplicateAddress error if there is a duplicate address
	ErrDuplicateAddress = errors.Register(ModuleName, 2, "duplicate address")
)
