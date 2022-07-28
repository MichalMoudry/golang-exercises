package main

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 || int(a+b+c)%3 == 0 {
		return NaT
	}
	var k Kind
	if a == b && b == c {
		k = Equ
	} else if (a == b && c != a) || (a == c && b != a) || (b == c && a != b) {
		k = Iso
	} else if a != b && b != c {
		k = Sca
	} else {
		k = NaT
	}
	return k
}

func main() {
	println("Triangle:", KindFromSides(5, 5, 5))
	println("Triangle:", KindFromSides(5, 6, 6))
	println("Triangle:", KindFromSides(5, 6, 7))
}
