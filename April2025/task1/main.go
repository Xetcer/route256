package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		var str string
		fmt.Fscanln(in, &str)
		result := IsPattern(str)
		// fmt.Println("n:", n, "str:", str, "result", result)
		fmt.Fprintf(out, "%s\n", result)
	}
}

func IsPattern(s string) string {
	if len(s) == 1 {
		return "YES"
	}

	patternStr := (string)(s[0])
	// Формируем шаблон Regexp
	byteS := []byte(s)
	regexpStr := fmt.Sprintf("^(%s+)(([^%s]{1}[%s]{1})|[%s+])*$", patternStr, patternStr, patternStr, patternStr)

	// fmt.Println(regexpStr)
	re := regexp.MustCompile(regexpStr)
	if re.Match(byteS) {
		return "YES"
	} else {
		return "NO"
	}
}

func IsPatternbyBuf(s string) string {
	if len(s) == 1 {
		return "YES"
	}
	if s[0] != s[len(s)-1] {
		return "NO"
	}
	// var patternStr rune
	// var prev, current rune
	// checkIsPattern := false
	// // patternStr := (string)(s[0])
	// for i, value := range s {
	// 	if i == 0 {
	// 		patternStr = value
	// 		prev = value
	// 	} else {
	// 		if checkIsPattern {
	// 			if value != patternStr {
	// 				return "NO"
	// 			}
	// 			checkIsPattern = false
	// 		}

	// 		if value != prev {
	// 			checkIsPattern = true
	// 		}
	// 		prev = current
	// 		current = value
	// 	}
	// }
	patternStr := (string)(s[0])
	// Формируем шаблон Regexp
	byteS := []byte(s)
	regexpStr := fmt.Sprintf("^(%s+)(([^%s]{1}[%s]{1})|[%s+])*$", patternStr, patternStr, patternStr, patternStr)

	// fmt.Println(regexpStr)
	re := regexp.MustCompile(regexpStr)
	if re.Match(byteS) {
		return "YES"
	} else {
		return "NO"
	}
}
