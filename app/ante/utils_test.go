package ante_test

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/ethermint/app"
	ante "github.com/cosmos/ethermint/app/ante"
	"github.com/cosmos/ethermint/crypto/ethsecp256k1"
	evmtypes "github.com/cosmos/ethermint/x/evm/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type AnteTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *app.EthermintApp
	clientCtx   client.Context
	txBuilder   client.TxBuilder
	anteHandler sdk.AnteHandler
}

func (suite *AnteTestSuite) SetupTest() {
	checkTx := false
	suite.app = app.Setup(checkTx)
	suite.ctx = suite.app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 2, ChainID: "ethermint-1", Time: time.Now().UTC()})
	suite.ctx = suite.ctx.WithMinGasPrices(sdk.NewDecCoins(sdk.NewDecCoin(evmtypes.DefaultEVMDenom, sdk.OneInt())))
	suite.ctx = suite.ctx.WithBlockGasMeter(sdk.NewGasMeter(1000000000000000000))
	suite.app.EvmKeeper.WithChainID(suite.ctx)

	infCtx := suite.ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	suite.app.AccountKeeper.SetParams(infCtx, authtypes.DefaultParams())
	suite.app.EvmKeeper.SetParams(infCtx, evmtypes.DefaultParams())

	encodingConfig := app.MakeEncodingConfig()
	// We're using TestMsg amino encoding in some tests, so register it here.
	encodingConfig.Amino.RegisterConcrete(&testdata.TestMsg{}, "testdata.TestMsg", nil)

	suite.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	suite.anteHandler = ante.NewAnteHandler(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.EvmKeeper, encodingConfig.TxConfig.SignModeHandler())
}

func TestAnteTestSuite(t *testing.T) {
	suite.Run(t, new(AnteTestSuite))
}

// CreateTestTx is a helper function to create a tx given multiple inputs.
func (suite *AnteTestSuite) CreateTestTx(privs []cryptotypes.PrivKey, accNums []uint64, accSeqs []uint64) (authsigning.Tx, error) {
	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  suite.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err := suite.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := authsigning.SignerData{
			ChainID:       suite.ctx.ChainID(),
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		sigV2, err := tx.SignWithPrivKey(
			suite.clientCtx.TxConfig.SignModeHandler().DefaultMode(), signerData,
			suite.txBuilder, priv, suite.clientCtx.TxConfig, accSeqs[i])
		if err != nil {
			return nil, err
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = suite.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	return suite.txBuilder.GetTx(), nil
}

func newTestAddrKey() (common.Address, cryptotypes.PrivKey) {
	privkey, _ := ethsecp256k1.GenerateKey()
	addr := crypto.PubkeyToAddress(privkey.ToECDSA().PublicKey)

	return addr, privkey
}

var _ sdk.Tx = &invalidTx{}

type invalidTx struct{}

func (invalidTx) GetMsgs() []sdk.Msg   { return []sdk.Msg{nil} }
func (invalidTx) ValidateBasic() error { return nil }
