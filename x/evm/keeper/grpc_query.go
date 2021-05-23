package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ethcmn "github.com/ethereum/go-ethereum/common"

	ethermint "github.com/cosmos/ethermint/types"
	"github.com/cosmos/ethermint/x/evm/types"
)

var _ types.QueryServer = Keeper{}

// Account implements the Query/Account gRPC method
func (k Keeper) Account(c context.Context, req *types.QueryAccountRequest) (*types.QueryAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := ethermint.ValidateAddress(req.Address); err != nil {
		return nil, status.Error(
			codes.InvalidArgument, err.Error(),
		)
	}

	addr := ethcmn.HexToAddress(req.Address)

	ctx := sdk.UnwrapSDKContext(c)
	k.WithContext(ctx)

	return &types.QueryAccountResponse{
		Balance:  k.GetBalance(addr).String(),
		CodeHash: k.GetCodeHash(addr).Bytes(), // TODO: use hex
		Nonce:    k.GetNonce(addr),
	}, nil
}

func (k Keeper) CosmosAccount(c context.Context, req *types.QueryCosmosAccountRequest) (*types.QueryCosmosAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := ethermint.ValidateAddress(req.Address); err != nil {
		return nil, status.Error(
			codes.InvalidArgument, err.Error(),
		)
	}

	ctx := sdk.UnwrapSDKContext(c)
	k.WithContext(ctx)

	ethAddr := ethcmn.HexToAddress(req.Address)
	cosmosAddr := sdk.AccAddress(ethAddr.Bytes())

	account := k.accountKeeper.GetAccount(ctx, cosmosAddr)
	res := types.QueryCosmosAccountResponse{
		CosmosAddress: cosmosAddr.String(),
	}

	if account != nil {
		res.Sequence = account.GetSequence()
		res.AccountNumber = account.GetAccountNumber()
	}

	return &res, nil
}

// Balance implements the Query/Balance gRPC method
func (k Keeper) Balance(c context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := ethermint.ValidateAddress(req.Address); err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			types.ErrZeroAddress.Error(),
		)
	}

	ctx := sdk.UnwrapSDKContext(c)
	k.WithContext(ctx)

	balanceInt := k.GetBalance(ethcmn.HexToAddress(req.Address))
	balance, err := ethermint.MarshalBigInt(balanceInt)
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			"failed to marshal big.Int to string",
		)
	}

	return &types.QueryBalanceResponse{
		Balance: balance,
	}, nil
}

// Storage implements the Query/Storage gRPC method
func (k Keeper) Storage(c context.Context, req *types.QueryStorageRequest) (*types.QueryStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := ethermint.ValidateAddress(req.Address); err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			types.ErrZeroAddress.Error(),
		)
	}

	if ethermint.IsEmptyHash(req.Key) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			types.ErrEmptyHash.Error(),
		)
	}

	ctx := sdk.UnwrapSDKContext(c)
	k.WithContext(ctx)

	address := ethcmn.HexToAddress(req.Address)
	key := ethcmn.HexToHash(req.Key)

	state := k.GetState(address, key)
	stateHex := state.Hex()

	if ethermint.IsEmptyHash(stateHex) {
		return nil, status.Error(
			codes.NotFound, "contract code not found for given address",
		)
	}

	return &types.QueryStorageResponse{
		Value: stateHex,
	}, nil
}

// Code implements the Query/Code gRPC method
func (k Keeper) Code(c context.Context, req *types.QueryCodeRequest) (*types.QueryCodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := ethermint.ValidateAddress(req.Address); err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			types.ErrZeroAddress.Error(),
		)
	}

	ctx := sdk.UnwrapSDKContext(c)
	k.WithContext(ctx)

	address := ethcmn.HexToAddress(req.Address)
	code := k.GetCode(address)

	return &types.QueryCodeResponse{
		Code: code,
	}, nil
}

// TxLogs implements the Query/TxLogs gRPC method
func (k Keeper) TxLogs(c context.Context, req *types.QueryTxLogsRequest) (*types.QueryTxLogsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if ethermint.IsEmptyHash(req.Hash) {
		return nil, status.Error(
			codes.InvalidArgument,
			types.ErrEmptyHash.Error(),
		)
	}

	ctx := sdk.UnwrapSDKContext(c)
	k.WithContext(ctx)

	hash := ethcmn.HexToHash(req.Hash)
	logs := k.GetTxLogs(hash)

	return &types.QueryTxLogsResponse{
		Logs: types.NewTransactionLogsFromEth(hash, logs).Logs,
	}, nil
}

