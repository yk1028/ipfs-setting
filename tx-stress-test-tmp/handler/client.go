package handler

import (
	"fmt"

	"github.com/Moonyongjung/tx-stress-test/util"

	"github.com/Moonyongjung/xpla.go/client"
	"github.com/Moonyongjung/xpla.go/key"
	"github.com/Moonyongjung/xpla.go/types"
	xutil "github.com/Moonyongjung/xpla.go/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MakeXplaClient(mnemonic string) (*client.XplaClient, error) {
	priKey, err := key.NewPrivKey(mnemonic)
	if err != nil {
		fmt.Println(err)
	}

	addr, err := key.Bech32AddrString(priKey)
	if err != nil {
		return nil, err
	}

	xutil.LogInfo("sender:", addr)

	chainId := util.Config().Get().XplaClientConfig.ChainId
	clientOption := util.Config().Get().XplaClientConfig.ClientOption

	var feeGranter sdk.AccAddress
	if clientOption.FeeGranter != "" {
		feeGranter, err = sdk.AccAddressFromBech32(clientOption.FeeGranter)
		if err != nil {
			return nil, err
		}
	}

	newClientOption := client.Options{
		PrivateKey:     priKey,
		AccountNumber:  clientOption.AccountNumber,
		Sequence:       clientOption.Sequence,
		BroadcastMode:  clientOption.BroadcastMode,
		GasLimit:       clientOption.GasLimit,
		GasPrice:       clientOption.GasPrice,
		GasAdjustment:  clientOption.GasAdj,
		FeeAmount:      clientOption.FeeAmount,
		FeeGranter:     feeGranter,
		TimeoutHeight:  clientOption.TimeoutHeight,
		LcdURL:         clientOption.LcdUrl,
		GrpcURL:        clientOption.GrpcUrl,
		RpcURL:         clientOption.RpcUrl,
		EvmRpcURL:      clientOption.EvmRpcURL,
		OutputDocument: clientOption.OutputDocument,
	}

	xplac := client.NewXplaClient(chainId).WithOptions(newClientOption)

	res, err := querySequence(xplac, addr)
	if err != nil {
		return nil, err
	}

	sequence := util.ParsingQueryAccount(res)
	SequenceMng().NewSequence(sequence)

	xplac.WithSequence(sequence)

	return xplac, nil
}

func querySequence(xplac *client.XplaClient, addr string) (string, error) {
	queryAccAddressMsg := types.QueryAccAddressMsg{
		Address: addr,
	}
	res, err := xplac.AccAddress(queryAccAddressMsg).Query()
	if err != nil {
		return "", err
	}

	return res, nil
}
