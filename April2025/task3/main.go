package main

import (
	"bufio"
	"fmt"
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
	// перебор по наборам входных данных
	for range t {
		fmt.Fscanln(in, &n)
		var newStr string // каждая новая строка
		strMap := make(map[string]int)
		evenMap := make(map[string]int)
		oddMap := make(map[string]int)
		for range n {
			fmt.Fscanln(in, &newStr)
			even, odd := getKeys(newStr)
			evenMap[even]++
			if odd != "" {
				oddMap[odd]++
				strMap[newStr]++
			}
		}

		oddSum := getPairsCount(oddMap)
		evenSum := getPairsCount(evenMap)
		intersecSum := getPairsCount(strMap)
		pairs := oddSum + evenSum - intersecSum
		fmt.Fprintf(out, "%d\n", pairs)
		clear(strMap)
	}
}

func getPairsCount(counts map[string]int) int {
	sum := 0
	for _, count := range counts {
		sum += count * (count - 1) / 2 // количество пар
	}
	return sum
}

func getKeys(s string) (even, odd string) {
	even, odd = "", ""
	for i := range s {
		if i%2 == 0 {
			even += string(s[i])
		} else {
			odd += string(s[i])
		}
	}
	return even, odd
}
