package account

var fundsByCodes = map[string]int32{
	"double": 2,
	"triple": 3,
}

const fundsByCodeDefault = 1

func getFundsByCode(code string) int32 {
	funds, found := fundsByCodes[code]
	if !found {
		return fundsByCodeDefault
	}

	return funds
}
