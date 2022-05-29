package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	r := regexp.MustCompile(`^((\[TRC\])|(\[DBG\])|(\[INF\])|(\[WRN\])|(\[ERR\])|(\[FTL\])).*$`)
	return r.MatchString(text)
}

func SplitLogLine(text string) []string {
	r := regexp.MustCompile(`<(\*|~|=|-)*>`)
	return r.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (counter int) {
	r := regexp.MustCompile(`(?i)` + `^.*".*password.*".*$`)
	for _, line := range lines {
		if r.MatchString(line) {
			counter++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile(`end-of-line[0-9]+`)
	return r.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	retArr := make([]string, len(lines))
	rUser := regexp.MustCompile(`.*User\s+.*`)
	rUserSplit := regexp.MustCompile(`\s+User\s+`)
	rUsername := regexp.MustCompile(`^\S+\s`)
	for i, line := range lines {
		if rUser.MatchString(line) {
			splices := rUserSplit.Split(line, -1)
			username := rUsername.FindString(splices[1])
			retArr[i] = "[USR] " + username + line
		} else {
			retArr[i] = line
		}
	}
	return retArr
}
