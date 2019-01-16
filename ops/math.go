package ops

// type Summer interface {
// 	Sum() int
// }

// type Complex struct {
// 	Real      int
// 	Imaginary int
// }

// func (c *Complex) Sum() int {
// 	return c.Real + c.Imaginary
// }

// func (c *Complex) String() string {
// 	return fmt.Sprintf("%d + i%d", c.Real, c.Imaginary)
// }

// func NewComplex(real, imag int) *Complex {
// 	cpx := Complex{
// 		Real:      real,
// 		Imaginary: imag,
// 	}
// 	return &cpx
// }

// //PerformSumDiff - performs sum and difference from given parameters. Gives
// //error if param1 < param2
// func PerformSumDiff(param1, param2 int) (sum int, diff int, err error) {
// 	sum = param1 + param2
// 	if param1 >= param2 {
// 		diff = param1 - param2
// 	} else {
// 		err = errors.New("Invalid parameters")
// 	}
// 	return
// }

// func PrintSumOf(summer Summer) {
// 	fmt.Printf("%d\n", summer.Sum())
// }
