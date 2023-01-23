package v5

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/evm/types"

	v5types "github.com/evmos/ethermint/x/evm/migrations/v5/types"
)

var (
	ParamStoreKeyEVMDenom            = []byte("EVMDenom")
	ParamStoreKeyEnableCreate        = []byte("EnableCreate")
	ParamStoreKeyEnableCall          = []byte("EnableCall")
	ParamStoreKeyExtraEIPs           = []byte("EnableExtraEIPs")
	ParamStoreKeyChainConfig         = []byte("ChainConfig")
	ParamStoreKeyAllowUnprotectedTxs = []byte("AllowUnprotectedTxs")
)

// MigrateStore migrates the x/evm module state from the consensus version 4 to
// version 5. Specifically, it takes the parameters that are currently stored
// in separate keys and stores them directly into the x/evm module state using
// a single params key.
func MigrateStore(
	ctx sdk.Context,
	storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
) error {
	var (
		extraEIPs   v5types.V5ExtraEIPs
		chainConfig types.ChainConfig
		params      types.Params
	)

	store := ctx.KVStore(storeKey)

	denom := string(store.Get(ParamStoreKeyEVMDenom))

	extraEIPsBz := store.Get(ParamStoreKeyExtraEIPs)
	cdc.MustUnmarshal(extraEIPsBz, &extraEIPs)

	chainCfgBz := store.Get(ParamStoreKeyChainConfig)
	cdc.MustUnmarshal(chainCfgBz, &chainConfig)

	params.EvmDenom = denom
	params.ExtraEIPs = extraEIPs.EIPs
	params.ChainConfig = chainConfig
	params.EnableCreate = store.Has(ParamStoreKeyEnableCreate)
	params.EnableCall = store.Has(ParamStoreKeyEnableCall)
	params.AllowUnprotectedTxs = store.Has(ParamStoreKeyAllowUnprotectedTxs)

	store.Delete(ParamStoreKeyChainConfig)
	store.Delete(ParamStoreKeyExtraEIPs)
	store.Delete(ParamStoreKeyEVMDenom)
	store.Delete(ParamStoreKeyEnableCreate)
	store.Delete(ParamStoreKeyEnableCall)
	store.Delete(ParamStoreKeyAllowUnprotectedTxs)

	if err := params.Validate(); err != nil {
		return err
	}

	bz := cdc.MustMarshal(&params)

	store.Set(types.KeyPrefixParams, bz)
	return nil
}
