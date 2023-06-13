package modules

import (
	"fmt"
	"time"

	"github.com/Moonyongjung/tx-stress-test/handler"
	"github.com/Moonyongjung/tx-stress-test/types"
	"github.com/Moonyongjung/tx-stress-test/util"
	"github.com/mitchellh/mapstructure"

	"github.com/Moonyongjung/xpla.go/client"
	xtypes "github.com/Moonyongjung/xpla.go/types"
	xutil "github.com/Moonyongjung/xpla.go/util"
)

func StartExecuteContract(testerOption types.TesterOption) {
	var val string
	xutil.LogInfo("select function")
	xutil.LogInfo("1. store wasm - cw721 base")
	xutil.LogInfo("2. instantiate")
	xutil.LogInfo("3. execute test - mint token ID")

	fmt.Scan(&val)

	mykey2 := util.GetKeyRecord2().Get()
	xplac, err := handler.MakeXplaClient(mykey2.Mnemonic)
	if err != nil {
		xutil.LogInfo(err)
		return
	}

	switch {
	case val == "1":
		wasmFileName := util.Tester().Get().Wasm.WasmFile
		// Store msg test
		storeMsg := xtypes.StoreMsg{
			FilePath:              wasmFileName,
			InstantiatePermission: "",
		}

		txbytes, err := xplac.StoreCode(storeMsg).CreateAndSignTx()
		if err != nil {
			xutil.LogInfo(err)
			return
		}

		res, err := xplac.BroadcastBlock(txbytes)
		if err != nil {
			xutil.LogInfo(err)
			return

		} else {
			xutil.LogInfo(*res)

			handler.SaveCodeIdAfterStoreContract(res)
		}

	case val == "2":
		initMsg :=
			`{
				"name":"cw721-base-onchain",
				"symbol":"CW721",
				"minter":"` + mykey2.Address + `"
			}`

			// Instantiate msg test

		code, err := getCodeInfo()
		if err != nil {
			xutil.LogInfo(err)
			return
		}
		instantiateMsg := xtypes.InstantiateMsg{
			CodeId:  code.CodeId,
			Amount:  "0",
			Label:   "Contract instant",
			InitMsg: initMsg,
			Admin:   mykey2.Address,
		}
		txbytes, err := xplac.InstantiateContract(instantiateMsg).CreateAndSignTx()
		if err != nil {
			xutil.LogInfo(err)
			return
		}

		res, err := xplac.BroadcastBlock(txbytes)
		if err != nil {
			xutil.LogInfo(err)
			return

		} else {
			xutil.LogInfo(*res)

			handler.SaveContractAddrAfterInitiate(res)
		}

	case val == "3":
		latestTokenID, err := getLatestTokenID()
		if err != nil {
			xutil.LogInfo(err)
			return
		}

		executeTxTest(xplac, testerOption, mykey2.Address, latestTokenID.LatestTokenID)
		handler.SaveTokenIDNumber(latestTokenID.LatestTokenID + types.Txs)

	default:
		xutil.LogInfo("invalid number")
	}

}

func executeTxTest(xplac *client.XplaClient, testerOption types.TesterOption, address string, latestTokenID int) {
	nftTokenId := util.Tester().Get().Wasm.NftTokenId
	amount := "0"

	contractInfo, err := getContractInfo()
	if err != nil {
		xutil.LogInfo(err)
		return
	}

	executeMsg := xtypes.ExecuteMsg{
		ContractAddress: contractInfo.ContractAddress,
		Amount:          amount,
	}

	targetDuration := time.Second * time.Duration(testerOption.TestTime)
	go util.Loading()
	go func() {
		time.Sleep(targetDuration)
		types.TimeCheckChan <- true
	}()

	startTime := time.Now()

Loop1:
	for {
		select {
		case c := <-types.TimeCheckChan:
			types.CloseLoading <- true
			WasmTx(xplac, c, executeMsg, nftTokenId, address)
			break Loop1

		default:
			xplac = WasmTx(xplac, false, executeMsg, nftTokenId, address)
		}
	}
	endTime := time.Now()

	logResult(startTime, endTime, targetDuration)
}

func WasmTx(xplac *client.XplaClient, isTargetTime bool, executeMsg xtypes.ExecuteMsg, nftTokenId, address string) *client.XplaClient {

	execMsg := `
		{
			"mint":{
				"token_id":"` + nftTokenId + "_" + xutil.FromIntToString(types.Txs) + `",
				"owner":"` + address + `",
				"token_uri":"token_uri"
			}
		}`

	executeMsg.ExecMsg = execMsg

	txbytes, _ := xplac.ExecuteContract(executeMsg).CreateAndSignTx()

	if isTargetTime {
		_, err := xplac.BroadcastBlock(txbytes)
		if err != nil {
			types.ErrTxs++
		} else {
			xplac = sequenceMng(xplac)
			types.SuccTxs++
		}
	} else {
		_, err := xplac.Broadcast(txbytes)
		if err != nil {
			types.ErrTxs++
		} else {
			xplac = sequenceMng(xplac)
			types.SuccTxs++
		}
	}
	types.Txs++

	return xplac
}

func getCodeInfo() (handler.ContractCodeJsonStruct, error) {
	var code handler.ContractCodeJsonStruct
	codeIdDir := util.Tester().Get().Wasm.CodeIdDir
	contractCodeData, err := xutil.JsonUnmarshal(code, codeIdDir)
	if err != nil {
		return handler.ContractCodeJsonStruct{}, err
	}
	mapstructure.Decode(contractCodeData, &code)

	return code, nil
}

func getContractInfo() (handler.ContractAddressJsonStruct, error) {
	var c handler.ContractAddressJsonStruct
	contractAddressDir := util.Tester().Get().Wasm.ContractAddrDir
	contractAddressData, err := xutil.JsonUnmarshal(c, contractAddressDir)
	if err != nil {
		return handler.ContractAddressJsonStruct{}, err
	}
	mapstructure.Decode(contractAddressData, &c)

	return c, nil
}

func getLatestTokenID() (handler.TokenIDNumberStruct, error) {
	var c handler.TokenIDNumberStruct
	tokenIdDir := util.Tester().Get().Wasm.SaveTokenIdDir
	tokenIdData, err := xutil.JsonUnmarshal(c, tokenIdDir)
	if err != nil {
		return handler.TokenIDNumberStruct{}, err
	}
	mapstructure.Decode(tokenIdData, &c)

	return c, nil
}
