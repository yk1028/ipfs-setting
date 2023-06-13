package modules

import (
	"time"

	"github.com/Moonyongjung/tx-stress-test/handler"
	"github.com/Moonyongjung/tx-stress-test/types"

	"github.com/Moonyongjung/xpla.go/client"
	xutil "github.com/Moonyongjung/xpla.go/util"
)

func sequenceMng(xplac *client.XplaClient) *client.XplaClient {
	handler.SequenceMng().AddSequence()
	xplac.WithSequence(handler.SequenceMng().NowSequence())

	return xplac
}

func logResult(startTime, endTime time.Time, targetDuration time.Duration) {
	xutil.LogInfo("\nstart time:", startTime)
	xutil.LogInfo("end time:", endTime)
	xutil.LogInfo("target duration:", targetDuration)
	xutil.LogInfo("total duration(waiting latest tx's response - in the block):", endTime.Sub(startTime))
	xutil.LogInfo("total try to send txs:", types.Txs)
	xutil.LogInfo("success sended txs:", types.SuccTxs)
	xutil.LogInfo("failed sended txs:", types.ErrTxs)
}
