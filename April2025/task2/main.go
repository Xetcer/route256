package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type opType int

const banksCount int = 3

const (
	RuToUSD opType = iota
	RuToEUR
	USDToRu
	USDToEUR
	EURToRu
	EURToUSD
)

type Bank struct {
	Operation map[opType]float64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscanln(in, &n)
	// перебор по наборам входных данных
	for range n {
		banks := make([]*Bank, 0)
		// заполняем банки данными теста
		for range banksCount {
			banks = append(banks, FillBank(in))
		}
		result := MaxDollars(banks)
		// Универсальный вывод
		fmt.Fprintf(out, "%g\n", result)
	}

}

func NewBank() *Bank {
	return &Bank{
		Operation: make(map[opType]float64),
	}
}

func FillBank(in *bufio.Reader) *Bank {
	bank := NewBank()
	bank.Operation = make(map[opType]float64)
	for i := RuToUSD; i <= EURToUSD; i++ {
		var from, to int
		fmt.Fscan(in, &from, &to)
		bank.Operation[i] = float64(to) / float64(from)
	}
	return bank
}

// Функция для нахождения максимального количества долларов
func MaxDollars(banks []*Bank) float64 {
	// Инициализация максимальной суммы долларов
	maxDollars := -math.MaxFloat64

	// Определяем максимальное значение при прямом обмене среди всех банков Рубль-Доллар
	maxDollars = max(banks[0].Operation[RuToUSD], banks[1].Operation[RuToUSD], banks[2].Operation[RuToUSD])

	// Теперь максимальное через евро
	AB := banks[0].Operation[RuToEUR] * banks[1].Operation[EURToUSD]
	AC := banks[0].Operation[RuToEUR] * banks[2].Operation[EURToUSD]
	BA := banks[1].Operation[RuToEUR] * banks[0].Operation[EURToUSD]
	BC := banks[1].Operation[RuToEUR] * banks[2].Operation[EURToUSD]
	CA := banks[2].Operation[RuToEUR] * banks[0].Operation[EURToUSD]
	CB := banks[2].Operation[RuToEUR] * banks[1].Operation[EURToUSD]

	tempMax := max(AB, BC, AC, BA, CA, CB)
	if maxDollars < tempMax {
		maxDollars = tempMax
	}

	// рубль доллар/ доллар рубль /рубль доллар
	ABC := banks[0].Operation[RuToUSD] * banks[1].Operation[USDToRu] * banks[2].Operation[RuToUSD]
	ACB := banks[0].Operation[RuToUSD] * banks[2].Operation[USDToRu] * banks[1].Operation[RuToUSD]
	BAC := banks[1].Operation[RuToUSD] * banks[0].Operation[USDToRu] * banks[2].Operation[RuToUSD]
	BCA := banks[1].Operation[RuToUSD] * banks[2].Operation[USDToRu] * banks[0].Operation[RuToUSD]
	CAB := banks[2].Operation[RuToUSD] * banks[0].Operation[USDToRu] * banks[1].Operation[RuToUSD]
	CBA := banks[2].Operation[RuToUSD] * banks[1].Operation[USDToRu] * banks[0].Operation[RuToUSD]
	tempMax = max(ABC, ACB, BAC, BCA, CAB, CBA)
	if maxDollars < tempMax {
		maxDollars = tempMax
	}

	// рубль евро/ евро рубль /рубль доллар
	ABC = banks[0].Operation[RuToEUR] * banks[1].Operation[EURToRu] * banks[2].Operation[RuToUSD]
	ACB = banks[0].Operation[RuToEUR] * banks[2].Operation[EURToRu] * banks[1].Operation[RuToUSD]
	BAC = banks[1].Operation[RuToEUR] * banks[0].Operation[EURToRu] * banks[2].Operation[RuToUSD]
	BCA = banks[1].Operation[RuToEUR] * banks[2].Operation[EURToRu] * banks[0].Operation[RuToUSD]
	CAB = banks[2].Operation[RuToEUR] * banks[0].Operation[EURToRu] * banks[1].Operation[RuToUSD]
	CBA = banks[2].Operation[RuToEUR] * banks[1].Operation[EURToRu] * banks[0].Operation[RuToUSD]
	tempMax = max(ABC, ACB, BAC, BCA, CAB, CBA)
	if maxDollars < tempMax {
		maxDollars = tempMax
	}

	// рубль доллар/ доллар евро /евро доллар
	ABC = banks[0].Operation[RuToUSD] * banks[1].Operation[USDToEUR] * banks[2].Operation[EURToUSD]
	ACB = banks[0].Operation[RuToUSD] * banks[2].Operation[USDToEUR] * banks[1].Operation[EURToUSD]
	BAC = banks[1].Operation[RuToUSD] * banks[0].Operation[USDToEUR] * banks[2].Operation[EURToUSD]
	BCA = banks[1].Operation[RuToUSD] * banks[2].Operation[USDToEUR] * banks[0].Operation[EURToUSD]
	CAB = banks[2].Operation[RuToUSD] * banks[0].Operation[USDToEUR] * banks[1].Operation[EURToUSD]
	CBA = banks[2].Operation[RuToUSD] * banks[1].Operation[USDToEUR] * banks[0].Operation[EURToUSD]
	tempMax = max(ABC, ACB, BAC, BCA, CAB, CBA)
	if maxDollars < tempMax {
		maxDollars = tempMax
	}
	return maxDollars
}