// TxReceipt implements the Query/TxReceipt gRPC method
func (k Keeper) TxReceipt(c context.Context, req *types.QueryTxReceiptRequest) (*types.QueryTxReceiptResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if ethermint.IsEmptyHash(req.Hash) {
		return nil, status.Error(
			codes.InvalidArgument,
			types.ErrEmptyHash.Error(),
		)
	}

	ctx := sdk.UnwrapSDKContext(c)

	hash := ethcmn.HexToHash(req.Hash)
	receipt, found := k.GetTxReceiptFromHash(ctx, hash)
	if !found {
		return nil, status.Errorf(
			codes.NotFound, "%s: %s", types.ErrTxReceiptNotFound.Error(), req.Hash,
		)
	}

	return &types.QueryTxReceiptResponse{
		Receipt: receipt,
	}, nil
}

// TxReceiptsByBlockHeight implements the Query/TxReceiptsByBlockHeight gRPC method
func (k Keeper) TxReceiptsByBlockHeight(c context.Context, _ *types.QueryTxReceiptsByBlockHeightRequest) (*types.QueryTxReceiptsByBlockHeightResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	receipts := k.GetTxReceiptsByBlockHeight(ctx, ctx.BlockHeight())
	return &types.QueryTxReceiptsByBlockHeightResponse{
		Receipts: receipts,
	}, nil
}

// TxReceiptsByBlockHash implements the Query/TxReceiptsByBlockHash gRPC method
func (k Keeper) TxReceiptsByBlockHash(c context.Context, req *types.QueryTxReceiptsByBlockHashRequest) (*types.QueryTxReceiptsByBlockHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if ethermint.IsEmptyHash(req.Hash) {
		return nil, status.Error(
			codes.InvalidArgument,
			types.ErrEmptyHash.Error(),
		)
	}

	ctx := sdk.UnwrapSDKContext(c)

	hash := ethcmn.HexToHash(req.Hash)
	receipts := k.GetTxReceiptsByBlockHash(ctx, hash)

	return &types.QueryTxReceiptsByBlockHashResponse{
		Receipts: receipts,
	}, nil
}

// BlockLogs implements the Query/BlockLogs gRPC method
func (k Keeper) BlockLogs(c context.Context, req *types.QueryBlockLogsRequest) (*types.QueryBlockLogsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if ethermint.IsEmptyHash(req.Hash) {
		return nil, status.Error(
			codes.InvalidArgument,
			types.ErrEmptyHash.Error(),
		)
	}

	ctx := sdk.UnwrapSDKContext(c)

	txLogs := k.GetAllTxLogs(ctx)

	return &types.QueryBlockLogsResponse{
		TxLogs: txLogs,
	}, nil
}

// BlockBloom implements the Query/BlockBloom gRPC method
func (k Keeper) BlockBloom(c context.Context, _ *types.QueryBlockBloomRequest) (*types.QueryBlockBloomResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	bloom, found := k.GetBlockBloom(ctx, ctx.BlockHeight())
	if !found {
		return nil, status.Error(
			codes.NotFound, types.ErrBloomNotFound.Error(),
		)
	}

	return &types.QueryBlockBloomResponse{
		Bloom: bloom.Bytes(),
	}, nil
}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}

// StaticCall implements Query/StaticCall gRPCP method
func (k Keeper) StaticCall(c context.Context, req *types.QueryStaticCallRequest) (*types.QueryStaticCallResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	return nil, nil

	// ctx := sdk.UnwrapSDKContext(c)
	// k.WithContext(ctx)

	// var recipient *ethcmn.Address
	// if len(req.Address) > 0 {
	// 	addr := ethcmn.HexToAddress(req.Address)
	// 	recipient = &addr
	// }

	// // TODO: why is this hardcoded
	// sender := ethcmn.HexToAddress("0xaDd00275E3d9d213654Ce5223f0FADE8b106b707")

	// msg := types.NewMsgEthereumTx(
	// 	k.eip155ChainID, 0, recipient, big.NewInt(0), 100000000, big.NewInt(0), req.Input, nil, // TODO: get account nonce?
	// )
	// msg.From = sender.Hex()

	// if err := msg.ValidateBasic(); err != nil {
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	// ethMsg, err := msg.AsMessage()
	// if err != nil {
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	// cfg, found := k.GetChainConfig(ctx)
	// if !found {
	// 	return nil, types.ErrChainConfigNotFound
	// }

	// gasLimit := uint64(0) // TODO: define
	// evm := k.NewEVM(ethMsg, cfg.EthereumConfig(k.eip155ChainID), gasLimit)

	// evm.StaticCall()

	// return &types.QueryStaticCallResponse{Data: ret}, nil
}
