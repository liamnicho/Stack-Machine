package main

import (
	"errors"
	"strconv"
	"strings"
)

func StackMachine(commands string) (int, error) {
	if commands == "" {
		return 0, errors.New("")
	}
	commandsList := strings.Split(commands, " ")

	stack := make([]int, 0) // creating a list of integers
	for i := 0; i < len(commandsList); i++ {

		if commandsList[i] == "POP" {
			stack = stack[:len(stack)-1]

		} else if commandsList[i] == "DUP" {
			last := stack[len(stack)-1]
			stack = append(stack, last)

		} else if commandsList[i] == "+" {
			if len(stack) < 2 {
				return 0, errors.New("")
			}

			a := stack[len(stack)-1]     // read the last element
			stack = stack[:len(stack)-1] // delete the last element

			b := stack[len(stack)-1]     // read the last element
			stack = stack[:len(stack)-1] //

			add := a + b
			if add > 50000 {
				return 0, errors.New("")
			}

			stack = append(stack, add) // insert in to stack

		} else if commandsList[i] == "-" {
			if len(stack) < 2 {
				return 0, errors.New("")
			}

			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			sub := a - b
			if sub < 0 {
				return 0, errors.New("")
			}

			stack = append(stack, sub) // push the result
		} else if commandsList[i] == "*" {
			if len(stack) < 2 {
				return 0, errors.New("")
			}

			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			mul := a * b
			if mul > 50000 {
				return 0, errors.New("")
			}

			stack = append(stack, mul)
		} else if commandsList[i] == "SUM" {
			if len(stack) == 0 {
				return 0, errors.New("")
			}

			s := 0
			for i := 0; i < len(stack); i++ {
				s = s + stack[i]
			}

			stack = stack[:0]
			stack = append(stack, s) // push the result

		} else if commandsList[i] == "CLEAR" {
			stack = stack[:0]
		} else {
			a, err := strconv.ParseInt(commandsList[i], 10, 64)

			if err != nil {
				return 0, errors.New("")
			} else if (a > 50000) || a < 0 {
				return 0, errors.New("")
			}

			stack = append(stack, int(a)) // insert in to the stack
		}
	}

	if len(stack) == 0 {
		return 0, errors.New("")
	}
	return stack[len(stack)-1], nil
}

func main() {
	// main is unused - run using
	// go test ./...
}
