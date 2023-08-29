package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsBalanced("(()()()(()))"))
}

func IsBalanced(str string) bool {

	if str == "" || len(str) == 0 {
		return false
	}

	br := map[rune]int{
		'(': 0,
		')': 1,
		'{': 2,
		'}': 3,
		'[': 4,
		']': 5,
	}

	rstr := []rune(str)
	var st []rune
	var lastBrInd = -1
	for i := 0; i < len(rstr); i++ {
		ch := rstr[i]
		ind, ok := br[ch]
		if !ok {
			return false
		}
		if ind%2 != 0 {
			if len(st) == 0 {
				return false
			}
			lastBr := st[lastBrInd]
			st = st[:lastBrInd]
			if br[lastBr] != br[ch]-1 {
				return false
			}
			lastBrInd--
		} else {
			st = append(st, ch)
			lastBrInd++
		}
	}
	return len(st) == 0
}
