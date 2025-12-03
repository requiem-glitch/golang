package isogram

import "strings"

func IsIsogram(word string) bool {
	nWord := strings.ToLower(word)
    for _, r := range nWord {
        cnt := 0
        for _, l := range nWord {
            if r == l  && r != '-' && r != ' '{
                cnt ++
            }
        }
        if cnt >= 2 {
            return false
        }
    }
    return true
    // panic("Please implement the IsIsogram function")
}
