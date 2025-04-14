package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Delivery struct {
	Time   int
	Image  int
	Server int
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int // количество тестов
	fmt.Fscanln(in, &t)
	// fmt.Println("Tests:", t)
	for range t {
		var n int // количество серверов
		var m int // количество изображений
		fmt.Fscanln(in, &n)
		// fmt.Println("ServersCount:", n)
		var line string // временная строка для считывания параметров сервера и изображений
		line, _ = in.ReadString('\n')
		// fmt.Fscanln(in, &line) // читаем пропускную способность серверов
		// fmt.Println("servers:", line)
		serverThrouput := parseInts(line)
		fmt.Fscanln(in, &m)
		// fmt.Println("ImagesCount:", m)
		line, _ = in.ReadString('\n')
		// fmt.Fscanln(in, &line) // читаем  вес изображений
		// fmt.Println("images:", line)
		imageWeight := parseInts(line)
		// fmt.Println(t)
		// fmt.Println(n)
		// fmt.Println(serverThrouput)
		// fmt.Println(m)
		// fmt.Println(imageWeight)

		delivery := make([]Delivery, 0)
		// Сохраним все пары (время доставки, номер изображения, номер сервера) в срез delivery
		for serverNum, throuput := range serverThrouput {
			for imageNum, weight := range imageWeight {
				time := (weight + throuput - 1) / throuput
				delivery = append(delivery, Delivery{Time: time, Image: imageNum, Server: serverNum})
			}
		}
		// отсортируем срез от меньшего(быстрого) до большего(медленного) времени
		sort.Slice(delivery, func(i, j int) bool {
			return delivery[i].Time < delivery[j].Time
		})
		// сколько раз изображение j встречается в текущем диапазоне
		imageCounter := make([]int, m)
		// Сколько разных изображений покрыто в текущем диапазоне
		processedImagesCount := 0
		// увеличивает счётчик изображения, отмечая его как обработанное.
		inc := func(image int) {
			if imageCounter[image] == 0 {
				processedImagesCount++
			}
			imageCounter[image]++
		}
		// уменьшает счётчик, и, если изображение больше не покрывается, оно удаляется из диапазона.
		dec := func(image int) {
			imageCounter[image]--
			if imageCounter[image] == 0 {
				processedImagesCount--
			}
		}
		/*
		   Запускаем метод двух указателей:
		   – right движется вперёд, добавляя варианты доставки.
		   – Если все изображения покрыты (processedImagesCount
		   == imagesCount), пытаемся минимизировать разницу
		   между delivery[right].Time и delivery[left].Time.
		   – Если нашли лучший вариант, обновляем minDeliveryGap.
		   – Затем двигаем left, убирая вариант доставки.
		*/
		minDeliveryGap := math.MaxInt32
		var rangeStart, rangeEnd int
		for left, right := 0, 0; right < len(delivery); right++ {
			inc(delivery[right].Image)
			for left <= right && processedImagesCount == m {
				deliveryGap := delivery[right].Time - delivery[left].Time
				if minDeliveryGap > deliveryGap {
					minDeliveryGap = deliveryGap
					rangeStart, rangeEnd = left, right
				}
				dec(delivery[left].Image)
				left++
			}
		}
		/*
			После нахождения оптимального диапазона, каждое
			изображение imageServer[i] закрепляется за сервером,
			который его обработает в найденном диапазоне.
		*/
		imageServer := make([]int, m)
		for i := rangeStart; i <= rangeEnd; i++ {
			imageServer[delivery[i].Image] = delivery[i].Server
		}
		// fmt.Println(delivery)

		fmt.Fprintln(out, minDeliveryGap)
		OutputImageServers(imageServer, out)
	}
}

// Функция для преобразования строки в срез целых чисел
func parseInts(s string) []int {
	stringNums := strings.Fields(s) // Разделяем строку по пробелам
	nums := make([]int, len(stringNums))

	for i, str := range stringNums {
		num, err := strconv.Atoi(str) // Преобразуем строку в целое число
		if err != nil {
			continue // Обработка ошибки (в данном случае игнорируем)
		}
		nums[i] = num
	}

	return nums
}

// вывод серверов
func OutputImageServers(imageServer []int, out *bufio.Writer) {
	for i, troughput := range imageServer {
		if i > 0 {
			out.WriteString(" ") // Добавляем пробел перед каждым элементом, кроме первого
		}
		out.WriteString(strconv.Itoa(troughput + 1)) // Преобразуем число в строку и записываем
	}
	out.WriteString(" \n") // Добавляем переход на новую строку после всех элементов
}
