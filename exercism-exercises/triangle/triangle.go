package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// Method for validating input.
func isTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	} else if ((a + b) < c) || ((b + c) < a) || ((a + c) < b) {
		return false
	}
	return true
}

// Function determining triangle's kind based on length of its sides.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	var k Kind
	if a == b && b == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}

func Run() {
	println("Triangle:", KindFromSides(1, 1, 5))
	println("Triangle:", KindFromSides(1, 1, 1))
	println("Triangle:", KindFromSides(1, 6, 1))
}
