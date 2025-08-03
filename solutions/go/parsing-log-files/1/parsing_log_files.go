package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	if err != nil {
		panic(err)
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<(~|\*|=|-)*>`)
	if err != nil {
		panic(err)
	}
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re, err := regexp.Compile(`".*password.*"`)
	if err != nil {
		panic(err)
	}
	for _, v := range lines {
		if re.MatchString(strings.ToLower(v)) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d+`)
	if err != nil {
		panic(err)
	}
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User +(\w*) `)
	if err != nil {
		panic(err)
	}
	for i, v := range lines {
		if username := re.FindAllStringSubmatch(v, -1); username != nil {
			lines[i] = fmt.Sprintf("[USR] %s ", username[0][1]) + lines[i]
		}
	}
	return lines
}
