package modules

import (
	"time"

	"github.com/Moonyongjung/tx-stress-test/handler"
	"github.com/Moonyongjung/tx-stress-test/types"
	"github.com/Moonyongjung/tx-stress-test/util"

	"github.com/Moonyongjung/xpla.go/client"
	xtypes "github.com/Moonyongjung/xpla.go/types"
	xutil "github.com/Moonyongjung/xpla.go/util"
)

func StartSendTx(testerOption types.TesterOption) {
	mykey2 := util.GetKeyRecord2().Get()
	mykey3 := util.GetKeyRecord3().Get()
	xplac, err := handler.MakeXplaClient(mykey2.Mnemonic)
	if err != nil {
		xutil.LogInfo(err)
		return
	}

	bankSendMsg := xtypes.BankSendMsg{
		FromAddress: mykey2.Address,
		ToAddress:   mykey3.Address,
		Amount:      "1",
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
			BankTx(xplac, c, bankSendMsg)
			break Loop1

		default:
			xplac = BankTx(xplac, false, bankSendMsg)
		}
	}
	endTime := time.Now()

	logResult(startTime, endTime, targetDuration)
}

func BankTx(xplac *client.XplaClient, isTargetTime bool, bankSendMsg xtypes.BankSendMsg) *client.XplaClient {
	txbytes, _ := xplac.BankSend(bankSendMsg).CreateAndSignTx()

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
