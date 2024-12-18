// https://leetcode.com/problems/construct-string-with-repeat-limit/description/

package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type CharCount struct {
	char  rune
	count int
}

type CharCounter []*CharCount

var (
	errEmptyCharCounter     = errors.New("empty char counter")
	errIndexOutOfRange      = errors.New("index out of range")
	errZeroCountCharCounter = errors.New("char counter zeroed")
	nilRune                 = '\x00'
)

type test struct {
	s           string
	repeatLimit int
	answer      string
}

func main() {
	var testCase = []test{
		{
			s:           "cczazcc",
			repeatLimit: 3,
			answer:      "zzcccac",
		},
		{
			s:           "aababab",
			repeatLimit: 2,
			answer:      "bbabaa",
		},
	}

	for _, test := range testCase {
		res := repeatLimitedString(test.s, test.repeatLimit)
		fmt.Println("input:", test.s, "output:", res, "answer:", test.answer, "| correct:", res == test.answer)
	}
}

func repeatLimitedString(s string, repeatLimit int) string {
	charCounter := newCharCounter(s)

	var (
		builder strings.Builder
		counter int
		temp    = nilRune
	)

	for {
		char, err := charCounter.Peek()
		if err != nil {
			break
		}

		if char == temp {
			counter++
		} else {
			counter = 1
		}
		temp = char

		var c rune
		if counter <= repeatLimit {
			c, err = charCounter.PopChar()
			if err != nil {
				break
			}

		} else {
			counter = 1
			c, err = charCounter.CharFromIndex(1)
			if err != nil {
				break
			}
			temp = c
		}
		builder.WriteRune(c)
	}

	return builder.String()
}

func newCharCounter(s string) CharCounter {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	var cc CharCounter
	for i := 'z'; i >= 'a'; i-- {
		if v, ok := m[i]; ok {
			cc = append(cc, &CharCount{
				char:  i,
				count: v,
			})
		}
	}

	return cc
}

func (cc CharCounter) Print() {
	for _, c := range cc {
		fmt.Println(string(c.char), c.count)
	}
}

func (cc CharCounter) PeekIndex(idx int) (rune, error) {
	if idx >= len(cc) {
		return nilRune, errIndexOutOfRange
	}

	return cc[idx].char, nil
}

func (cc CharCounter) Peek() (rune, error) {
	if len(cc) == 0 {
		return nilRune, errEmptyCharCounter
	}

	return cc[0].char, nil
}

func (cc *CharCounter) CharFromIndex(idx int) (rune, error) {
	if idx >= len(*cc) {
		return nilRune, errIndexOutOfRange
	}

	if (*cc)[idx].count == 0 {
		return nilRune, errZeroCountCharCounter
	}

	c := (*cc)[idx].char
	(*cc)[idx].count--
	if (*cc)[idx].count == 0 {
		if idx >= len(*cc) {
			*cc = (*cc)[:len(*cc)-1]
		} else {
			*cc = slices.Delete(*cc, idx, idx+1)
		}
	}

	return c, nil
}

func (cc *CharCounter) PopChar() (rune, error) {
	for len(*cc) != 0 {
		if (*cc)[0].count != 0 {
			c := (*cc)[0].char

			(*cc)[0].count--
			if (*cc)[0].count == 0 {
				if len(*cc) != 0 {
					*cc = (*cc)[1:]
				} else {
					*cc = CharCounter{}
				}
			}

			return c, nil
		}

		if len(*cc) != 0 {
			*cc = slices.Delete(*cc, 0, 1)
		} else {
			*cc = CharCounter{}
		}
	}

	return nilRune, errEmptyCharCounter
}
