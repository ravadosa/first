package ops

type ComplexNumber struct {
	Real      int
	Imaginary int
}

func (this *ComplexNumber) Add() (sum int) {
	sum = this.Real + this.Imaginary
	return
}
