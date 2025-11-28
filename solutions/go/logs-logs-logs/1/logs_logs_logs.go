package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, i := range log {
        if i == '❗' {
            return "recommendation"
        } else if i == '🔍' {
            return "search"
        } else if i == '☀' {
            return "weather"
        }
        
    }
    return "default"
    // panic("Please implement the Application() function")
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) (result string) {
    for _, char := range log {
        if oldRune == char {
            result += string(newRune) 
        } else {
            result += string(char)
        }
    }
    return
    // panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	for _, _ = range log {
        limit--
    }
    if limit >= 0 {
        return true
    }
    return false
    // panic("Please implement the WithinLimit() function")
}
