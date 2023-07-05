package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/PikeEcosystem/cosmos-sdk/codec"
	sdk "github.com/PikeEcosystem/cosmos-sdk/types"
	ocbytes "github.com/PikeEcosystem/tendermint/libs/bytes"
	octypes "github.com/PikeEcosystem/tendermint/types"

	clienttypes "github.com/PikeEcosystem/ibc-go/v3/modules/core/02-client/types"
	ibcoctypes "github.com/PikeEcosystem/ibc-go/v3/modules/light-clients/07-tendermint/types"
	ibctesting "github.com/PikeEcosystem/ibc-go/v3/testing"
	ibctestingmock "github.com/PikeEcosystem/ibc-go/v3/testing/mock"
	"github.com/PikeEcosystem/ibc-go/v3/testing/simapp"
)

const (
	chainID                        = "fnsa"
	chainIDRevision0               = "fnsa-revision-0"
	chainIDRevision1               = "fnsa-revision-1"
	clientID                       = "fnsamainnet"
	trustingPeriod   time.Duration = time.Hour * 24 * 7 * 2
	ubdPeriod        time.Duration = time.Hour * 24 * 7 * 3
	maxClockDrift    time.Duration = time.Second * 10
)

var (
	height          = clienttypes.NewHeight(0, 4)
	newClientHeight = clienttypes.NewHeight(1, 1)
	upgradePath     = []string{"upgrade", "upgradedIBCState"}
)

type TendermintTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain

	// TODO: deprecate usage in favor of testing package
	ctx        sdk.Context
	cdc        codec.Codec
	privVal    octypes.PrivValidator
	valSet     *octypes.ValidatorSet
	valsHash   ocbytes.HexBytes
	header     *ibcoctypes.Header
	now        time.Time
	headerTime time.Time
	clientTime time.Time
}

func (suite *TendermintTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))
	// commit some blocks so that QueryProof returns valid proof (cannot return valid query if height <= 1)
	suite.coordinator.CommitNBlocks(suite.chainA, 2)
	suite.coordinator.CommitNBlocks(suite.chainB, 2)

	// TODO: deprecate usage in favor of testing package
	checkTx := false
	app := simapp.Setup(checkTx)

	suite.cdc = app.AppCodec()

	// now is the time of the current chain, must be after the updating header
	// mocks ctx.BlockTime()
	suite.now = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	suite.clientTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	// Header time is intended to be time for any new header used for updates
	suite.headerTime = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)

	suite.privVal = ibctestingmock.NewPV()

	pubKey, err := suite.privVal.GetPubKey()
	suite.Require().NoError(err)

	heightMinus1 := clienttypes.NewHeight(0, height.RevisionHeight-1)

	val := ibctesting.NewTestValidator(pubKey, 10)
	suite.valSet = octypes.NewValidatorSet([]*octypes.Validator{val})
	suite.valsHash = suite.valSet.Hash()
	suite.header = suite.chainA.CreateOCClientHeader(chainID, int64(height.RevisionHeight), heightMinus1, suite.now, suite.valSet, suite.valSet, []octypes.PrivValidator{suite.privVal})
	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1, Time: suite.now})
}

func getSuiteSigners(suite *TendermintTestSuite) []octypes.PrivValidator {
	return []octypes.PrivValidator{suite.privVal}
}

func getBothSigners(suite *TendermintTestSuite, altVal *octypes.Validator, altPrivVal octypes.PrivValidator) (*octypes.ValidatorSet, []octypes.PrivValidator) {
	// Create bothValSet with both suite validator and altVal. Would be valid update
	bothValSet := octypes.NewValidatorSet(append(suite.valSet.Validators, altVal))
	// Create signer array and ensure it is in same order as bothValSet
	_, suiteVal := suite.valSet.GetByIndex(0)
	bothSigners := ibctesting.CreateSortedSignerArray(altPrivVal, suite.privVal, altVal, suiteVal)
	return bothValSet, bothSigners
}

func TestTendermintTestSuite(t *testing.T) {
	suite.Run(t, new(TendermintTestSuite))
}
