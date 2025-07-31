package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var count int
	for i, v := range cb {
		if i == file {
			for _, x := range v {
				if x {
					count++
				}
			}
			break
		}
	}
	if count < 1 || count > 8 {
		return 0
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var count int
	for _, v := range cb {
		for i, x := range v {
			if i+1 == rank && x {
				count++
			}
		}
	}
	if count < 1 || count > 8 {
		return 0
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int
	for _, v := range cb {
		count += len(v)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
	for _, v := range cb {
		for _, x := range v {
			if x {
				count++
			}
		}
	}
	return count
}
