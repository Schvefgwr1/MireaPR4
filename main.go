package main

import (
	"fmt"
)

func main() {
	var taskId int
	fmt.Println("Введите номер задачи")
	fmt.Scanf("%d\n", &taskId)
	switch taskId {
	case 1:
		task1()
	case 2:
		task2()
	case 3:
		task3()
	case 4:
		task4()
	case 5:
		task5()
	case 6:
		task6()
	case 7:
		task7()
	case 8:
		task8()
	case 9:
		task9()
	case 10:
		task10()
	default:
		fmt.Printf("%d задача не найдена\n", taskId)
	}
}

func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

func task1() {
	var base, height float64
	fmt.Print("Введите основание треугольника: ")
	fmt.Scanf("%f", &base)
	fmt.Print("Введите высоту треугольника: ")
	fmt.Scanf("%f", &height)
	fmt.Println("Площадь треугольника:", triangleArea(base, height))
}

func sortArray(arr []float64) []float64 {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func task2() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scanf("%d", &n)
	arr := make([]float64, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &arr[i])
	}
	fmt.Println("Отсортированный массив:", sortArray(arr))
}

func sumOfSquares(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i * i
	}
	return sum
}

func task3() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Сумма квадратов чётных чисел:", sumOfSquares(n))
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func task4() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Scanf("%s", &s)
	if isPalindrome(s) {
		fmt.Println("Это палиндром.")
	} else {
		fmt.Println("Это не палиндром.")
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func task5() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &n)
	if isPrime(n) {
		fmt.Println("Число простое.")
	} else {
		fmt.Println("Число не является простым.")
	}
}

func generatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func task6() {
	var limit int
	fmt.Print("Введите предел для генерации простых чисел: ")
	fmt.Scanf("%d", &limit)
	fmt.Println("Простые числа:", generatePrimes(limit))
}

func toBinary(n int) string {
	binary := ""
	for n > 0 {
		binary = fmt.Sprintf("%d", n%2) + binary
		n /= 2
	}
	return binary
}

func task7() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Двоичное представление числа:", toBinary(n))
}

func findMax(arr []int) int {
	maxT := arr[0]
	for _, value := range arr {
		if value > maxT {
			maxT = value
		}
	}
	return maxT
}

func task8() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Максимальный элемент:", findMax(arr))
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func task9() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Введите второе число: ")
	fmt.Scanf("%d", &b)
	fmt.Println("НОД:", gcd(a, b))
}

func sumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func task10() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Сумма элементов массива:", sumArray(arr))
}
