package types

var (
	Txs     = 0
	SuccTxs = 0
	ErrTxs  = 0

	TimeCheckChan = make(chan bool, 1)
	CloseLoading  = make(chan bool, 1)

	TestTxTypeSend     = "send"
	TestTxTypeContract = "contract"
)
