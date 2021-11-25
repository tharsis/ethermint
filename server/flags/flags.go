package flags

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Tendermint full-node start flags
const (
	WithTendermint = "with-tendermint"
	Address        = "address"
	Transport      = "transport"
	TraceStore     = "trace-store"
	CPUProfile     = "cpu-profile"
)

// GRPC-related flags.
const (
	GRPCEnable     = "grpc.enable"
	GRPCAddress    = "grpc.address"
	GRPCWebEnable  = "grpc-web.enable"
	GRPCWebAddress = "grpc-web.address"
)

// JSON-RPC flags
const (
	JSONRPCEnable       = "json-rpc.enable"
	JSONRPCAPI          = "json-rpc.api"
	JSONRPCAddress      = "json-rpc.address"
	JSONWsAddress       = "json-rpc.ws-address"
	JSONRPCGasCap       = "json-rpc.gas-cap"
	JSONRPCEVMTimeout   = "json-rpc.evm-timeout"
	JSONRPCTxFeeCap     = "json-rpc.txfee-cap"
	JSONRPCFilterCap    = "json-rpc.filter-cap"
	JSONRPFeeHistoryCap = "json-rpc.feehistory-cap"
)

// EVM flags
const (
	EVMTracer = "evm.tracer"
)

// TLS flags
const (
	TLSCertPath = "tls.certificate-path"
	TLSKeyPath  = "tls.key-path"
)

// AddTxFlags adds common flags for commands to post tx
func AddTxFlags(cmd *cobra.Command) (*cobra.Command, error) {
	cmd.Flags().String(flags.FlagChainID, "testnet", "Specify Chain ID for sending Tx")
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	cmd.Flags().String(flags.FlagFees, "", "Fees to pay along with transaction; eg: 10aphoton")
	cmd.Flags().String(flags.FlagGasPrices, "", "Gas prices to determine the transaction fee (e.g. 10aphoton)")
	cmd.Flags().String(flags.FlagNode, "tcp://localhost:26657", "<host>:<port> to tendermint rpc interface for this chain")
	cmd.Flags().Float64(flags.FlagGasAdjustment, flags.DefaultGasAdjustment, "adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored ")
	cmd.Flags().StringP(flags.FlagBroadcastMode, "b", flags.BroadcastSync, "Transaction broadcasting mode (sync|async|block)")
	cmd.Flags().String(flags.FlagKeyringBackend, keyring.BackendOS, "Select keyring's backend")

	// --gas can accept integers and "simulate"
	// cmd.PersistentFlags().Var(&flags.GasFlagVar, "gas", fmt.Sprintf(
	//	"gas limit to set per-transaction; set to %q to calculate required gas automatically (default %d)",
	//	flags.GasFlagAuto, flags.DefaultGasLimit,
	// ))

	// viper.BindPFlag(flags.FlagTrustNode, cmd.Flags().Lookup(flags.FlagTrustNode))

	if err := viper.BindPFlag(flags.FlagNode, cmd.Flags().Lookup(flags.FlagNode)); err != nil {
		return nil, err
	}
	if err := viper.BindPFlag(flags.FlagKeyringBackend, cmd.Flags().Lookup(flags.FlagKeyringBackend)); err != nil {
		return nil, err
	}
	if err := cmd.MarkFlagRequired(flags.FlagChainID); err != nil {
		return nil, err
	}
	return cmd, nil
}
