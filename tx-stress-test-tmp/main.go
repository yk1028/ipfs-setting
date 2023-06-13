package main

import (
	"github.com/Moonyongjung/tx-stress-test/modules"
	"github.com/Moonyongjung/tx-stress-test/types"
	"github.com/Moonyongjung/tx-stress-test/util"

	xutil "github.com/Moonyongjung/xpla.go/util"
)

var configPath = "./config.yaml"
var testerPath = "./tester/tester.yaml"
var keyRecord1Path = "./keylist/keyRecord1.json"
var keyRecord2Path = "./keylist/keyRecord2.json"
var keyRecord3Path = "./keylist/keyRecord3.json"

func init() {
	util.Config().Read(configPath)
	util.Tester().Read(testerPath)
	util.GetKeyRecord1().Read(keyRecord1Path)
	util.GetKeyRecord2().Read(keyRecord2Path)
	util.GetKeyRecord3().Read(keyRecord3Path)
}

func main() {
	testerOption := util.Config().Get().TesterOption

	if err := testerOption.TesterConfigValidate(); err != nil {
		xutil.LogInfo(err)
		return
	}

	switch {
	case testerOption.TestTx == types.TestTxTypeSend:
		modules.StartSendTx(testerOption)

	case testerOption.TestTx == types.TestTxTypeContract:
		modules.StartExecuteContract(testerOption)
	}
}
