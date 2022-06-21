// На вхід подано стрінг з цілими числа, котри розділені пробілами.
// Треба повернути найбільше та найменше число.
// Наприклад:
// input := "1 2 3 4 5" // повертає "5 1"
// input := "1 9 3 4 -5" // повертає "9 -5"
// Уточнення:
// 1. Всі числа є не більше, ніж int32. Використовуйте цей тип даних.
// 2. В стрінгі завжди буде принаймні одне число.
// 3. Результатом має бути стрінг, в якому два числа розділені пробілом (або одне, якщо дано було лише одне
// число). Найбільше число має бути першим.
// 4. Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
// func main(){
//     input := "1 9 3 4 -5"
//     var result string
//     //...
//     // тут має бути ваш код
//     // змінна result в кінці функції має тримати стрінг з правильними результатами, згідно до умови задачі
// }

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := "0 -4 -5 0 10 18 33"
	var result string
	var min, max, number int
	var initFlag bool
	for _, v := range strings.Split(input, " ") {

		number, _ = strconv.Atoi(v)
		if !initFlag {
			min = number
			max = min
			initFlag = true
			continue
		}
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}

	}
	result = strconv.Itoa(max)
	if min != max {
		result += " "
		result += strconv.Itoa(min)
	}
	fmt.Println(result)
}
