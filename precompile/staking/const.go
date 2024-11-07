package staking

const (
	hexAddress = "0x1000000000000000000000000000000000000002"
	abiFile    = "staking.json"
)

type MethodStaking string

const (
	Delegate        MethodStaking = "delegate"
	BeginRedelegate MethodStaking = "beginRedelegate"
	Undelegate      MethodStaking = "undelegate"
)
