// Package markdown implements a simple Markdown parser
package markdown

// implementation to refactor

import (
	"fmt"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	var inEm, inStrong bool
	list := 0
	header := 0

	html := ""
	for pos := 0; pos < len(markdown); pos++ {
		char := markdown[pos]
		switch char {
		case '#':
			for ; char == '#'; char = markdown[pos] {
				header++
				pos++
			}
			html += fmt.Sprintf("<h%d>", header)
		case '*':
			if list == 0 {
				html += "<ul>"
			}
			list++
			html += "<li>"
			pos++
		case '\n':
			if list > 0 {
				html += "</li>"
			} else if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		case '_':
			if pos+1 < len(markdown) && markdown[pos+1] == '_' {
				inStrong = !inStrong
				if inStrong {
					html += "<strong>"
				} else {
					html += "</strong>"
				}
				pos++
			} else {
				inEm = !inEm
				if inEm {
					html += "<em>"
				} else {
					html += "</em>"
				}
			}
		default:
			html += string(char)
		}
	}

	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}

	if list > 0 {
		return html + "</li></ul>"
	}

	return "<p>" + html + "</p>"
}
