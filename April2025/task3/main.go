package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int // количество наборов строк
	var n int // количество строк в текущем наборе
	fmt.Fscanln(in, &t)
	result := 0 // число пар похожих строк
	// перебор по наборам входных данных
	for range t {
		fmt.Fscanln(in, &n)
		var str string // каждая новая строка
		strings := make([]string, 0)
		for range n {
			fmt.Fscanln(in, &str)
			strings = append(strings, str)
		}
		if len(strings) > 1 {
			result = GetSimilarStrCount(strings)
		} else {
			result = 0
		}

		fmt.Fprintf(out, "%d\n", result)
		clear(strings)
		result = 0
	}
}

func GetSimilarStrCount(strs []string) int {
	SimilarPairs := 0
	// fmt.Println("New strings:", strs)
	for i := range len(strs) - 1 {
		for j := i + 1; j < len(strs); j++ {
			SimilarPairs += IsSimilar(strs[i], strs[j])
		}
	}
	return SimilarPairs
}

func IsSimilar(s1, s2 string) int {
	lenS1 := len(s1)
	lenS2 := len(s2)
	if math.Abs(float64(lenS1)-float64(lenS2)) > 1 {
		return 0
	}
	if lenS1 == 1 && lenS2 == 1 {
		if s1[0] == s2[0] {
			return 1
		} else {
			return 0
		}
	} else if s1 == s2 {
		return 1
	}
	// выделим из нее четные и нечетные подстроки
	even1, odd1 := GetSubStr(s1)
	even2, odd2 := GetSubStr(s2)

	if even1 == even2 || odd1 == odd2 {
		return 1
	} else {
		return 0
	}
}

func GetSubStr(s string) (even string, odd string) {
	even = ""
	odd = ""
	for i, char := range s {
		if i%2 == 0 {
			even += string(char)
		} else {
			odd += string(char)
		}
	}
	return even, odd
}
