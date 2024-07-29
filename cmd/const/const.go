package constx

// '状态 PENDINxG:待处理; PENDING_CHECK:待check; SUCCESS:处理成功; FAILED：回滚;',
const (
	Version = "2.0.1"

	ChainBNB   = "bnb"
	Multiplier = 9
	MUSDT      = 18

	EventTransfer = "Transfer"

	ChainBlockUnSafe           = false
	ChainBlockSafe             = true
	ChainBlockNotSyncTx        = false
	ChainBlockAlreadySyncTx    = true
	ChainBlockNotConfirmTx     = false
	ChainBlockAlreadyConfirmTx = true

	ChainTxPendingCheck = "PENDING_CHECK"
	ChainTxSuccess      = "SUCCESS"
	ChainTxFailed       = "FAILED"

	ExecuteStatusPending = 0
	ExecuteStatusSuccess = 1

	N10E8 = 1000000000000000000

	TrueInt  = 1
	FALSEInt = 2

	USTDFlag = 1
	P3Flag   = 2
)

// todo !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//
//		The private key is compiled into the binary,    !!!!
//		It's important, don't reveal it to anyone    !!!!!!!
//	 Go and change it to the private key you need  !!!!!!
//		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
const PrivateKeyConst = ""
