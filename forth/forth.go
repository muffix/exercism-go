// Package forth implement an evaluator for a very simple subset of Forth
package forth

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []int
type operation func(stack) (stack, error)

func (s stack) push(v int) stack {
	return append(s, v)
}

func (s stack) pop() (stack, int, error) {
	l := len(s)
	if l == 0 {
		return nil, -1, fmt.Errorf("Stack empty")
	}
	return s[:l-1], s[l-1], nil
}

func buildOperations(input []string) map[string]operation {
	var operations = map[string]operation{
		"+":    func(s stack) (stack, error) { return s.twoAdic("+") },
		"-":    func(s stack) (stack, error) { return s.twoAdic("-") },
		"*":    func(s stack) (stack, error) { return s.twoAdic("*") },
		"/":    func(s stack) (stack, error) { return s.twoAdic("/") },
		"dup":  (stack).dup,
		"drop": (stack).drop,
		"swap": func(s stack) (stack, error) { return s.twoAdic("swap") },
		"over": (stack).over,
	}

	for _, params := range input {
		overrideParams := strings.Split(params, " ")
		if overrideParams[0] == ":" && overrideParams[len(overrideParams)-1] == ";" {
			operations[strings.ToLower(overrideParams[1])] = func(s stack) (stack, error) {
				return s.runInstructions(overrideParams[2:len(overrideParams)-1], operations)
			}
		}
	}

	return operations
}

// Forth returns the stack after performing the instructions in the input
func Forth(input []string) ([]int, error) {
	s := stack{}
	operations := buildOperations(input[:len(input)-1])

	s, err := s.runInstructions(strings.Split(input[len(input)-1], " "), operations)

	return s, err
}

func (s stack) runInstructions(instructions []string, operations map[string]operation) (stack, error) {
	for _, instr := range instructions {

		if op, ok := operations[strings.ToLower(instr)]; ok {
			st, err := op(s)
			if err != nil {
				return nil, err
			}
			s = st
			continue
		}

		number, err := strconv.Atoi(instr)
		if err != nil {
			return nil, err
		}

		s = s.push(number)

	}
	return s, nil
}

func (s stack) twoAdic(op string) (stack, error) {
	if len(s) < 2 {
		return s, fmt.Errorf("Empty stack")
	}

	s, y, _ := s.pop()
	s, x, _ := s.pop()

	var res int
	switch op {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		if y == 0 {
			return s, fmt.Errorf("Division by zero")
		}
		res = x / y
	case "swap":
		s = s.push(y)
		s = s.push(x)
		return s, nil
	}

	s = s.push(res)

	return s, nil
}

func (s stack) dup() (stack, error) {
	if len(s) < 1 {
		return s, fmt.Errorf("Empty stack")
	}
	s = s.push(s[len(s)-1])
	return s, nil
}

func (s stack) drop() (stack, error) {
	s, _, err := s.pop()
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s stack) over() (stack, error) {
	if len(s) < 2 {
		return s, fmt.Errorf("Empty stack")
	}
	s = s.push(s[len(s)-2])
	return s, nil
}
