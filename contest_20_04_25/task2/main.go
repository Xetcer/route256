package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type factsType int

const factsCount int = 3

// Типы фактов
const (
	is factsType = iota
	same
	younger
	older
)

// общая структура факта
type factInfo struct {
	age   int       // величина возраста
	name  string    // имя
	ratio factsType // соотношение
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int //число наборов данных
	fmt.Fscanln(in, &t)
	for range t {
		sentense, _ := in.ReadString('\n')
		QPerson := GetPersonName(sentense) // Имя человека, возраст которого надо определить
		fmt.Println("Person Name is ", QPerson)
		personFacts := make(map[string][]factInfo, 0)
		for range factsCount {
			factSentense, _ := in.ReadString('\n')
			fmt.Println(factSentense)
			words := strings.Fields(factSentense)
			fact := factInfo{}
			switch len(words) {
			case 5:
				{
					age, _ := strconv.Atoi(words[2])
					fact.age = age
					fact.ratio = is
					fact.name = words[0]
				}
			default:
				{
					// получаем имя
					fact.name = words[len(words)-1]
					// Получаем соотношение
					if words[3] == "same" {
						fact.ratio = same
					} else {
						if words[4] == "younger" {
							fact.ratio = younger
						} else {
							fact.ratio = older
						}
						age, _ := strconv.Atoi(words[2])
						fact.age = age
					}

				}
			}
			personFacts[words[0]] = append(personFacts[words[0]], fact)
		}
		fmt.Println("Facts:", personFacts)
		// TODO
		// var str string
		// fmt.Fscanln(in, &str)
		// result := IsPattern(str)
		// // fmt.Println("n:", n, "str:", str, "result", result)
		// fmt.Fprintf(out, "%s\n", result)
	}
}

func GetPersonName(s string) string {
	words := strings.Fields(s)
	return words[len(words)-1]
}

func NewFact() *factInfo {
	return &factInfo{}
}
