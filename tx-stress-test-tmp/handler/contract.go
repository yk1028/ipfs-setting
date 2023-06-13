package handler

import (
	"github.com/Moonyongjung/tx-stress-test/util"
	"github.com/Moonyongjung/xpla.go/types"
	xutil "github.com/Moonyongjung/xpla.go/util"
)

type ContractCodeJsonStruct struct {
	TypeName string
	CodeId   string
}

type ContractAddressJsonStruct struct {
	TypeName        string
	ContractAddress string
}

type TokenIDNumberStruct struct {
	LatestTokenID int
}

func SaveCodeIdAfterStoreContract(res *types.TxRes) {
	cwRes := res.Response

	key := cwRes.Logs[0].Events[1].Attributes[0].Key
	value := cwRes.Logs[0].Events[1].Attributes[0].Value

	jsonData := ContractCodeJsonStruct{
		TypeName: key,
		CodeId:   value,
	}

	f := util.Tester().Get().Wasm.CodeIdDir

	err := util.JsonMarshal(jsonData, f)
	if err != nil {
		xutil.LogInfo(err)
	}

	SaveTokenIDNumber(0)
}

func SaveContractAddrAfterInitiate(res *types.TxRes) {
	cwRes := res.Response

	typeName := cwRes.Logs[0].Events[0].Attributes[0].Key
	contractAddress := cwRes.Logs[0].Events[0].Attributes[0].Value

	jsonData := ContractAddressJsonStruct{
		TypeName:        typeName,
		ContractAddress: contractAddress,
	}

	f := util.Tester().Get().Wasm.ContractAddrDir

	err := util.JsonMarshal(jsonData, f)
	if err != nil {
		xutil.LogInfo(err)
	}
}

func SaveTokenIDNumber(tokenID int) {
	jsonData := TokenIDNumberStruct{
		LatestTokenID: tokenID,
	}

	f := util.Tester().Get().Wasm.SaveTokenIdDir

	err := util.JsonMarshal(jsonData, f)
	if err != nil {
		xutil.LogInfo(err)
	}
}
