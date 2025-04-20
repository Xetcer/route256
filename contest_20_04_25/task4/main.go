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
	var t int
	//
	fmt.Fscanln(in, &t)
	for range t {
		// TODO
		// var str string
		// fmt.Fscanln(in, &str)
		// result := IsPattern(str)
		// // fmt.Println("n:", n, "str:", str, "result", result)
		// fmt.Fprintf(out, "%s\n", result)
	}
}
