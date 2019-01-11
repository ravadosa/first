package ops

import "errors"

//PerformSumDiff - performs sum and difference from given parameters. Gives
//error if param1 < param2
func PerformSumDiff(param1, param2 int) (sum int, diff int, err error) {
	sum = param1 + param2
	if param1 >= param2 {
		diff = param1 - param2
	} else {
		err = errors.New("Invalid parameters")
	}
	return
}
