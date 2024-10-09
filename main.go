package main

import (
	"fmt"
	"math"
)

func main() {
	var taskId, moduleId int
	fmt.Println("Введите номер модуля")
	fmt.Scanf("%d\n", &moduleId)
	fmt.Println("Введите номер задачи")
	fmt.Scanf("%d\n", &taskId)
	xy := fmt.Sprintf("%d_%d", moduleId, taskId)
	switch xy {
	case "1_1":
		task1_1()
	case "1_2":
		task1_2()
	case "1_3":
		task1_3()
	case "1_4":
		task1_4()
	case "1_5":
		task1_5()
	case "2_1":
		task2_1()
	case "2_2":
		task2_2()
	case "2_3":
		task2_3()
	case "2_4":
		task2_4()
	case "2_5":
		task2_5()
	case "3_1":
		task3_1()
	case "3_2":
		task3_2()
	case "3_3":
		task3_3()
	case "3_4":
		task3_4()
	case "3_5":
		task3_5()

	default:
		fmt.Printf("%d задача в модуле %d не найдена\n", taskId, moduleId)

	}
}

func task1_1() {
	fmt.Println("Enter a 4-th value")
	var number int
	fmt.Scanf("%d", &number)
	res := number/1000 + number%1000/100 + number%100/10 + number%10
	fmt.Println(res)
}

func task1_2() {
	var syst string
	var temp float32
	fmt.Println("Enter a base temperature system: C/F")
	fmt.Scanf("%s", &syst)
	fmt.Println(syst)
	switch syst {
	case "C":
		fmt.Println("Enter a value in Celsius")
		fmt.Scanf("%f", &temp)
		fmt.Printf("Temperature in Fahrenheit: %.2f\n", temp*1.8+32)
	case "F":
		fmt.Println("Enter a value in Fahrenheit")
		fmt.Scanf("%f", &temp)
		fmt.Printf("Temperature in Celsia: %.2f\n", (temp-32)/1.8)
	default:
		fmt.Println("Incorrect system, please restart the program and retry entering")
	}
}

func task1_3() {
	var length int
	var number float32
	fmt.Println("Enter a length of array:")
	fmt.Scanf("%d", &length)
	arr := make([]float32, length)
	for i := 0; i < length; i++ {
		fmt.Println("Enter a number of array: ")
		fmt.Scanf("%f", &number)
		arr[i] = number * 2
	}
	fmt.Println("Array: ", arr)
}

func task1_4() {
	var str string
	flag := false
	for !flag {
		fmt.Println("Enter a new part of string or \"end\" for end of program")
		var partOfStr string
		fmt.Scanf("%s", &partOfStr)
		if partOfStr == "end" {
			flag = true
		} else {
			str += partOfStr
			str += " "
		}
	}
	fmt.Println("String: ", str)
}

type Point struct {
	x float64
	y float64
}

func (p Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func task1_5() {
	p1 := Point{x: 0, y: 0}
	p2 := Point{x: 0, y: 0}
	fmt.Println("Введите x и y в формате: x_1 y_1")
	fmt.Scanf("%f %f", &p1.x, &p1.y)
	fmt.Println("Введите x и y в формате: x_2 y_2")
	fmt.Scanf("%f %f", &p2.x, &p2.y)
	distance := p1.DistanceTo(&p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

func task2_1() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanf("%d", &num)
	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

func task2_2() {
	var year int
	fmt.Println("Введите номер года")
	fmt.Scanf("%d", &year)
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}
}

func task2_3() {
	var x1, x2, x3 float32
	fmt.Println("Введите три числа через пробел")
	fmt.Scanf("%f %f, %f", &x1, &x2, &x3)
	fmt.Println("Наибольшее число")
	if x1 >= x2 && x1 >= x3 {
		fmt.Println(x1)
	} else if x2 >= x1 && x2 >= x3 {
		fmt.Println(x2)
	} else {
		fmt.Println(x3)
	}
}

func task2_4() {
	var age int
	fmt.Println("Введите возраст")
	fmt.Scanf("%d", &age)
	fmt.Println("Возрастная группа: ")
	switch {
	case age < 12:
		fmt.Println("Ребенок")
	case age >= 12 && age < 18:
		fmt.Println("Подросток")
	case age >= 18 && age < 65:
		fmt.Println("Взрослый")
	default:
		fmt.Println("Пожилой")
	}
}

func task2_5() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanf("%d", &num)
	if num%5 == 0 && num%3 == 0 {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func task3_1() {
	var n int
	fmt.Println("Введите число:")
	fmt.Scanf("%d", &n)
	result := factorial(n)
	fmt.Printf("Факториал числа %d равен %d\n", n, result)
}

func fibonacci(n int) []int {
	fib := make([]int, n)
	if n >= 1 {
		fib[0] = 0
	}
	if n >= 2 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func task3_2() {
	var n int
	fmt.Println("Введите число:")
	fmt.Scanf("%d", &n)
	result := fibonacci(n)
	fmt.Println("Первые ", n, " чисел Фибонначи: ", result)
}

func readArray() *[]int {
	var arr []int
	var elem, n int
	fmt.Print("Введите количество чисел в масиве: ")
	fmt.Scan(&n)
	fmt.Println("Введите исходный массив:")
	for i := 0; i < n; i++ {
		fmt.Scan(&elem)
		arr = append(arr, elem)
	}
	return &arr
}

func task3_3() {
	arr := *readArray()
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println("Перевернутый массив:", arr)
}

func task3_4() {
	var n int
	fmt.Print("Введите число для поиска простых чисел: ")
	fmt.Scan(&n)

	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	fmt.Println("Простые числа до", n, ":", primes)
}

func task3_5() {
	arr := *readArray()
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("Сумма чисел в массиве:", sum)
}
