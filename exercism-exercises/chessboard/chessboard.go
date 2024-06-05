package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	value, exists := cb[rank]
	if !exists {
		return 0
	} else {
		var res int = 0
		for i := 0; i < len(value); i++ {
			if value[i] {
				res += 1
			}
		}
		return res
	}
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	file -= 1
	if file < 0 || file > 7 {
		return 0
	}
	var res int = 0
	for _, ranks := range cb {
		if ranks[file] {
			res += 1
		}
	}
	return res
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	columns, _ := cb["A"]
	return len(cb) * len(columns)
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var result int = 0
	for key := range cb {
		result += CountInRank(cb, key)
	}
	return result
}

func Run() {
	var chessboard Chessboard = map[string]Rank{
		"A": []bool{true, false, true, false, true, false, false, false},
		"B": []bool{true, false, true, false, true, false, false, false},
		"C": []bool{false, false, false, false, true, false, false, false},
		"D": []bool{false, false, true, false, true, false, false, false},
		"E": []bool{false, false, true, false, true, false, false, false},
		"F": []bool{true, false, true, false, false, false, false, false},
		"G": []bool{false, false, true, false, true, false, false, false},
		"H": []bool{true, false, false, false, true, false, false, true},
	}
	println("Count in rank:", CountInRank(chessboard, "A"))
	println("Count in file:", CountInFile(chessboard, 1))
	println("Count all:", CountAll(chessboard))
	println("Count occupied:", CountOccupied(chessboard))
}
