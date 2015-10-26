// Copyright 2015 The Go Authors. All rights reserved.
/*
Package owl implements a regular expression parser.
Feel free to contact the author 18817874087@163.com if you have any question.
*/
package owl

import (
	"errors"
	"fmt"
)

// RegToPost convert a regular expression to a postfix expression.
func RegToPost(reg string) (string, error) {
	var post string
	type Paren struct {
		NumOfOperand  int
		NumOfParallel int
	}
	var NumOfOperand = 0
	var NumOfParallel = 0
	p := []Paren{}
	pIndex := 0
	addDot := func() {
		NumOfOperand--
		for ; NumOfOperand > 0; NumOfOperand-- {
			post += "."
		}
	}
	addVerticalBar := func() {
		for ; NumOfParallel > 0; NumOfParallel-- {
			post += "|"
		}
	}
	for _, r := range reg {
		if NumOfOperand > 2 {
			fmt.Println("oh my god!")
		}
		switch r {
		case '(':
			if NumOfOperand > 1 {
				NumOfOperand--
				post += "."
			}
			p = append(p, Paren{NumOfOperand, NumOfParallel})
			pIndex++
			NumOfOperand = 0
			NumOfParallel = 0
		case '|':
			if NumOfOperand == 0 {
				return "", errors.New("Wrong exp")
			}
			addDot()
			NumOfParallel++
		case ')':
			if pIndex == 0 || NumOfOperand == 0 {
				return "", errors.New("Wrong exp")
			}
			addDot()
			addVerticalBar()
			pIndex--
			NumOfOperand = p[pIndex].NumOfOperand + 1
			NumOfParallel = p[pIndex].NumOfParallel
			p = p[:pIndex]
		case '*':
			fallthrough
		case '+':
			fallthrough
		case '?':
			if NumOfOperand == 0 {
				return "", errors.New("Wrong exp")
			}
			post += string(r)
		default:
			if NumOfOperand > 1 {
				NumOfOperand--
				post += "."
			}
			post += string(r)
			NumOfOperand++
		}
	}
	if pIndex != 0 {
		return "", errors.New("Wrong exp")
	}
	addDot()
	addVerticalBar()
	return post, nil
}

const (
	Match = 256
	Split = 257
)

type State struct {
	content  interface{}
	out1     *State
	out2     *State
	lastlist int
}

var NumState int

func NewState(c interface{}, out1 *State, out2 *State) *State {
	NumState++
	s = &State{c, out1, out2}
	return s
}

type Fragment struct {
	start *State
	out   []*State
}

func NewFragment(start *State, out []*State) *Fragment {
	f = &Fragment{start, out}
}
func Patch(out []*State, s *State) {
	length := len(out)
	for i := 0; i < length; i++ {
		out[i] = s
	}
}
func PostToNfa(postfix string) *State {

}
