package client

import (
	"github.com/osmosis-labs/osmosis/v19/x/poolmanager/client/cli"
	"github.com/osmosis-labs/osmosis/v19/x/poolmanager/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	DenomPairTakerFeeProposalHandler = govclient.NewProposalHandler(cli.NewCmdHandleDenomPairTakerFeeProposal, rest.ProposalDenomPairTakerFeeRESTHandler)
)