package nucleotide

import (
	"errors"
	"fmt"
)

var (
	InvalidNucleotideErr = errors.New("invalid nucleotide in the sequence")
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, v := range d {
		if v == 65 {
			h['A'] += 1
		} else if v == 67 {
			h['C'] += 1
		} else if v == 71 {
			h['G'] += 1
		} else if v == 84 {
			h['T'] += 1
		} else {
			return Histogram{}, InvalidNucleotideErr
		}
	}
	return h, nil
}

func Run() {
	dna := DNA("")
	fmt.Println(dna.Counts())
}
