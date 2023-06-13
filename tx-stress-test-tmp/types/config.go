package types

import (
	"errors"
)

type ConfigType struct {
	XplaClientConfig XplaClientConfig `yaml:"xpla_client"`
	TesterOption     TesterOption     `yaml:"tester_option"`
}

type XplaClientConfig struct {
	ChainId      string       `yaml:"chain_id"`
	ClientOption ClientOption `yaml:"client_option"`
}

type ClientOption struct {
	AccountNumber  string `yaml:"account_number"`
	Sequence       string `yaml:"sequence"`
	BroadcastMode  string `yaml:"broadcast_mode"`
	GasLimit       string `yaml:"gas_limit"`
	GasPrice       string `yaml:"gas_price"`
	GasAdj         string `yaml:"gas_adjustment"`
	FeeAmount      string `yaml:"fee_amount"`
	FeeGranter     string `yaml:"fee_granter"`
	TimeoutHeight  string `yaml:"timeout_height"`
	LcdUrl         string `yaml:"lcd_url"`
	GrpcUrl        string `yaml:"grpc_url"`
	RpcUrl         string `yaml:"rpc_url"`
	EvmRpcURL      string `yaml:"evm_rpc_url"`
	OutputDocument string `yaml:"output_document"`
}

type TesterOption struct {
	TestTime int    `yaml:"test_time"`
	TestTx   string `yaml:"test_tx"`
}

func (t TesterOption) TesterConfigValidate() error {
	if !(t.TestTx == TestTxTypeSend || t.TestTx == TestTxTypeContract) {
		return errors.New("invalid tx type, select \"send\" or \"contract\"")
	}

	return nil
}
