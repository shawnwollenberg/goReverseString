package Reversestring

func Reversestring(str string) (result string) {
	//fmt.Println(str)
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Reversstring_backup(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

/*func main() {
	fmt.Println(Reversestring("shawn"))
}*/

//create _test.go
//go mod init
//run "go test"
