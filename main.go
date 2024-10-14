package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"unicode"
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

func fromBaseToDecimal(number string, base int) int {
	var decimal int
	for _, digit := range strings.ToUpper(number) {
		var value int
		if digit >= '0' && digit <= '9' {
			value = int(digit - '0')
		} else {
			value = int(digit - 'A' + 10)
		}
		decimal = decimal*base + value
	}
	return decimal
}

func fromDecimalToBase(decimal, base int) string {
	if decimal == 0 {
		return "0"
	}

	var result string
	for decimal > 0 {
		remainder := decimal % base
		var digit rune
		if remainder < 10 {
			digit = rune('0' + remainder)
		} else {
			digit = rune('A' + remainder - 10)
		}
		result = string(digit) + result
		decimal /= base
	}
	return result
}

func task1_1() {
	var number string
	var fromBase, toBase int

	fmt.Print("Введите число: ")
	fmt.Scanf("%s", &number)
	fmt.Print("Введите исходную систему счисления (2-36): ")
	fmt.Scanf("%d", &fromBase)
	fmt.Print("Введите конечную систему счисления (2-36): ")
	fmt.Scanf("%d", &toBase)

	decimal := fromBaseToDecimal(number, fromBase)
	result := fromDecimalToBase(decimal, toBase)

	fmt.Printf("Число %s из системы %d в системе %d: %s\n", number, fromBase, toBase, result)
}

func task1_2() {
	var a, b, c float64

	fmt.Print("Введите коэффициент a: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Введите коэффициент b: ")
	fmt.Scanf("%f", &b)
	fmt.Print("Введите коэффициент c: ")
	fmt.Scanf("%f", &c)

	if a == 0 {
		if b != 0 {
			x := -c / b
			fmt.Printf("Линейное уравнение. Корень: %.2f\n", x)
		} else {
			if c == 0 {
				fmt.Println("Бесконечное множество решений.")
			} else {
				fmt.Println("Нет решений.")
			}
		}
		return
	}

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Два действительных корня: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if discriminant == 0 {
		x := -b / (2 * a)
		fmt.Printf("Один действительный корень: x = %.2f\n", x)
	} else {
		realPart := -b / (2 * a)
		imagPart := math.Sqrt(-discriminant) / (2 * a)
		fmt.Printf("Комплексные корни: x1 = %.2f+%.2fi, x2 = %.2f-%.2fi\n", realPart, imagPart, realPart, imagPart)
	}
}

func task1_3() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scanf("%d", &n)

	numbers := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	sort.Slice(numbers, func(i, j int) bool {
		return math.Abs(float64(numbers[i])) < math.Abs(float64(numbers[j]))
	})

	fmt.Println("Отсортированный массив по модулю:", numbers)
}

func mergeSortedArrays(a, b []int) []int {
	var merged []int
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	for i < len(a) {
		merged = append(merged, a[i])
		i++
	}
	for j < len(b) {
		merged = append(merged, b[j])
		j++
	}

	return merged
}

func task1_4() {
	var n, m int

	fmt.Print("Введите количество элементов первого массива: ")
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	fmt.Println("Введите элементы первого отсортированного массива:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	fmt.Print("Введите количество элементов второго массива: ")
	fmt.Scanf("%d", &m)
	b := make([]int, m)
	fmt.Println("Введите элементы второго отсортированного массива:")
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}

	merged := mergeSortedArrays(a, b)
	fmt.Println("Слияние двух отсортированных массивов:", merged)
}

func indexOf(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func task1_5() {
	var haystack, needle string

	fmt.Print("Введите основную строку: ")
	fmt.Scanf("%s", &haystack)
	fmt.Print("Введите подстроку для поиска: ")
	fmt.Scanf("%s", &needle)

	index := indexOf(haystack, needle)
	fmt.Println("Индекс первого вхождения подстроки:", index)
}

func task2_1() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Введите оператор (+, -, *, /, ^, %): ")
	fmt.Scanf("%s", &operator)
	fmt.Print("Введите второе число: ")
	fmt.Scanf("%f", &b)

	switch operator {
	case "+":
		fmt.Printf("Результат: %.2f\n", a+b)
	case "-":
		fmt.Printf("Результат: %.2f\n", a-b)
	case "*":
		fmt.Printf("Результат: %.2f\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: Деление на ноль.")
		} else {
			fmt.Printf("Результат: %.2f\n", a/b)
		}
	case "^":
		fmt.Printf("Результат: %.2f\n", math.Pow(a, b))
	case "%":
		ai, bi := int(a), int(b)
		if bi == 0 {
			fmt.Println("Ошибка: Деление на ноль.")
		} else {
			fmt.Printf("Результат: %d\n", ai%bi)
		}
	default:
		fmt.Println("Ошибка: Недопустимый оператор.")
	}
}

