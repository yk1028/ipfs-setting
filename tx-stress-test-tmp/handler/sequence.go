package handler

import (
	"sync"

	xutil "github.com/Moonyongjung/xpla.go/util"
)

var SequenceInstance *SequenceStruct
var SequenceOnce sync.Once

// manage account Sequence
type SequenceStruct struct {
	Sequence string
}

func SequenceMng() *SequenceStruct {
	SequenceOnce.Do(func() {
		SequenceInstance = &SequenceStruct{}
	})
	return SequenceInstance
}

func (n *SequenceStruct) NewSequence(sequence string) {
	n.Sequence = sequence
}

func (n *SequenceStruct) NowSequence() string {
	return n.Sequence
}

func (n *SequenceStruct) AddSequence() {
	Sequence := n.Sequence
	SequenceNum := xutil.FromStringToUint64(Sequence)
	SequenceNum = SequenceNum + 1
	Sequence = xutil.FromUint64ToString(SequenceNum)
	n.Sequence = Sequence
}
