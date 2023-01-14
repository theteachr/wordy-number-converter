package main

import (
	"fmt"
	"os"
	"strings"
)

type transformer struct {
	value    int
	modifier func(int, int) int
}

type stack struct {
	ts  []transformer
	top int
}

func newStack() *stack {
	return &stack{top: -1}
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (s *stack) push(t transformer) {
	s.ts = append(s.ts, t)
	s.top++
}

func (s *stack) pop() (transformer, bool) {
	if s.isEmpty() {
		return transformer{}, false
	}

	popped := s.ts[s.top]
	s.top--

	return popped, true
}

func curried_stack_mul(s *stack) func(int, int) int {
	return func(a int, b int) int {
		n, ok := s.pop()

		if !ok {
			n.value = 1
		}

		return a + b*n.value
	}
}

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "No input provided :o\n")
		os.Exit(1)
	}

	value := 0
	s := newStack()

	word_to_number := map[string]transformer{
		"one":       transformer{value: 1, modifier: add},
		"two":       transformer{value: 2, modifier: add},
		"three":     transformer{value: 3, modifier: add},
		"four":      transformer{value: 4, modifier: add},
		"five":      transformer{value: 5, modifier: add},
		"six":       transformer{value: 6, modifier: add},
		"seven":     transformer{value: 7, modifier: add},
		"eight":     transformer{value: 8, modifier: add},
		"nine":      transformer{value: 9, modifier: add},
		"ten":       transformer{value: 10, modifier: add},
		"eleven":    transformer{value: 11, modifier: add},
		"twelve":    transformer{value: 12, modifier: add},
		"thirteen":  transformer{value: 13, modifier: add},
		"fourteen":  transformer{value: 14, modifier: add},
		"fifteen":   transformer{value: 15, modifier: add},
		"sixteen":   transformer{value: 16, modifier: add},
		"seventeen": transformer{value: 17, modifier: add},
		"eighteen":  transformer{value: 18, modifier: add},
		"nineteen":  transformer{value: 19, modifier: add},
		"twenty":    transformer{value: 20, modifier: add},
		"thirty":    transformer{value: 30, modifier: add},
		"forty":     transformer{value: 40, modifier: add},
		"fifty":     transformer{value: 50, modifier: add},
		"sixty":     transformer{value: 60, modifier: add},
		"seventy":   transformer{value: 70, modifier: add},
		"eighty":    transformer{value: 80, modifier: add},
		"ninety":    transformer{value: 90, modifier: add},
		"hundred":   transformer{value: 100, modifier: mul},
		"thousand":  transformer{value: 1000, modifier: mul},
		"million":   transformer{value: 1_000_000, modifier: mul},
		"billion":   transformer{value: 1_000_000_000, modifier: mul},
	}

	worded_number := os.Args[1]
	words := strings.Split(strings.Trim(worded_number, " "), " ")

	for _, word := range words {
		t, ok := word_to_number[word]

		if !ok {
			fmt.Fprintf(os.Stderr, "Seems like you had a typo :<\n")
			os.Exit(2)
		}

		fmt.Println(t.value, value)
		value = t.modifier(value, t.value)
	}

	for {
		t, ok := s.pop()

		if !ok {
			break
		}

		fmt.Println(t.value, value)
	}

	fmt.Println(value)
}
