package types

type TesterType struct {
	Wasm TesterWasm `yaml:"wasm"`
}

type TesterWasm struct {
	WasmFile        string `yaml:"wasm_file"`
	CodeIdDir       string `yaml:"code_id_dir"`
	ContractAddrDir string `yaml:"contract_addr_dir"`
	NftTokenId      string `yaml:"nft_token_id"`
	SaveTokenIdDir  string `yaml:"save_token_id_dir"`
}
