package main

import "fmt"

func main() {
	var first, second bool //значення по замовчуваню для булевого типу = false
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)  // false
	fmt.Println("second = ", second) // false
	//Перші два записи отримали значення по замовчуванню

	fmt.Println("third  = ", third)  // true
	fmt.Println("fourth = ", fourth) // false - оператор НЕ (!) інвертує значення. Саме тому true змінилось на false

	fmt.Println("fifth  = ", fifth, "\n") // true

	fmt.Println("!true  = ", !true)        // false
	fmt.Println("!false = ", !false, "\n") // true - оператор !
	// знову використаний оператор НЕ (!)

	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false
	// Логічне І. Дає true лише в тому випадку, коли одна з умов = true

	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false
	// Логічне АБО. Дає true коли одна з умов = true

	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false
	// Вище записи - порівняння чисел. Коли умова істинна - повертає true

	//Задание.
	//1. Пояснить результаты операций
}
