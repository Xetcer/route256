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
	var t int // количество наборов данных
	fmt.Fscanln(in, &t)

	// перебор по наборам входных данных
	for range t {
		var k, n, m int32
		var s string            // текущая строка доски
		fmt.Fscanln(in, &k)     // длина победной комбинации
		fmt.Fscanln(in, &n, &m) // количество строк и столбцов
		a := make([][]int8, n)
		// заполняем игральную доску
		for i := range n {
			fmt.Fscanln(in, &s)
			a[i] = make([]int8, m)
			for j := range m {
				switch s[j] {
				case '.':
					a[i][j] = 2
				case '0':
					a[i][j] = 0
				case 'X':
					a[i][j] = 1
				}
			}
		}
		result := getAnswer(a, m, n, k)
		clear(a)
		// Универсальный вывод
		fmt.Fprintf(out, "%s\n", result)
	}
}

func getAnswer(a [][]int8, m, n, k int32) string {
	/*
		X и Y являются текущими позициями таргета. N и M — количество
		столбцов и строк в доске. Функция вернёт true в случае, если
		позиция таргета корректна, иначе вернёт false.
	*/
	ok := func(x, y int32) bool {
		return 0 <= x && x < n && 0 <= y && y < m
	}

	/*
		Функция win используется для проверки факта того, есть ли в
		заданных координатах последовательность из X или 0. Направление
		проверки задается вектором (dx, dy).
		(1, 0) для проверки вертикального направления (сверху вниз);
		(1, 1) — для диагонали (сверху слева → вниз направо).
	*/
	win := func(x, y, dx, dy int32) bool {
		sum := make([]int32, 3)
		for ok(x, y) { // начинаем двигаться в заданном направлении
			sum[a[x][y]]++ // считаем суммы 1(Х) 0 2(.) в заданном направлении
			x += dx
			y += dy
			if ok(x-k*dx, y-k*dy) { // если сдвинулись на k элементов от начальной координаты
				if sum[0] == k || sum[1] == k { // смотрим есть ли в данном направлении k крестиков или 0
					return true
				}
				sum[a[x-k*dx][y-k*dy]]-- // уменьшаем сумму текущих элементов в заданном направлении.
			}
		}
		return false
	}
	/*
	   Функция preWin — аналог функции выше с одним исключением:
	   здесь происходит проверка того, можем ли мы поставить один X так,
	   чтобы собрать выигрышную позицию.
	   Функция считает количество крестиков sum[1] и пустых клеток
	   sum[2] по заданному k. Если есть одна пустая клетка и количество
	   крестиков равно (k − 1), функция вернёт true, иначе вернёт false.
	*/
	preWin := func(x, y, dx, dy int32) bool {
		sum := make([]int32, 3)
		for ok(x, y) {
			sum[a[x][y]]++
			x += dx
			y += dy
			if ok(x-k*dx, y-k*dy) { // если сдвинулись на k элементов от начальной координаты
				if sum[1] == k-1 && sum[2] == 1 { // смотрим если сумма крестиков равна k-1 и есть 1 пустое место в направлении
					return true
				}
				sum[a[x-k*dx][y-k*dy]]-- // уменьшаем сумму текущих элементов в заданном направлении.
			}
		}
		return false
	}

	type pair struct {
		f      func(int32, int32, int32, int32) bool
		answer string
	}

	// переберем все варианты, спера с NO потом с YES до первого совпадения
	for _, p := range []pair{{win, "NO"}, {preWin, "YES"}} {
		// двигаемся в строке направо по всем строкам
		for i := int32(0); i < n; i++ {
			if p.f(i, 0, 0, 1) {
				return p.answer
			}
		}
		// двигаемся в столбце вниз
		for i := int32(0); i < m; i++ {
			if p.f(0, i, 1, 0) {
				return p.answer
			}
		}
		// двигаемся по диагоняли направо вниз
		for i := int32(0); i < n; i++ { // сдвигаемся по строке вправо на каждой итерации
			if p.f(i, 0, 1, 1) {
				return p.answer
			}
		}
		for j := int32(0); j < m; j++ { // сдвигаемся по столбцу вниз на каждой итерации
			if p.f(0, j, 1, 1) {
				return p.answer
			}
		}

		// двигаемся по диагонали вниз влево
		for j := int32(0); j < m; j++ { // сдвигаемся по столбцу вниз на каждой итерации
			if p.f(0, j, 1, -1) {
				return p.answer
			}
		}
		for i := int32(0); i < n; i++ { // сдвигаемся по строке вправо на каждой итерации
			if p.f(i, m-1, 1, -1) {
				return p.answer
			}
		}

	}
	return "NO"
}
