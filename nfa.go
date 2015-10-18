// Copyright 2015 The Go Authors. All rights reserved.
/*
Package owl implements a regular expression parser.
Feel free to contact the author 18817874087@163.com if you have any question.
*/
package owl

import (
	"errors"
	_ "fmt"
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
	for _, r := range reg {
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
			for {
				NumOfOperand--
				if NumOfOperand > 0 {
					post += "."
				} else {
					break
				}
			}
			NumOfParallel++
		case ')':
			if pIndex == 0 || NumOfOperand == 0 {
				return "", errors.New("Wrong exp")
			}
			for {
				NumOfOperand--
				if NumOfOperand > 0 {
					post += "."
				} else {
					break
				}
			}
			for ; NumOfParallel > 0; NumOfParallel-- {
				post += "|"
			}
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
	for {
		NumOfOperand--
		if NumOfOperand > 0 {
			post += "."
		} else {
			break
		}
	}
	for ; NumOfParallel > 0; NumOfParallel-- {
		post += "|"
	}
	return post, nil
}
