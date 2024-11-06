package bank

const (
	hexAddress = "0x1000000000000000000000000000000000000001"
	abiFile    = "bank.json"
)

type MethodBank string

const (
	Balance MethodBank = "balance"
	Send    MethodBank = "send"
	Supply  MethodBank = "supplyOf"
)
