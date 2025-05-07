package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)
	for range t {
		var n int         // количество высказываний
		var action string // совершаемое действие
		fmt.Fscanln(in, &n)
		peoples := make(map[string]int)    // карта, в которой будем вести подсчет
		peoplesAns := make(map[string]int) // карта, в которой будем формировать вывод людей
		parseSentense := func(s string) {
			words := strings.Fields(s)
			name := strings.TrimSuffix(words[0], ":")
			// Если человека нет в карте, добавим его
			if _, isExist := peoples[name]; !isExist {
				peoples[name] = 0
			}
			// отрицаем или утверждаем?
			isNot := false
			if len(words) == 5 {
				isNot = true
			}
			// Получим имя о ком говорим
			who := words[1]
			// проверим говорим о ком-то ?
			isMyself := true
			if words[2] == "is" {
				isMyself = false
			}
			if isMyself {
				if isNot {
					peoples[name]--
				} else {
					peoples[name] += 2
				}
			} else {
				if isNot {
					peoples[who]--
				} else {
					peoples[who]++
				}
			}
			// Получаем действие
			action = strings.TrimSuffix(words[len(words)-1], "!")
		}
		// считаем все утверждения данного набора и заполним карту
		for range n {
			var sentense string // временная строка для считывания одного утверждения
			sentense, _ = in.ReadString('\n')
			parseSentense(sentense)
		}
		// fmt.Println("Peoples: ", peoples)
		// fmt.Println("action: ", action)
		// создадим срез ключей для того чтобы отсортировать в лексикографичесом порядке
		// namesByOrder := make([]string, 0, len(peoples))
		// for key := range peoples {
		// 	namesByOrder = append(namesByOrder, key)
		// }

		// fmt.Println("names", names)
		// отсортируем ключи в лексикографическом порядке
		// sort.Strings(namesByOrder)
		// fmt.Println("names sort", names)
		maxPoints := -0xffff // его очки
		for cName := range peoples {
			if peoples[cName] > maxPoints {
				clear(peoplesAns)
				maxPoints = peoples[cName]
				peoplesAns[cName] = maxPoints
			} else if peoples[cName] == maxPoints {
				peoplesAns[cName] = maxPoints
			}
		}
		// Теперь отсортируем массив для вывода по порядку
		namesByOrder := make([]string, 0, len(peoplesAns))
		for key := range peoplesAns {
			namesByOrder = append(namesByOrder, key)
		}
		sort.Strings(namesByOrder)
		// Выводим результат
		for _, name := range namesByOrder {
			fmt.Fprintf(out, "%s is %s.\n", name, action)
		}
	}
}
