package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File 
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	_, exists := cb[file]
    if !exists {
        return 0
    }
    cntr := 0
    for i := 0; i < len(cb[file]); i++ {
        if cb[file][i] == true {
            cntr++
        }
    }
    return cntr
    // panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
        return 0
    }
    cntr := 0
    for _, line := range cb {
        if line[rank-1] {
            cntr++
        }
    }
    return cntr
    // panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	cntr := 0
    for _, line := range cb {
        for _, _ = range line {
            cntr++
        } 
    }
    return cntr
    // panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	cntr := 0
    for _, line := range cb {
        for _, i := range line {
            if i {
                cntr++
            }
        }
    }
    return cntr
    // panic("Please implement CountOccupied()")
}
