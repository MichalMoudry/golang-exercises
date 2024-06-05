package logparser

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	validLineRegex, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	foundString := validLineRegex.FindString(text)
	return foundString != ""
}

func SplitLogLine(text string) []string {
	splitRegex, _ := regexp.Compile(`<(\*|~|=|-)*>`)
	return splitRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	passwordRegex, _ := regexp.Compile("\".*[pP][aA][sS]{2}[wW][oO][rR][dD].*\"")
	var tempFind string
	var result int = 0
	for i := 0; i < len(lines); i++ {
		tempFind = passwordRegex.FindString(lines[i])
		if tempFind != "" {
			result++
		}
	}
	return result
}

func RemoveEndOfLineText(text string) string {
	eolRegex, _ := regexp.Compile("end-of-line\\d*")
	return eolRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	tagRegex, _ := regexp.Compile("User\\s+\\w*")
	usernameRegex, _ := regexp.Compile("\\w*$")
	var submatch []string
	var result []string
	for i := 0; i < len(lines); i++ {
		submatch = tagRegex.FindStringSubmatch(lines[i])
		if len(submatch) > 0 {
			result = append(result, fmt.Sprintf("[USR] %s %s", usernameRegex.FindString(submatch[0]), lines[i]))
		} else {
			result = append(result, lines[i])
		}
	}
	return result
}

func Run() {
	println(IsValidLine("[ERR] A good error here"))
	println(IsValidLine("Any old [ERR] text"))
	println(IsValidLine("[BOB] Any old text"), "\n")

	println("Split log line:")
	split := SplitLogLine("section 1<=>section 2<-*~*->section 3")
	for i := 0; i < len(split); i++ {
		println(split[i])
	}

	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,     // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}
	println("Quoted passwords (expected - 2):", CountQuotedPasswords(lines))
	println("Remove end of line:", RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27"), "\n")

	result := TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	})
	for x := 0; x < len(result); x++ {
		println(result[x])
	}
}
