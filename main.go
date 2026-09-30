package main

import "fmt"

//lab do 26.09 1/2

func calculateStats(numbers []int) (int, float64) {
	if len(numbers) == 0 {
		return 0, 0.0
	}

	sum := 0
	for _, v := range numbers {
		sum += v
	}

	avg := float64(sum) / float64(len(numbers))
	return sum, avg
}

func findElement(slice []string, value string) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}
	return -1, false
}

//lab do 26.09 2/2

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

type Person struct {
	Name string
	Age  int
}

func celebrateBirthday(p *Person) {
	p.Age++
}

func main() {
//lab do 26.09 1/2

	nums := []int{1, 2, 3, 4, 5}
	sum, avg := calculateStats(nums)
	fmt.Printf("Срез чисел: %v\n", nums)
	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее: %.2f\n", avg)

	empty := []int{}
	sum2, avg2 := calculateStats(empty)
	fmt.Printf("Пустой срез чисел: %v\n", empty)
	fmt.Printf("Сумма: %v\n", sum2)
	fmt.Printf("Среднее: %.2f\n", avg2)

	words := []string{"Marlen", "Maria", "Ilona"}

	index, found := findElement(words, "Maria")
	fmt.Printf("Ищем %v: Индекс: %v, Найден: %v\n", "Maria", index, found)

	index, found = findElement(words, "Madina")
	fmt.Printf("Ищем %v: Индекс: %v, Найден: %v\n", "Madina", index, found)

//lab do 26.09 2/2

	a := 10
	b := 20

	fmt.Printf("До swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("После swap: a=%d, b=%d\n\n", a, b)

	vasya := Person{Name: "Василий", Age: 30}

	fmt.Printf("До дня рождения: %s, возраст=%d\n", vasya.Name, vasya.Age)
	celebrateBirthday(&vasya)
	fmt.Printf("После дня рождения: %s, возраст=%d\n", vasya.Name, vasya.Age)
}
