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
	charsMap := make(map[rune]int)
	// составляем карту символов
	for _, char := range s {
		charsMap[char]++
	}

	// Находим главную букву в строке (должна повторяться больше всех раз)
	var patternStr string
	max := 0
	for k, value := range charsMap {
		if max < value {
			patternStr = string(k)
			max = value
		}
	}

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
