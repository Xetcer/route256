package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		var qty int
		fmt.Fscanln(in, &qty)
		for j := 0; j < qty; j++ {
			var str string
			fmt.Fscanln(in, &str)
			result := IsPattern(str)
			fmt.Println("str:", str, "result", result)
			fmt.Fprintf(out, "%s\n", result)
		}
	}
}

func IsPattern(s string) string {
	if len(s) == 1 {
		return "YES"
	}
	charsMap := make(map[rune]int)
	for _, char := range s {
		char = unicode.ToLower(char)
		charsMap[char]++
		if (len(charsMap) >= 2) ||
			(charsMap[char] >= 3) {
			return "YES"
		}
	}
	return "NO"
}
