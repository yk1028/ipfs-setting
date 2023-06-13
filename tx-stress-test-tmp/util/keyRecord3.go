package util

import (
	"sync"

	"github.com/Moonyongjung/tx-stress-test/types"

	xutil "github.com/Moonyongjung/xpla.go/util"
	"github.com/mitchellh/mapstructure"
)

type KeyRecord3 types.KeyRecordStruct

var onceKeyRecord3 sync.Once
var instanceKeyRecord3 *singletonKeyRecord3

type singletonKeyRecord3 struct {
	keyRecord3 KeyRecord3
}

func GetKeyRecord3() *singletonKeyRecord3 {
	onceKeyRecord3.Do(func() {
		instanceKeyRecord3 = &singletonKeyRecord3{}
	})
	return instanceKeyRecord3
}

func (s *singletonKeyRecord3) Get() KeyRecord3 {
	return s.keyRecord3
}

func (s *singletonKeyRecord3) Read(filePath string) {
	keyRecordStructData, _ := xutil.JsonUnmarshal(s.keyRecord3, filePath)
	mapstructure.Decode(keyRecordStructData, &s.keyRecord3)
}
