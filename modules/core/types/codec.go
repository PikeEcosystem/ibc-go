package types

import (
	codectypes "github.com/PikeEcosystem/cosmos-sdk/codec/types"

	clienttypes "github.com/PikeEcosystem/ibc-go/v3/modules/core/02-client/types"
	connectiontypes "github.com/PikeEcosystem/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/PikeEcosystem/ibc-go/v3/modules/core/04-channel/types"
	commitmenttypes "github.com/PikeEcosystem/ibc-go/v3/modules/core/23-commitment/types"
	solomachinetypes "github.com/PikeEcosystem/ibc-go/v3/modules/light-clients/06-solomachine/types"
	ibcoctypes "github.com/PikeEcosystem/ibc-go/v3/modules/light-clients/07-tendermint/types"
	localhosttypes "github.com/PikeEcosystem/ibc-go/v3/modules/light-clients/09-localhost/types"
)

// RegisterInterfaces registers x/ibc interfaces into protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	clienttypes.RegisterInterfaces(registry)
	connectiontypes.RegisterInterfaces(registry)
	channeltypes.RegisterInterfaces(registry)
	solomachinetypes.RegisterInterfaces(registry)
	ibcoctypes.RegisterInterfaces(registry)
	localhosttypes.RegisterInterfaces(registry)
	commitmenttypes.RegisterInterfaces(registry)
}
