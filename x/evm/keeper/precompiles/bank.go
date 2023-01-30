package precompiles

import (
	"errors"
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/x/evm/statedb"
	"github.com/evmos/ethermint/x/evm/types"
)

const EVMDenomPrefix = "evm/"

var (
	MintMethod      abi.Method
	BalanceOfMethod abi.Method

	_ statedb.StatefulPrecompiledContract = (*BankContract)(nil)
)

func init() {
	addressType, _ := abi.NewType("address", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	MintMethod = abi.NewMethod(
		"mint", "mint", abi.Function, "", false, false, abi.Arguments{abi.Argument{
			Name: "recipient",
			Type: addressType,
		}, abi.Argument{
			Name: "amount",
			Type: uint256Type,
		}},
		nil,
	)
	BalanceOfMethod = abi.NewMethod(
		"balanceOf", "balanceOf", abi.Function, "", false, false, abi.Arguments{abi.Argument{
			Name: "token",
			Type: addressType,
		}, abi.Argument{
			Name: "address",
			Type: addressType,
		}},
		abi.Arguments{abi.Argument{
			Name: "amount",
			Type: uint256Type,
		}},
	)
}

func EVMDenom(token common.Address) string {
	return EVMDenomPrefix + token.Hex()
}

type BankContract struct {
	ctx        sdk.Context
	bankKeeper types.BankKeeper
}

// NewBankContractCreator creates the precompiled contract to manage native tokens
func NewBankContractCreator(bankKeeper types.BankKeeper) statedb.PrecompiledContractCreator {
	return func(ctx sdk.Context) statedb.StatefulPrecompiledContract {
		return &BankContract{
			ctx:        ctx,
			bankKeeper: bankKeeper,
		}
	}
}

// RequiredGas calculates the contract gas use
func (bc *BankContract) RequiredGas(input []byte) uint64 {
	// TODO estimate required gas
	return 0
}

func (bc *BankContract) Run(evm *vm.EVM, input []byte, caller common.Address, value *big.Int, readonly bool) ([]byte, error) {
	stateDB, ok := evm.StateDB.(ExtStateDB)
	if !ok {
		return nil, errors.New("not run in ethermint")
	}

	// parse input
	methodID := input[:4]
	switch string(methodID) {
	case string(MintMethod.ID):
		if readonly {
			return nil, errors.New("the method is not readonly")
		}
		args, err := MintMethod.Inputs.Unpack(input[4:])
		if err != nil {
			return nil, errors.New("fail to unpack input arguments")
		}
		recipient := args[0].(common.Address)
		amount := args[1].(*big.Int)
		if amount.Sign() <= 0 {
			return nil, errors.New("invalid amount")
		}
		denom := EVMDenom(caller)
		err = stateDB.ExecuteNativeAction(func(ctx sdk.Context) error {
			addr := sdk.AccAddress(recipient.Bytes())
			amt := sdk.NewCoins(sdk.NewCoin(denom, sdkmath.NewIntFromBigInt(amount)))
			if err := bc.bankKeeper.MintCoins(ctx, types.ModuleName, amt); err != nil {
				return sdkerrors.Wrap(err, "fail to mint coins in precompiled contract")
			}
			if err := bc.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, amt); err != nil {
				return sdkerrors.Wrap(err, "fail to send mint coins to account")
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	case string(BalanceOfMethod.ID):
		args, err := BalanceOfMethod.Inputs.Unpack(input[4:])
		if err != nil {
			return nil, errors.New("fail to unpack input arguments")
		}
		token := args[0].(common.Address)
		addr := args[1].(common.Address)
		// query from storage
		amount := bc.bankKeeper.GetBalance(bc.ctx, sdk.AccAddress(addr.Bytes()), EVMDenom(token)).Amount.BigInt()
		return BalanceOfMethod.Outputs.Pack(amount)
	default:
		return nil, errors.New("unknown method")
	}
	return nil, nil
}
