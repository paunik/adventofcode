package two

const (
	step = 4
)

// AlarmOpcode calculates op code
func AlarmOpcode(ocis []int) []int {

	count := len(ocis)

	for index := 0; index < count-4; index++ {

		if index%step == 0 {
			operation := ocis[index]
			firstParamAddress := ocis[index+1]
			secondParamAddress := ocis[index+2]
			storageParamAddress := ocis[index+3]

			switch operation {
			case 1:
				ocis[storageParamAddress] = ocis[firstParamAddress] + ocis[secondParamAddress]
			case 2:
				ocis[storageParamAddress] = ocis[firstParamAddress] * ocis[secondParamAddress]
			case 99:
				return ocis
			}
		}
	}
	return ocis
}
