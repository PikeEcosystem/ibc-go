package client

import (
	govclient "github.com/PikeEcosystem/cosmos-sdk/x/gov/client"

	"github.com/PikeEcosystem/ibc-go/v3/modules/core/02-client/client/cli"
)

var (
	UpdateClientProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateClientProposal)
	UpgradeProposalHandler      = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
)
