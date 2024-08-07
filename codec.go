package example

import (
	"cosmossdk.io/core/registry"
	"cosmossdk.io/core/transaction"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry registry.InterfaceRegistrar) {
	registry.RegisterImplementations((*transaction.Msg)(nil),
		&MsgUpdateParams{},
		&MsgIncrementCounter{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
