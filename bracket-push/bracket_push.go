// Package brackets contains tools around brackets
package brackets

import "fmt"

type bracketStack []rune

// push pushes a new opening bracket onto the stack
func (s bracketStack) push(v rune) bracketStack {
	return append(s, v)
}

// pop pops a matching opening bracket from the stack. Returns an error if the beackets don't match
func (s bracketStack) pop(close rune) (bracketStack, error) {
	if len(s) == 0 {
		return s, fmt.Errorf("Can't pop from empty stack")
	}

	open := s[len(s)-1]

	if !bracketsMatch(open, close) {
		return s[:len(s)-1], fmt.Errorf("Bracket %v on stack doesn't match closing bracket %v", open, close)
	}

	return s[:len(s)-1], nil
}

// bracketsMatch returns whether the brackets match
func bracketsMatch(open, close rune) bool {
	if open == '(' && close == ')' {
		return true
	}
	if (open == '{' || open == '[') && close == open+2 {
		return true
	}
	return false
}

// Bracket returns whether all brackets are nested correctly and all are closed
func Bracket(in string) (bool, error) {
	var stack = bracketStack{}

	for _, c := range in {
		if c == '{' || c == '[' || c == '(' {
			stack = stack.push(c)
		}
		if c == '}' || c == ']' || c == ')' {
			s, err := stack.pop(c)
			if err != nil {
				return false, nil
			}
			stack = s
		}
	}

	if len(stack) > 0 {
		return false, nil
	}

	return true, nil
}
