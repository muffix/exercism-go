// Package bob implements Bob's responses
package bob

import "strings"

// Hey returns Bob's response based on what is said to him
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	uppercaseRemark := strings.ToUpper(remark)

	isQuestion := strings.HasSuffix(remark, "?")
	isAllCaps := remark == uppercaseRemark && uppercaseRemark != strings.ToLower(uppercaseRemark)

	if isAllCaps && isQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isAllCaps {
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
