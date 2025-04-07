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
	even, odd := GetSubStr(s1, s2, lenS1, lenS2)

	if even != "" || odd != "" {
		return 1
	} else {
		return 0
	}
}

func GetSubStr(s1, s2 string, Len1, Len2 int) (even string, odd string) {
	even = ""
	odd = ""
	firsIndex := 0
	step := 1
	checkEven, checkOdd := true, true // если длины равны то ищем обе комбинации
	if Len1 != Len2 {
		if max(Len1, Len2)%2 == 0 {
			checkOdd = false
		} else {
			checkEven = false
			firsIndex = 1
		}
		step = 2
	}

	for i := firsIndex; i < min(Len1, Len2); i += step {
		if i%2 == 0 {
			if checkEven {
				if s1[i] == s2[i] {
					even += string(s1[i])
				} else {
					even = ""
					checkEven = false
				}
			}
		} else {
			if checkOdd {
				if s1[i] == s2[i] {
					odd += string(s1[i])
				} else {
					odd = ""
					checkOdd = false
				}
			}
		}
	}

	return even, odd
}
