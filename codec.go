package example

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	examplev1 "github.com/cosmosregistry/example/api/v1"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&examplev1.MsgUpdateParams{},
		&examplev1.MsgIncrementCounter{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &examplev1.Msg_ServiceDesc)
}
