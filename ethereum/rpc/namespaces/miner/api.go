package miner

import (
	"errors"
	"math/big"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/tharsis/ethermint/ethereum/rpc/backend"
	"github.com/tharsis/ethermint/ethereum/rpc/namespaces/eth"
	rpctypes "github.com/tharsis/ethermint/ethereum/rpc/types"
	ethermint "github.com/tharsis/ethermint/types"
)

// API is the miner prefixed set of APIs in the Miner JSON-RPC spec.
type API struct {
	ctx     *server.Context
	logger  log.Logger
	ethAPI  *eth.PublicAPI
	backend backend.Backend
}

// NewMinerAPI creates an instance of the Miner API.
func NewMinerAPI(
	ctx *server.Context,
	ethAPI *eth.PublicAPI,
	backend backend.Backend,
) *API {
	return &API{
		ctx:     ctx,
		ethAPI:  ethAPI,
		logger:  ctx.Logger.With("module", "miner"),
		backend: backend,
	}
}

// SetEtherbase sets the etherbase of the miner
func (api *API) SetEtherbase(etherbase common.Address) bool {
	api.logger.Debug("miner_setEtherbase")

	delAddr, err := api.backend.GetCoinbase()
	if err != nil {
		api.logger.Debug("failed to get coinbase address", "error", err.Error())
		return false
	}

	withdrawAddr := sdk.AccAddress(etherbase.Bytes())
	msg := distributiontypes.NewMsgSetWithdrawAddress(delAddr, withdrawAddr)

	if err := msg.ValidateBasic(); err != nil {
		api.logger.Debug("tx failed basic validation", "error", err.Error())
		return false
	}

	// Assemble transaction from fields
	builder, ok := api.ethAPI.ClientCtx().TxConfig.NewTxBuilder().(authtx.ExtensionOptionsTxBuilder)
	if !ok {
		api.logger.Debug("clientCtx.TxConfig.NewTxBuilder returns unsupported builder", "error", err.Error())
	}

	err = builder.SetMsgs(msg)
	if err != nil {
		api.logger.Error("builder.SetMsgs failed", "error", err.Error())
	}

	denom, err := sdk.GetBaseDenom()
	if err != nil {
		api.logger.Debug("Could not get the denom of smallest unit registered.")
		return false
	}

	// req := &txtypes.SimulateRequest{
	// 	TxBytes: nil,
	// }

	// res, err := api.ethAPI.QueryClient().Simulate(api.ethAPI.Ctx(), req)
	// if err != nil {
	// 	api.logger.Debug("failed to simulate transaction to obtain gas estimation", "error", err.Error())
	// 	return false
	// }

	// res.GasInfo.GasUsed
	delCommonAddr := common.BytesToAddress(delAddr.Bytes())
	nonce, err := api.ethAPI.GetTransactionCount(delCommonAddr, rpctypes.EthPendingBlockNumber)
	if err != nil {
		api.logger.Debug("failed to get nonce", "error", err.Error())
		return false
	}

	txFactory := tx.Factory{}
	txFactory = txFactory.
		WithChainID(api.ethAPI.ClientCtx().ChainID).
		WithKeybase(api.ethAPI.ClientCtx().Keyring).
		WithTxConfig(api.ethAPI.ClientCtx().TxConfig).
		WithSequence(uint64(*nonce))

	_, gas, err := tx.CalculateGas(api.ethAPI.ClientCtx(), txFactory, msg)
	if err != nil {
		api.logger.Debug("failed to calculate gas", "error", err.Error())
	}
	//txFactory = txFactory.WithGas(gas)

	// TODO: is there a way to calculate this message fee and gas limit?
	// value := big.NewInt(gas)
	value := new(big.Int).SetUint64(gas)
	fees := sdk.Coins{sdk.NewCoin(denom, sdk.NewIntFromBigInt(value))}
	builder.SetFeeAmount(fees)
	builder.SetGasLimit(ethermint.DefaultRPCGasLimit)

	keyInfo, err := api.ethAPI.ClientCtx().Keyring.KeyByAddress(delAddr)
	if err != nil {
		return false
	}

	if err := tx.Sign(txFactory, keyInfo.GetName(), builder, false); err != nil {
		api.logger.Debug("failed to sign tx", "error", err.Error())
		return false
	}

	// Encode transaction by default Tx encoder
	txEncoder := api.ethAPI.ClientCtx().TxConfig.TxEncoder()
	txBytes, err := txEncoder(builder.GetTx())
	if err != nil {
		api.logger.Debug("failed to encode eth tx using default encoder", "error", err.Error())
		return false
	}

	tmHash := common.BytesToHash(tmtypes.Tx(txBytes).Hash())

	// Broadcast transaction in sync mode (default)
	// NOTE: If error is encountered on the node, the broadcast will not return an error
	syncCtx := api.ethAPI.ClientCtx().WithBroadcastMode(flags.BroadcastSync)
	rsp, err := syncCtx.BroadcastTx(txBytes)
	if err != nil || rsp.Code != 0 {
		if err == nil {
			err = errors.New(rsp.RawLog)
		}
		api.logger.Debug("failed to broadcast tx", "error", err.Error())
		return false
	}

	// ethermintd tx distribution withdraw-all-rewards

	api.logger.Debug("broadcasted tx to set miner withdraw address (etherbase)", "hash", tmHash.String())
	return true
}

// SetGasPrice sets the minimum accepted gas price for the miner.
// NOTE: this function accepts only integers to have the same interface than go-eth
// to use float values, the gas prices must be configured using the configuration file
func (api *API) SetGasPrice(gasPrice hexutil.Big) bool {
	api.logger.Info(api.ctx.Viper.ConfigFileUsed())
	appConf, err := config.ParseConfig(api.ctx.Viper)
	if err != nil {
		api.logger.Error("failed to parse file.", "file", api.ctx.Viper.ConfigFileUsed(), "error:", err.Error())
		return false
	}

	var unit string
	minGasPrices := appConf.GetMinGasPrices()

	// fetch the base denom from the sdk Config in case it's not currently defined on the node config
	if len(minGasPrices) == 0 || minGasPrices.Empty() {
		unit, err = sdk.GetBaseDenom()
		if err != nil {
			api.logger.Debug("Could not get the denom of smallest unit registered.")
			return false
		}
	} else {
		unit = minGasPrices[0].Denom
	}

	c := sdk.NewDecCoin(unit, sdk.NewIntFromBigInt(gasPrice.ToInt()))

	appConf.SetMinGasPrices(sdk.DecCoins{c})
	config.WriteConfigFile(api.ctx.Viper.ConfigFileUsed(), appConf)
	api.logger.Info("Your configuration file was modified. Please RESTART your node.", "gas-price", c.String())
	return true
}
