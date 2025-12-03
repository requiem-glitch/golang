package scrabble

import "strings"

func Score(word string) (sum int) {
	WORD := strings.ToUpper(word)
    for _, r := range WORD {
        if r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'L' || r == 'N' || r == 'R' || r == 'S' || r == 'T' {
            sum++
        } else if r == 'D' || r == 'G' {
            sum += 2
        } else if r == 'B' || r == 'C' || r == 'M' || r == 'P' {
            sum += 3
        } else if r == 'F' || r == 'H' || r == 'V' || r == 'W' || r == 'Y' {
            sum += 4
        } else if r == 'K' {
            sum += 5
        } else if r == 'J' || r == 'X' {
            sum += 8
        } else if r == 'Q' || r == 'Z' {
            sum += 10
        }
    }
    return
    
    // panic("Please implement the Score function")
}
