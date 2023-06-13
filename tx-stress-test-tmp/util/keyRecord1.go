package util

import (
	"sync"

	"github.com/Moonyongjung/tx-stress-test/types"

	xutil "github.com/Moonyongjung/xpla.go/util"
	"github.com/mitchellh/mapstructure"
)

type KeyRecord1 types.KeyRecordStruct

var onceKeyRecord1 sync.Once
var instanceKeyRecord1 *singletonKeyRecord1

type singletonKeyRecord1 struct {
	keyRecord1 KeyRecord1
}

func GetKeyRecord1() *singletonKeyRecord1 {
	onceKeyRecord1.Do(func() {
		instanceKeyRecord1 = &singletonKeyRecord1{}
	})
	return instanceKeyRecord1
}

func (s *singletonKeyRecord1) Get() KeyRecord1 {
	return s.keyRecord1
}

func (s *singletonKeyRecord1) Read(filePath string) {
	keyRecordStructData, _ := xutil.JsonUnmarshal(s.keyRecord1, filePath)
	mapstructure.Decode(keyRecordStructData, &s.keyRecord1)
}
