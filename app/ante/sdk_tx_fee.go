package ante

import (
	"fmt"
	"math"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ethermint "github.com/evmos/ethermint/types"
	"github.com/evmos/ethermint/x/evm/types"
)

// NewSDKTxFeeChecker returns a `TxFeeChecker` that applies feemarket to cosmos tx.
// could be called in both checkTx and deliverTx modes.
//
// - use the same eip-1559 mechanism as the evm transactions:
//   - feeCap = tx.fees / tx.gas
//   - tipFeeCap = tx.extOpt.MaxPriorityPrice or MaxInt64
// - when `ExtensionOptionDynamicFeeTx` is omitted, `tipFeeCap` defaults to `MaxInt64`.
// - when london hardfork is not enabled, it fallbacks to sdk default behaviour (validator min-gas-prices).
// - tx priority is set to `effectiveGasPrice / DefaultPriorityReduction`.
func NewSDKTxFeeChecker(k EVMKeeper) authante.TxFeeChecker {
	return func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		if ctx.BlockHeight() == 0 {
			// genesis transactions: fallback to min-gas-price logic
			return checkTxFeeWithValidatorMinGasPrices(ctx, tx)
		}

		params := k.GetParams(ctx)
		denom := params.EvmDenom
		ethCfg := params.ChainConfig.EthereumConfig(k.ChainID())
		baseFee := k.GetBaseFee(ctx, ethCfg)

		if baseFee == nil {
			// london hardfork is not enabled: fallback to min-gas-prices logic
			return checkTxFeeWithValidatorMinGasPrices(ctx, tx)
		}

		// default to `MaxInt64` when there's no extension option.
		prioPriceCap := sdkmath.NewInt(math.MaxInt64)
		if hasExtOptsTx, ok := tx.(authante.HasExtensionOptionsTx); ok {
			for _, opt := range hasExtOptsTx.GetExtensionOptions() {
				if extOpt, ok := opt.GetCachedValue().(*ethermint.ExtensionOptionDynamicFeeTx); ok {
					prioPriceCap = extOpt.MaxPriorityPrice
					break
				}
			}
		}
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return nil, 0, fmt.Errorf("tx must be a FeeTx")
		}

		gas := feeTx.GetGas()
		feeCoins := feeTx.GetFee()
		fee := feeCoins.AmountOfNoDenomValidation(denom)
		priceCap := fee.Quo(sdkmath.NewIntFromUint64(gas))

		basePrice := sdkmath.NewIntFromBigInt(baseFee)
		if priceCap.LT(basePrice) {
			return nil, 0, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient gas price; got: %s required: %s", priceCap, basePrice)
		}

		effectivePrice := sdkmath.NewIntFromBigInt(types.EffectiveGasPrice(basePrice.BigInt(), priceCap.BigInt(), prioPriceCap.BigInt()))
		effectiveFee := sdk.NewCoins(sdk.NewCoin(denom, effectivePrice.Mul(sdkmath.NewIntFromUint64(gas))))
		bigPriority := effectivePrice.Sub(basePrice).Quo(types.DefaultPriorityReduction)
		priority := int64(math.MaxInt64)
		if bigPriority.IsInt64() {
			priority = bigPriority.Int64()
		}
		return effectiveFee, priority, nil
	}
}

// checkTxFeeWithValidatorMinGasPrices implements the default fee logic, where the minimum price per
// unit of gas is fixed and set by each validator, and the tx priority is computed from the gas price.
func checkTxFeeWithValidatorMinGasPrices(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return nil, 0, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()

	// Ensure that the provided fees meet a minimum threshold for the validator,
	// if this is a CheckTx. This is only for local mempool purposes, and thus
	// is only ran on check tx.
	if ctx.IsCheckTx() {
		minGasPrices := ctx.MinGasPrices()
		if !minGasPrices.IsZero() {
			requiredFees := make(sdk.Coins, len(minGasPrices))

			// Determine the required fees by multiplying each required minimum gas
			// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
			glDec := sdk.NewDec(int64(gas))
			for i, gp := range minGasPrices {
				fee := gp.Amount.Mul(glDec)
				requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
			}

			if !feeCoins.IsAnyGTE(requiredFees) {
				return nil, 0, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, requiredFees)
			}
		}
	}

	priority := getTxPriority(feeCoins)
	return feeCoins, priority, nil
}

// getTxPriority returns a naive tx priority based on the amount of the smallest denomination of the fee
// provided in a transaction.
func getTxPriority(fee sdk.Coins) int64 {
	var priority int64
	for _, c := range fee {
		amt := c.Amount.Quo(types.DefaultPriorityReduction)
		p := int64(math.MaxInt64)
		if amt.IsInt64() {
			p = amt.Int64()
		}
		if priority == 0 || p < priority {
			priority = p
		}
	}

	return priority
}
