package parsinglogfiles
import (
    "regexp"
    "fmt"
    )
func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
    // panic("Please implement the IsValidLine function")
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~=*]*>`)
    return re.Split(text, -1)
    // panic("Please implement the SplitLogLine function")
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"(?i)[^"]*password[^"]*"`)
    count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
    // panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
    // panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	result := make([]string, len(lines))
    for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			username := matches[1]
			result[i] = fmt.Sprintf("[USR] %s %s", username, line)
		} else {
			result[i] = line
		}
	}
    return result
    // panic("Please implement the TagWithUserName function")
}
