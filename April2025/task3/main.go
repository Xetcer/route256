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
	var n int
	fmt.Fscanln(in, &n)
	result := "TODO"
	// перебор по наборам входных данных
	for range n {
		// banks := make([]*Bank, 0)
		// // заполняем банки данными теста
		// for range banksCount {
		// 	banks = append(banks, FillBank(in))
		// }
		// result := MaxDollars(banks)
		// // Универсальный вывод
		fmt.Fprintf(out, "%s\n", result)
	}

}
