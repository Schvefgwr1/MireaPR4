package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
	case 11:
		task11()
	case 12:
		task12()
	case 13:
		task13()
	case 14:
		task14()
	case 15:
		task15()
	default:
		fmt.Printf("%d задача не найдена\n", taskId)

	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("Число %d не является простым. Первый делитель: %d\n", n, i)
			return false
		}
	}
	return true
}

func task1() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if isPrime(number) {
		fmt.Printf("Число %d является простым.\n", number)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func task2() {
	var a, b int
	fmt.Print("Введите два числа: ")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("НОД чисел %d и %d: %d\n", a, b, gcd(a, b))
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Printf("Шаг %d: %v\n", i+1, arr)
	}
}

func task3() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите числа массива через пробел:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	arr := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Некорректный ввод. Пожалуйста, вводите только целые числа.")
			return
		}
		arr[i] = num
	}

	fmt.Println("Исходный массив:", arr)
	bubbleSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}

func task4() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

var memo = map[int]int{}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func task5() {
	var n int
	fmt.Print("Введите номер числа Фибоначчи: ")
	fmt.Scan(&n)

	fmt.Printf("Число Фибоначчи %d: %d\n", n, fibonacci(n))
}

func reverseNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

func task6() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("Обратное число: %d\n", reverseNumber(number))
}

func pascalTriangle(levels int) {
	triangle := make([][]int, levels)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	for _, row := range triangle {
		fmt.Println(row)
	}
}

func task7() {
	var levels int
	fmt.Print("Введите количество уровней треугольника Паскаля: ")
	fmt.Scanf("%d", &levels)
	pascalTriangle(levels)
}

func isPalindrome(n int) bool {
	original, reversed := n, 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}

func task8() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &number)

	if isPalindrome(number) {
		fmt.Printf("Число %d является палиндромом.\n", number)
	} else {
		fmt.Printf("Число %d не является палиндромом.\n", number)
	}
}

func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	minArr, maxArr := arr[0], arr[0]
	for _, num := range arr[1:] {
		if num < minArr {
			minArr = num
		}
		if num > maxArr {
			maxArr = num
		}
	}
	return minArr, maxArr
}

func task9() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите числа массива через пробел:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	arr := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Некорректный ввод. Пожалуйста, вводите только целые числа.")
			return
		}
		arr[i] = num
	}

	minArr, maxArr := findMinMax(arr)
	fmt.Printf("Минимум: %d, Максимум: %d\n", minArr, maxArr)
}

func task10() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1
	var guess, attempts int
	const maxAttempts = 10

	fmt.Println("Я загадал число от 1 до 100. Попробуйте угадать!")

	for attempts < maxAttempts {
		fmt.Print("Ваш ответ: ")
		fmt.Scanf("%d", &guess)
		attempts++

		if guess < secret {
			fmt.Println("Больше!")
		} else if guess > secret {
			fmt.Println("Меньше!")
		} else {
			fmt.Printf("Поздравляю! Вы угадали число %d за %d попыток!\n", secret, attempts)
			return
		}
	}

	fmt.Printf("Вы исчерпали попытки! Загаданное число было: %d\n", secret)
}

func isArmstrong(n int) bool {
	temp := n
	sum := 0
	digits := int(math.Log10(float64(n))) + 1

	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	return sum == n
}

func task11() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &number)

	if isArmstrong(number) {
		fmt.Printf("Число %d является числом Армстронга.\n", number)
	} else {
		fmt.Printf("Число %d не является числом Армстронга.\n", number)
	}
}

func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func task12() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	wordCount := countWords(input)
	fmt.Println("Количество уникальных слов:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

const size = 10

func printBoard(board [size][size]bool) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func countNeighbors(board [size][size]bool, x, y int) int {
	neighbors := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			if nx >= 0 && nx < size && ny >= 0 && ny < size && board[nx][ny] {
				neighbors++
			}
		}
	}
	return neighbors
}

func nextGeneration(board [size][size]bool) [size][size]bool {
	var newBoard [size][size]bool
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			neighbors := countNeighbors(board, i, j)
			if board[i][j] && (neighbors == 2 || neighbors == 3) {
				newBoard[i][j] = true
			} else if !board[i][j] && neighbors == 3 {
				newBoard[i][j] = true
			}
		}
	}
	return newBoard
}

func task13() {
	board := [size][size]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, true, true, true, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	for generation := 0; generation < 5; generation++ {
		fmt.Printf("Generation %d:\n", generation)
		printBoard(board)
		board = nextGeneration(board)
	}
}

func digitalRoot(n int) int {
	for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}
	return n
}

func task14() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("Цифровой корень числа %d: %d\n", number, digitalRoot(number))
}

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			roman += symbols[i]
			num -= val[i]
		}
	}
	return roman
}

func task15() {
	var number int
	fmt.Print("Введите арабское число: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("Римское представление числа %d: %s\n", number, intToRoman(number))
}
