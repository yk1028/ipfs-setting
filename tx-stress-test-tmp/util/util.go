package util

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Moonyongjung/tx-stress-test/types"
	"github.com/Moonyongjung/xpla.go/util"
)

func ParsingQueryAccount(res string) string {
	strList := strings.Split(res, "sequence")
	strList = strings.Split(strList[1], "\"")
	sequence := strList[2]

	return sequence
}

func Loading() {
	loadingChars := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}

	i := 0
Loop2:
	for {
		select {
		case <-types.CloseLoading:
			util.LogInfo("---done")
			break Loop2

		default:
			fmt.Printf("\r %s sending tx || try to send tx: %d | send success: %d | send fail: %d",
				loadingChars[i%len(loadingChars)],
				types.Txs,
				types.SuccTxs,
				types.ErrTxs,
			)
			time.Sleep(500 * time.Millisecond)
		}
		i++
	}
}

func JsonMarshal(jsonData interface{}, jsonFilePath string) error {
	byteData, err := util.JsonMarshalData(jsonData)
	if err != nil {
		util.LogInfo(err)
	}
	err = os.WriteFile(jsonFilePath, byteData, os.FileMode(0644))
	if err != nil {
		util.LogInfo(err)
		path := strings.Split(jsonFilePath, "/")
		pathPop := path[:len(path)-1]
		filePath := strings.Join(pathPop, "/")

		err := os.Mkdir(filePath, 0755)
		if err != nil {
			util.LogInfo(err)
		}
		err = os.WriteFile(jsonFilePath, byteData, os.FileMode(0644))
		if err != nil {
			return err
		}
	}

	return nil
}
