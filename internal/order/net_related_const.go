package order

const TheOnlyNet = "ARBI_USDT"
const TheOnlyAddr = "0x09197c3dd57E86Cb8b02A7ca2c315a7e59dE9383"

func getAddrByNet(net string) (addr string, isFound bool, minAmount int32) {
	if net == TheOnlyNet {
		return TheOnlyAddr, true, 1e6
	}

	return "", false, 0
}