func isPalindrome(s string) bool {
	var cleaned []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}

func task2_2() {
	var input string
	fmt.Print("Введите строку для проверки на палиндром: ")
	fmt.Scanf("%[^\n]s", &input)
	if isPalindrome(input) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

type Segment struct {
	start, end float64
}

func task2_3() {
	var segments [3]Segment

	for i := 0; i < 3; i++ {
		fmt.Printf("Введите начальную и конечную точку %d-го отрезка: ", i+1)
		fmt.Scanf("%f %f", &segments[i].start, &segments[i].end)
		if segments[i].start > segments[i].end {
			segments[i].start, segments[i].end = segments[i].end, segments[i].start
		}
	}
	maxStart := segments[0].start
	for i := 1; i < 3; i++ {
		if segments[i].start > maxStart {
			maxStart = segments[i].start
		}
	}

	minEnd := segments[0].end
	for i := 1; i < 3; i++ {
		if segments[i].end < minEnd {
			minEnd = segments[i].end
		}
	}

	if maxStart <= minEnd {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func filterWord(word string) string {
	filtered := []rune{}
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, r)
		}
	}
	return string(filtered)
}

func task2_4() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите предложение: ")
	sentence, _ := reader.ReadString('\n')

	words := strings.Fields(sentence)

	var longestWord string
	maxLength := 0

	for _, word := range words {
		filteredWord := filterWord(word)
		if len(filteredWord) > maxLength {
			longestWord = filteredWord
			maxLength = len(filteredWord)
		}
	}
	if maxLength == 0 {
		fmt.Println("Нет слов в предложении.")
	} else {
		fmt.Printf("Самое длинное слово: %s\n", longestWord)
	}
}

func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%100 != 0 {
		return true
	} else if year%400 == 0 {
		return true
	}
	return false
}

func task2_5() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scanf("%d", &year)
	if isLeapYear(year) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
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
	fmt.Print("Введите максимальное число для чисел Фибоначчи: ")
	fmt.Scanf("%d", &n)
	if n < 0 {
		fmt.Println("Числа Фибоначчи не могут быть отрицательными.")
		return
	}
	a, b := 0, 1
	for a <= n {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func task3_2() {
	var start, end int
	fmt.Print("Введите начальное число диапазона: ")
	fmt.Scanf("%d", &start)
	fmt.Print("Введите конечное число диапазона: ")
	fmt.Scanf("%d", &end)
	if start > end {
		start, end = end, start
	}
	fmt.Println("Простые числа в диапазоне:")
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func isArmstrong(num int) bool {
	if num < 0 {
		return false
	}

	n := num
	count := 0
	for n != 0 {
		n /= 10
		count++
	}

	sum := 0
	n = num
	for n != 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(count)))
		n /= 10
	}

	return sum == num
}

func task3_3() {
	var start, end int
	fmt.Print("Введите начальное число диапазона: ")
	fmt.Scanf("%d", &start)
	fmt.Print("Введите конечное число диапазона: ")
	fmt.Scanf("%d", &end)

	if start > end {
		start, end = end, start
	}

	fmt.Println("Числа Армстронга в диапазоне:")
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func task3_4() {
	var input string
	fmt.Print("Введите строку для реверса: ")
	fmt.Scanf("%[^\n]s", &input)
	reversed := reverseString(input)
	fmt.Println("Реверс строки:", reversed)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func task3_5() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Введите второе число: ")
	fmt.Scanf("%d", &b)

	result := gcd(a, b)
	fmt.Println("НОД:", result)
}
