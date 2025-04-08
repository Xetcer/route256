package main

import (
	"bufio"
	"fmt"
	"os"
)

type SubString struct {
	evenStr string // четная подстрока
	oddStr  string // нечетная подстрока
}

type Str struct {
	SubStr SubString // подстроки данной строки
	count  int       // количество таких строк
}

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
		var str string // каждая новая строка
		strMap := make(map[string]Str, 0)
		evenMap := make(map[string]int, 0)
		oddMap := make(map[string]int, 0)
		for range n {
			fmt.Fscanln(in, &str)
			if value, ok := strMap[str]; ok {
				subStr := value
				subStr.count++
				strMap[str] = subStr
				evenMap[subStr.SubStr.evenStr]++
				oddMap[subStr.SubStr.oddStr]++
			} else {
				even, odd := getKeys(str)
				strMap[str] = Str{SubStr: SubString{evenStr: even, oddStr: odd}, count: 1}
				if even != "" {
					evenMap[even]++
				}
				if odd != "" {
					oddMap[odd]++
				}
			}
		}
		pairs := 0
		for _, keyStrs := range strMap {
			isCompare := false
			if value, ok := evenMap[keyStrs.SubStr.evenStr]; ok {
				if value > 1 {
					pairs += keyStrs.count
					isCompare = true
				}
			}
			if !isCompare {
				if value, ok := oddMap[keyStrs.SubStr.oddStr]; ok {
					if value > 1 {
						pairs += keyStrs.count
					}
				}
			}
		}
		// fmt.Println("stringsMap:", strMap, "evenMap:", evenMap, "oddMap", oddMap, "Pairs", getPairsCount(pairs))

		fmt.Fprintf(out, "%d\n", getPairsCount(pairs))
		clear(strMap)
	}
}

func getPairsCount(count int) int {
	if count > 1 {
		return count * (count - 1) / 2 // количество пар
	}
	return 0
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
