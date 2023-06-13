package util

import (
	"sync"

	"github.com/Moonyongjung/tx-stress-test/types"
	xutil "github.com/Moonyongjung/xpla.go/util"

	"github.com/mitchellh/mapstructure"
)

type KeyRecord2 types.KeyRecordStruct

var onceKeyRecord2 sync.Once
var instanceKeyRecord2 *singletonKeyRecord2

type singletonKeyRecord2 struct {
	keyRecord2 KeyRecord2
}

func GetKeyRecord2() *singletonKeyRecord2 {
	onceKeyRecord2.Do(func() {
		instanceKeyRecord2 = &singletonKeyRecord2{}
	})
	return instanceKeyRecord2
}

func (s *singletonKeyRecord2) Get() KeyRecord2 {
	return s.keyRecord2
}

func (s *singletonKeyRecord2) Read(filePath string) {
	keyRecordStructData, _ := xutil.JsonUnmarshal(s.keyRecord2, filePath)
	mapstructure.Decode(keyRecordStructData, &s.keyRecord2)
}
