package Reversestring

import "github.com/golang/example/stringutil"

func ReverseBuiltIn(str string) string {
	return stringutil.Reverse(str)
}

func Reversestring(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Reversstring_Rune(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
